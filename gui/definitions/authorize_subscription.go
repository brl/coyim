package definitions

func init() {
	add(`AuthorizeSubscription`, &defAuthorizeSubscription{})
}

type defAuthorizeSubscription struct{}

func (*defAuthorizeSubscription) String() string {
	return `<interface>
  <object class="GtkMessageDialog" id="dialog">
    <property name="title" translatable="yes">Subscription request</property>

    <property name="window-position">GTK_WIN_POS_CENTER</property>
    <property name="modal">true</property>
    <property name="message-type">GTK_MESSAGE_QUESTION</property>
    <property name="buttons">GTK_BUTTONS_YES_NO</property>
  </object>
</interface>
`
}
