package definitions

func init() {
	add(`ConnectionFailureNotification`, &defConnectionFailureNotification{})
}

type defConnectionFailureNotification struct{}

func (*defConnectionFailureNotification) String() string {
	return `<interface>
  <object class="GtkInfoBar" id="infobar">
    <property name="message-type">GTK_MESSAGE_ERROR</property>
    <property name="show-close-button">false</property>

    <child internal-child="content_area">
      <object class="GtkBox" id="box">
        <property name="spacing">0</property>
        <property name="homogeneous">false</property>
        <property name="orientation">GTK_ORIENTATION_HORIZONTAL</property>
        <property name="hexpand">True</property>
        <property name="margin-start">0</property>
        <property name="margin-end">0</property>
        <child>
          <object class="GtkLabel" id="message">
            <property name="selectable">TRUE</property>
            <property name="ellipsize">PANGO_ELLIPSIZE_END</property>
            <property name="hexpand">True</property>
            <property name="wrap">True</property>
            <property name="margin-right">10</property>
          </object>
          <packing>
            <property name="expand">True</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkButton" id="button_more_info">
            <property name="hexpand">False</property>
            <property name="halign">end</property>
            <property name="relief">none</property>
            <property name="valign">center</property>
            <property name="border_width">0</property>
            <style>
              <class name="raised"/>
              <class name="close"/>
            </style>
            <signal name="clicked" handler="on_more_info_signal" />
            <child>
              <object class="GtkImage" id="more-info-btn">
                <property name="icon-name">info</property>
              </object>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkButton" id="button_close">
            <property name="can_focus">1</property>
            <property name="border_width">2</property>
            <property name="relief">none</property>
            <property name="valign">center</property>
            <style>
              <class name="raised"/>
              <class name="close"/>
            </style>
            <property name="hexpand">False</property>
            <property name="halign">end</property>
            <property name="valign">center</property>
            <signal name="clicked" handler="on_close_signal" />
            <child>
              <object class="GtkImage" id="close_image">
                <property name="visible">True</property>
                <property name="icon_name">window-close-symbolic</property>
              </object>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="position">2</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>
`
}
