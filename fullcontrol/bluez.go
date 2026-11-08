package fullcontrol

import (
	"github.com/godbus/dbus/v5"
)

var conn *dbus.Conn

func InitBluetooth() {
	var err error
	conn, err = dbus.ConnectSystemBus()
	if err != nil {
		panic(err)
	}
}

func CloseDBus() error {
    if conn != nil {
        return conn.Close()
    }
    return nil
}

func GetAdapterAlias() (string, error) {
	obj := conn.Object(
        "org.bluez",
        "/org/bluez/hci0",
    )

	value, err := obj.GetProperty("org.bluez.Adapter1.Alias")
    if err != nil {
        return "", err
    }

	alias := value.Value().(string)
	return alias, nil
}

func SetAdapterAlias(alias string) (error) {
	obj := conn.Object(
        "org.bluez",
        "/org/bluez/hci0",
    )

	res := obj.Call(
		"org.freedesktop.DBus.Properties.Set",
		0,
		"org.bluez.Adapter1",
		"Alias",
		dbus.MakeVariant(alias),
	)

	return res.Err
}
