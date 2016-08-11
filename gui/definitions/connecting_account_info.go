package definitions

func init() {
	add(`ConnectingAccountInfo`, &defConnectingAccountInfo{})
}

type defConnectingAccountInfo struct{}

func (*defConnectingAccountInfo) String() string {
	return `<interface>
  <object class="GtkInfoBar" id="infobar">
    <property name="message-type">GTK_MESSAGE_INFO</property>
    <child internal-child="content_area">
      <object class="GtkBox" id="box">
        <property name="homogeneous">false</property>
        <property name="orientation">GTK_ORIENTATION_HORIZONTAL</property>
        <child>
          <object class="GtkLabel" id="message">
            <property name="ellipsize">PANGO_ELLIPSIZE_END</property>
            <property name="wrap">true</property>
          </object>
        </child>
        <child>
          <object class="GtkSpinner" id="spinner">
            <property name="active">true</property>
            <property name="halign">GTK_ALIGN_END</property>
            <property name="hexpand">true</property>
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`
}
