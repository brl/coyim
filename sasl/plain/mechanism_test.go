package plain

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/glib_mock"
	. "github.com/twstrike/coyim/Godeps/_workspace/src/gopkg.in/check.v1"
	"github.com/twstrike/coyim/i18n"
	"github.com/twstrike/coyim/sasl"
)

func Test(t *testing.T) { TestingT(t) }

func init() {
	log.SetOutput(ioutil.Discard)
	i18n.InitLocalization(&glib_mock.Mock{})
}

type SASLPlain struct{}

var _ = Suite(&SASLPlain{})

func (s *SASLPlain) Test(c *C) {
	expected := sasl.Token("\x00foo\x00bar")

	client := Mechanism.NewClient()
	c.Check(client.NeedsMore(), Equals, true)

	client.SetProperty(sasl.AuthID, "foo")
	client.SetProperty(sasl.Password, "bar")

	t, err := client.Step(nil)

	c.Check(err, IsNil)
	c.Check(client.NeedsMore(), Equals, true)
	c.Check(t, DeepEquals, expected)

	t, err = client.Step(nil)
	c.Check(err, IsNil)
	c.Check(client.NeedsMore(), Equals, false)
}
