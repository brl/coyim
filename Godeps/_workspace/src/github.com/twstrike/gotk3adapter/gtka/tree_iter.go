package gtka

import (
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gtki"
)

type treeIter struct {
	*gtk.TreeIter
}

func wrapTreeIterSimple(v *gtk.TreeIter) *treeIter {
	if v == nil {
		return nil
	}
	return &treeIter{v}
}

func wrapTreeIter(v *gtk.TreeIter, e error) (*treeIter, error) {
	return wrapTreeIterSimple(v), e
}

func unwrapTreeIter(v gtki.TreeIter) *gtk.TreeIter {
	if v == nil {
		return nil
	}
	return v.(*treeIter).TreeIter
}
