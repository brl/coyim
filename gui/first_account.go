package gui

import "github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gtki"

func (u *gtkUI) showFirstAccountWindow() {
	result := make(chan func())
	var cleanup func()

	doInUIThread(func() {
		b := newBuilder("FirstAccountDialog")
		dialog := b.getObj("dialog").(gtki.Dialog)
		dialog.SetTransientFor(u.window)
		cleanup = dialog.Destroy

		b.ConnectSignals(map[string]interface{}{
			"on_import_signal": func() {
				result <- u.runImporter
			},
			"on_register_signal": func() {
				result <- u.showServerSelectionWindow
			},
			"on_existing_signal": func() {
				result <- u.showAddAccountWindow
			},
			"on_cancel_signal": func() {
				result <- func() {}
			},
		})

		dialog.ShowAll()
	})

	tp := <-result
	doInUIThread(cleanup)
	doInUIThread(tp)
}
