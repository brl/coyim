package definitions

func init() {
	add(`ConfirmAccountRemoval`, &defConfirmAccountRemoval{})
}

type defConfirmAccountRemoval struct{}

func (*defConfirmAccountRemoval) String() string {
	return `<interface>
  <object class="GtkMessageDialog" id="RemoveAccount">
    <property name="title" translatable="yes">Confirm account removal</property>
    <property name="secondary-text" translatable="yes">Are you sure you want to remove this account?</property>
    <property name="window-position">GTK_WIN_POS_CENTER</property>
    <property name="modal">true</property>
    <property name="message-type">GTK_MESSAGE_QUESTION</property>
    <property name="buttons">GTK_BUTTONS_YES_NO</property>
  </object>
</interface>
`
}
