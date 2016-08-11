package ui

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/glib_mock"
	"github.com/twstrike/coyim/i18n"

	. "github.com/twstrike/coyim/Godeps/_workspace/src/gopkg.in/check.v1"
)

var escapingTests = []string{
	"",
	"foo",
	"foo\\",
	"foo\\x01",
	"العربية",
}

func Test(t *testing.T) { TestingT(t) }

func init() {
	log.SetOutput(ioutil.Discard)
	i18n.InitLocalization(&glib_mock.Mock{})
}

type UISuite struct{}

var _ = Suite(&UISuite{})

func (s *UISuite) TestEscaping(t *C) {
	for _, test := range escapingTests {
		escaped := EscapeNonASCII(test)
		unescaped, err := UnescapeNonASCII(escaped)
		if err != nil {
			t.Errorf("Error unescaping '%s' (from '%s')", escaped, test)
			continue
		}
		if unescaped != test {
			t.Errorf("Unescaping didn't return the original value: '%s' -> '%s' -> '%s'", test, escaped, unescaped)
		}
	}
}

func (s *UISuite) TestHTMLStripping(t *C) {
	raw := []byte("<hr>This is some <font color='green'>html</font><br />.")
	exp := []byte("This is some html.")
	res := StripHTML(raw)

	t.Check(res, DeepEquals, exp)
}

func (s *UISuite) Test_StripSomeHTML(t *C) {
	raw := []byte("<p>This is <walloftext>some</walloftext> <FONT color='green'>html</font><br />.")
	exp := "This is <walloftext>some</walloftext> html."
	res := StripSomeHTML(raw)

	t.Check(string(res), DeepEquals, exp)
}
