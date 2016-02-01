package gui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/coyim/config"
	"github.com/twstrike/coyim/i18n"
)

func (u *gtkUI) showFingerprintsForPeer(jid string, account *account) {
	builder := builderForDefinition("PeerFingerprints")
	dialog := getObjIgnoringErrors(builder, "dialog").(*gtk.Dialog)
	info := getObjIgnoringErrors(builder, "information").(*gtk.Label)
	grid := getObjIgnoringErrors(builder, "grid").(*gtk.Grid)

	info.SetSelectable(true)

	fprs := []*config.Fingerprint{}
	p, ok := account.session.GetConfig().GetPeer(jid)
	if ok {
		fprs = p.Fingerprints
	}

	if len(fprs) == 0 {
		info.SetText(fmt.Sprintf(i18n.Local("There are no known fingerprints for %s"), jid))
	} else {
		info.SetText(fmt.Sprintf(i18n.Local("These are the fingerprints known for %s:"), jid))
	}

	for ix, fpr := range fprs {
		flabel, _ := gtk.LabelNew(config.FormatFingerprint(fpr.Fingerprint))
		flabel.SetSelectable(true)
		trusted := i18n.Local("not trusted")
		if fpr.Trusted {
			trusted = i18n.Local("trusted")
		}

		ftrusted, _ := gtk.LabelNew(trusted)
		ftrusted.SetSelectable(true)

		grid.Attach(flabel, 0, ix, 1, 1)
		grid.Attach(ftrusted, 1, ix, 1, 1)
	}

	builder.ConnectSignals(map[string]interface{}{
		"on_close_signal": func() {
			dialog.Destroy()
		},
	})

	dialog.SetTransientFor(u.window)
	dialog.ShowAll()
}
