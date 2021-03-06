package gtki

import (
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gdki"
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/glibi"
)

type Gtk interface {
	AboutDialogNew() (AboutDialog, error)
	AccelGroupNew() (AccelGroup, error)
	AcceleratorParse(string) (uint, gdki.ModifierType)
	AddProviderForScreen(gdki.Screen, StyleProvider, uint)
	ApplicationNew(string, glibi.ApplicationFlags) (Application, error)
	BuilderNew() (Builder, error)
	BuilderNewFromResource(string) (Builder, error)
	CellRendererTextNew() (CellRendererText, error)
	CheckButtonNewWithMnemonic(string) (CheckButton, error)
	CheckMenuItemNewWithMnemonic(string) (CheckMenuItem, error)
	CssProviderNew() (CssProvider, error)
	EntryNew() (Entry, error)
	FileChooserDialogNewWith2Buttons(string, Window, FileChooserAction, string, ResponseType, string, ResponseType) (FileChooserDialog, error)
	Init(*[]string)
	LabelNew(string) (Label, error)
	ListStoreNew(...glibi.Type) (ListStore, error)
	MenuItemNew() (MenuItem, error)
	MenuItemNewWithLabel(string) (MenuItem, error)
	MenuItemNewWithMnemonic(string) (MenuItem, error)
	MenuNew() (Menu, error)
	SeparatorMenuItemNew() (SeparatorMenuItem, error)
	TextBufferNew(TextTagTable) (TextBuffer, error)
	TextTagNew(string) (TextTag, error)
	TextTagTableNew() (TextTagTable, error)
	TextViewNew() (TextView, error)
	TreePathNew() TreePath
	WindowSetDefaultIcon(gdki.Pixbuf)
	SettingsGetDefault() (Settings, error)
}

func AssertGtk(_ Gtk) {}
