package xmpp

import (
	"encoding/xml"
	"io"
	"net"
	"strings"

	"golang.org/x/net/proxy"
)

// A Dialer connects and authenticates to an XMPP server
type Dialer struct {
	// JID represents the user's "bare JID" as specified in RFC 6120
	JID string

	// Password used to authenticate to the server
	Password string

	// ServerAddress associates a particular FQDN with the origin domain specified by the JID.
	ServerAddress string

	// Proxy configures a proxy used to connect to the server
	Proxy proxy.Dialer

	// Config configures the XMPP protocol
	Config Config
}

func (d *Dialer) hasCustomServer() bool {
	return d.ServerAddress != ""
}

func (d *Dialer) getJIDLocalpart() string {
	parts := strings.SplitN(d.JID, "@", 2)
	return parts[0]
}

func (d *Dialer) getJIDDomainpart() string {
	//TODO: remove any existing resourcepart although our doc says it is a bare JID (without resourcepart) but it would be nice
	parts := strings.SplitN(d.JID, "@", 2)
	return parts[1]
}

// GetServer returns the "hardcoded" server chosen if available, otherwise returns the domainpart from the JID. The server contains port information
func (d *Dialer) GetServer() string {
	if d.hasCustomServer() {
		return d.ServerAddress
	}

	return d.getFallbackServer()
}

func (d *Dialer) getFallbackServer() string {
	return net.JoinHostPort(d.getJIDDomainpart(), "5222")
}

// RegisterAccount registers an account on the server. The formCallback is used to handle XMPP forms.
func (d *Dialer) RegisterAccount(formCallback FormCallback) (*Conn, error) {
	//TODO: notify in case the feature is not supported
	d.Config.CreateCallback = formCallback
	return d.Dial()
}

// Dial creates a new connection to an XMPP server with the given proxy
// and authenticates as the given user.
func (d *Dialer) Dial() (*Conn, error) {
	// Starting an XMPP connection comprises two parts:
	// - Opening a transport channel (TCP)
	// - Opening an XML stream over the transport channel

	// RFC 6120, section 3
	conn, err := d.newTCPConn()
	if err != nil {
		return nil, err
	}

	// RFC 6120, section 4
	return d.setupStream(conn)
}

// RFC 6120, Section 4.2
func (d *Dialer) setupStream(conn net.Conn) (c *Conn, err error) {
	if d.hasCustomServer() {
		d.Config.TrustedAddress = true
	}

	c = newConn()
	c.config = d.Config
	c.originDomain = d.getJIDDomainpart()
	d.bindTransport(c, conn)

	if err := d.negotiateStreamFeatures(c, conn); err != nil {
		return nil, err
	}

	go c.watchKeepAlive(conn)
	go c.watchPings()

	return c, nil
}

// RFC 6120, section 4.3.2
func (d *Dialer) negotiateStreamFeatures(c *Conn, conn net.Conn) error {
	if err := c.sendInitialStreamHeader(); err != nil {
		return err
	}

	// STARTTLS MUST be the first feature to be negotiated
	if err := d.negotiateSTARTTLS(c, conn); err != nil {
		return err
	}

	if registered, err := d.negotiateInBandRegistration(c); err != nil || registered {
		return err
	}

	// SASL negotiation. RFC 6120, section 6
	if err := d.negotiateSASL(c); err != nil {
		return err
	}

	//TODO: negotiate other features

	return nil
}

func (d *Dialer) bindTransport(c *Conn, conn net.Conn) {
	c.in, c.out = makeInOut(conn, d.Config)
	c.rawOut = conn
	c.keepaliveOut = &timeoutableConn{conn, keepaliveTimeout}
}

func makeInOut(conn io.ReadWriter, config Config) (in *xml.Decoder, out io.Writer) {
	if config.InLog != nil {
		in = xml.NewDecoder(io.TeeReader(conn, config.InLog))
	} else {
		in = xml.NewDecoder(conn)
	}

	if config.OutLog != nil {
		out = io.MultiWriter(conn, config.OutLog)
	} else {
		out = conn
	}

	return
}
