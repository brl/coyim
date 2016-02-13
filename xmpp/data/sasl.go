package data

import (
	"encoding/xml"
	"fmt"
)

// SaslMechanisms contains information about SASL mechanisms
// RFC 3920  C.4  SASL name space
//TODO RFC 6120 obsoletes RFC 3920
type SaslMechanisms struct {
	XMLName   xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl mechanisms"`
	Mechanism []string `xml:"mechanism"`
}

type saslAuth struct {
	XMLName   xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl auth"`
	Mechanism string   `xml:"mechanism,attr"`
}

type saslChallenge string

type saslResponse string

// SaslAbort signifies a SASL abort
type SaslAbort struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl abort"`
}

// SaslSuccess signifies a SASL Success
type SaslSuccess struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl success"`
	Content []byte   `xml:",innerxml"`
}

// SaslFailure signifies a SASL Failure
type SaslFailure struct {
	XMLName          xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl failure"`
	Text             string   `xml:"text,omitempty"`
	DefinedCondition Any      `xml:",any"`
}

// Condition returns a SASL-related error condition
func (f SaslFailure) Condition() SASLErrorCondition {
	return SASLErrorCondition(f.DefinedCondition.XMLName.Local)
}

func (f SaslFailure) String() string {
	if f.Text != "" {
		return fmt.Sprintf("%s: %q", f.Condition(), f.Text)
	}

	return string(f.Condition())
}

// SASLErrorCondition represents a defined SASL-related error conditions as defined in RFC 6120, section 6.5
type SASLErrorCondition string

// SASL error conditions as defined in RFC 6120, section 6.5
const (
	SASLAborted              SASLErrorCondition = "aborted"
	SASLAccountDisabled                         = "account-disabled"
	SASLCredentialsExpired                      = "credentials-expired"
	SASLEncryptionRequired                      = "encryption-required"
	SASLIncorrectEncoding                       = "incorrect-encoding"
	SASLInvalidAuthzid                          = "invalid-authzid"
	SASLInvalidMechanism                        = "invalid-mechanism"
	SASLMalformedRequest                        = "malformed-request"
	SASLMechanismTooWeak                        = "mechanism-too-weak"
	SASLNotAuthorized                           = "not-authorized"
	SASLTemporaryAuthFailure                    = "temporary-auth-failure"
)
