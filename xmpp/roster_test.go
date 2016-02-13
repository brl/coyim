package xmpp

import (
	"encoding/xml"

	"github.com/twstrike/coyim/xmpp/data"

	. "gopkg.in/check.v1"
)

type RosterXmppSuite struct{}

var _ = Suite(&RosterXmppSuite{})

type testStanzaValue struct{}

func (s *RosterXmppSuite) Test_ParseRoster_failsIfItDoesntReceiveAClientIQ(c *C) {
	rep := data.Stanza{
		Name:  xml.Name{Local: "Foobarium"},
		Value: testStanzaValue{},
	}

	_, err := data.ParseRoster(rep)
	c.Assert(err.Error(), Equals, "xmpp: roster request resulted in tag of type Foobarium")
}

func (s *RosterXmppSuite) Test_ParseRoster_failsIfTheRosterContentIsIncorrect(c *C) {
	rep := data.Stanza{
		Name: xml.Name{Local: "iq"},
		Value: &data.ClientIQ{
			Query: []byte("<foo></bar>"),
		},
	}

	_, err := data.ParseRoster(rep)
	c.Assert(err.Error(), Equals, "expected element type <query> but have <foo>")
}
