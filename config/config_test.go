package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

func init() {
	log.SetOutput(ioutil.Discard)
}

type ConfigXmppSuite struct{}

var _ = Suite(&ConfigXmppSuite{})

func init() {
	log.SetOutput(ioutil.Discard)
}

func (s *ConfigXmppSuite) TestParseYes(c *C) {
	c.Assert(ParseYes("Y"), Equals, true)
	c.Assert(ParseYes("y"), Equals, true)
	c.Assert(ParseYes("YES"), Equals, true)
	c.Assert(ParseYes("yes"), Equals, true)
	c.Assert(ParseYes("Yes"), Equals, true)
	c.Assert(ParseYes("anything"), Equals, false)
}

func (s *ConfigXmppSuite) TestSerializeAccountsConfig(c *C) {
	expected := `{
	"Accounts": [
		{
			"Account": "bob@riseup.net",
			"Peers": [
				{
					"UserID": "bob@riseup.net",
					"Fingerprints": null,
					"Nickname": "boby"
				}
			],
			"HideStatusUpdates": false,
			"RequireTor": true,
			"OTRAutoTearDown": false,
			"OTRAutoAppendTag": false,
			"OTRAutoStartSession": false,
			"AlwaysEncrypt": true,
			"ConnectAutomatically": false
		},
		{
			"Account": "bob@riseup.net",
			"Peers": null,
			"HideStatusUpdates": false,
			"RequireTor": false,
			"OTRAutoTearDown": false,
			"OTRAutoAppendTag": false,
			"OTRAutoStartSession": false,
			"ConnectAutomatically": false
		}
	],
	"Bell": false,
	"ConnectAutomatically": false,
	"Display": {
		"MergeAccounts": false,
		"ShowOnlyOnline": false,
		"HideFeedbackBar": false
	},
	"AdvancedOptions": false
}`

	conf := ApplicationConfig{
		Accounts: []*Account{
			&Account{
				Account:       "bob@riseup.net",
				Peers: []*Peer{
					&Peer{
						UserID: "bob@riseup.net",
						Nickname: "boby",
					},
				},
				RequireTor:    true,
				AlwaysEncrypt: true,
			},
			&Account{
				Account: "bob@riseup.net",
			},
		},
	}

	contents, err := json.MarshalIndent(conf, "", "\t")
	c.Assert(err, IsNil)
	c.Assert(string(contents), Equals, expected)
}

func (s *ConfigXmppSuite) TestFindConfigFile(c *C) {
	conf := findConfigFile("")
	if strings.HasSuffix(conf, ".enc") {
		c.Assert(conf, Equals, os.Getenv("HOME")+"/.config/coyim/accounts.json.enc")
	} else {
		c.Assert(conf, Equals, os.Getenv("HOME")+"/.config/coyim/accounts.json")
	}
}
