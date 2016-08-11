package definitions

func init() {
	add(`BadUsernameNotification`, &defBadUsernameNotification{})
}

type defBadUsernameNotification struct{}

func (*defBadUsernameNotification) String() string {
	return `<interface>
  <object class="GtkInfoBar" id="infobar">
    <property name="message-type">GTK_MESSAGE_WARNING</property>

    <child internal-child="content_area">
      <object class="GtkBox" id="box">
        <property name="homogeneous">false</property>
        <property name="orientation">GTK_ORIENTATION_HORIZONTAL</property>
        <child>
          <object class="GtkLabel" id="message">
            <property name="wrap">true</property>
          </object>
        </child>
      </object>
    </child>

  </object>
</interface>
`
}
