package gui

import (
	"os"
	"strings"
)

type colorSet struct {
	rosterPeerBackground               string
	rosterPeerOfflineForeground        string
	rosterPeerOnlineForeground         string
	rosterGroupBackground              string
	rosterAccountOnlineBackground      string
	rosterAccountOfflineBackground     string
	conversationOutgoingUserForeground string
	conversationIncomingUserForeground string
	conversationOutgoingTextForeground string
	conversationIncomingTextForeground string
	conversationStatusTextForeground   string
}

var themeVariant string

func (u *gtkUI) isDarkThemeVariant() bool {
	if themeVariant != "" {
		return themeVariant == "dark"
	}
	themeVariant = "light"
	gtkTheme := os.Getenv("GTK_THEME")
	if gtkTheme != "" {
		toks := strings.Split(gtkTheme, ":")
		variant := toks[len(toks)-1:][0]
		if variant == "dark" {
			themeVariant = variant
			return true
		}
	}
	settings, err := g.gtk.SettingsGetDefault()
	if err != nil {
		panic(err)
	}
	prefDark, _ := settings.GetProperty("gtk-application-prefer-dark-theme")
	if prefDark == true {
		themeVariant = "dark"
		return true
	}
	return false
}

func (u *gtkUI) currentColorSet() colorSet {
	if u.isDarkThemeVariant() {
		return u.defaultDarkColorSet()
	}
	return u.defaultLightColorSet()
}

func (u *gtkUI) defaultLightColorSet() colorSet {
	return colorSet{
		rosterPeerBackground:               "#ffffff",
		rosterPeerOfflineForeground:        "#aaaaaa",
		rosterPeerOnlineForeground:         "#000000",
		rosterGroupBackground:              "#e9e7f3",
		rosterAccountOnlineBackground:      "#918caa",
		rosterAccountOfflineBackground:     "#d5d3de",
		conversationOutgoingUserForeground: "#3465a4",
		conversationIncomingUserForeground: "#a40000",
		conversationOutgoingTextForeground: "#555753",
		conversationIncomingTextForeground: "#000000",
		conversationStatusTextForeground:   "#4e9a06",
	}
}

func (u *gtkUI) defaultDarkColorSet() colorSet {
	return colorSet{
		rosterPeerBackground:               "#7f7f7f",
		rosterPeerOfflineForeground:        "#aaaaaa",
		rosterPeerOnlineForeground:         "#e5e5e5",
		rosterGroupBackground:              "#b8b6bf",
		rosterAccountOnlineBackground:      "#d5d3de",
		rosterAccountOfflineBackground:     "#918caa",
		conversationOutgoingUserForeground: "#3465a4",
		conversationIncomingUserForeground: "#a40000",
		conversationOutgoingTextForeground: "#f3f3f3",
		conversationIncomingTextForeground: "#7f7f7f",
		conversationStatusTextForeground:   "#4e9a06",
	}
}
