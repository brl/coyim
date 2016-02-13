// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xmpp implements the XMPP IM protocol, as specified in RFC 6120 and
// 6121.
package xmpp

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/twstrike/coyim/xmpp/data"
	xe "github.com/twstrike/coyim/xmpp/errors"
	"github.com/twstrike/coyim/xmpp/interfaces"

	"github.com/twstrike/coyim/sasl"
	"github.com/twstrike/coyim/sasl/digestmd5"
	"github.com/twstrike/coyim/sasl/plain"
	"github.com/twstrike/coyim/sasl/scram"
)

var (
	errUnsupportedSASLMechanism = errors.New("xmpp: server does not support any of the prefered SASL mechanism")
)

func init() {
	plain.Register()
	digestmd5.Register()
	scram.Register()
}

// SASL negotiation. RFC 6120, section 6
func (d *dialer) negotiateSASL(c interfaces.Conn) error {
	user := d.getJIDLocalpart()
	password := d.password

	if err := c.Authenticate(user, password); err != nil {
		return xe.ErrAuthenticationFailed
	}

	// RFC 6120, section 6.3.2. Restart the stream
	err := c.SendInitialStreamHeader()
	if err != nil {
		return err
	}

	return c.BindResource()
}

func (c *conn) Authenticate(user, password string) error {
	// TODO: Google accounts with 2-step auth MUST use app-specific passwords:
	// https://security.google.com/settings/security/apppasswords
	// An alternative would be implementing the Google authentication mechanisms
	// - X-OAUTH2: requires app registration on Google, and a server to receive the oauth callback
	// https://developers.google.com/talk/jep_extensions/oauth?hl=en
	// - X-GOOGLE-TOKEN: seems to be this https://developers.google.com/identity/protocols/AuthForInstalledApps

	return c.authenticateWithPreferedMethod(user, password)
}

func (c *conn) authenticateWithPreferedMethod(user, password string) error {
	//TODO: this should be configurable by the client
	preferedMechanisms := []string{"SCRAM-SHA-1", "DIGEST-MD5", "PLAIN"}

	log.Println("sasl: server supports mechanisms", c.features.Mechanisms.Mechanism)

	for _, prefered := range preferedMechanisms {
		if !sasl.ClientSupport(prefered) {
			continue
		}

		for _, m := range c.features.Mechanisms.Mechanism {
			if prefered == m {
				log.Println("sasl: authenticating via", m)
				return c.authenticatWith(prefered, user, password)
			}
		}
	}

	return errUnsupportedSASLMechanism
}

func clientNonce(r io.Reader) (string, error) {
	//TODO: what is the appropriate size for this?
	//TODO: what is the appropriate way to generate a cnonce?
	conceRand := make([]byte, 7)
	_, err := io.ReadFull(r, conceRand)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(conceRand), nil
}

func (c *conn) authenticatWith(mechanism string, user string, password string) error {
	clientAuth, err := sasl.NewClient(mechanism)
	if err != nil {
		return err
	}

	clientNonce, err := clientNonce(c.Rand())
	if err != nil {
		return err
	}

	clientAuth.SetProperty(sasl.AuthID, user)
	clientAuth.SetProperty(sasl.Password, password)

	clientAuth.SetProperty(sasl.Service, "xmpp")
	clientAuth.SetProperty(sasl.QOP, "auth")

	//TODO: this should come from username if it were a full JID
	//clientAuth.SetProperty(sasl.Realm, "")

	clientAuth.SetProperty(sasl.ClientNonce, clientNonce)

	t, err := clientAuth.Step(nil)
	if err != nil {
		return err
	}

	fmt.Fprintf(c.rawOut, "<auth xmlns='%s' mechanism='%s'>%s</auth>\n", NsSASL, mechanism, t.Encode())

	return c.challengeLoop(clientAuth)
}

func (c *conn) challengeLoop(clientAuth sasl.Session) error {
	for {
		t, success, err := c.receiveChallenge()
		if err != nil {
			return err
		}

		t, err = clientAuth.Step(t)
		if err != nil {
			return err
		}

		if success {
			return nil
		}

		if !clientAuth.NeedsMore() {
			break
		}

		fmt.Fprintf(c.rawOut, "<response xmlns='%s'>%s</response>\n", NsSASL, t.Encode())
	}

	return xe.ErrAuthenticationFailed
}

func (c *conn) receiveChallenge() (t sasl.Token, success bool, err error) {
	var encodedChallenge []byte

	name, val, _ := next(c)
	switch v := val.(type) {
	case *data.SaslFailure:
		err = errors.New("xmpp: authentication failure: " + v.String())
		return
	case *data.SaslSuccess:
		encodedChallenge = v.Content
		success = true
	case *string:
		if name.Local != "challenge" || name.Space != NsSASL {
			err = errors.New("xmpp: unexpected <" + name.Local + "> in " + name.Space)
			return
		}

		encodedChallenge = []byte(*v)
	}

	t, err = sasl.DecodeToken(encodedChallenge)
	return
}

// Resource binding. RFC 6120, section 7
func (c *conn) BindResource() error {
	// This is mandatory, so a missing features.Bind is a protocol failure
	fmt.Fprintf(c.out, "<iq type='set' id='bind_1'><bind xmlns='%s'/></iq>", NsBind)
	var iq data.ClientIQ
	if err := c.in.DecodeElement(&iq, nil); err != nil {
		return errors.New("unmarshal <iq>: " + err.Error())
	}
	c.jid = iq.Bind.Jid // our local id

	return c.establishSession()
}

// See RFC 3921, section 3.
func (c *conn) establishSession() error {
	if c.features.Session == nil {
		return nil
	}

	// The server needs a session to be established.
	fmt.Fprintf(c.out, "<iq to='%s' type='set' id='sess_1'><session xmlns='%s'/></iq>", c.originDomain, NsSession)
	var iq data.ClientIQ
	if err := c.in.DecodeElement(&iq, nil); err != nil {
		return errors.New("xmpp: unmarshal <iq>: " + err.Error())
	}

	if iq.Type != "result" {
		return errors.New("xmpp: session establishment failed")
	}

	return nil
}
