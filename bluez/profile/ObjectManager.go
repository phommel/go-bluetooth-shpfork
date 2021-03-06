package profile

import (
	"git.enexoma.de/r/smartcontrol/libraries/go-bluetooth.git/bluez"
	"github.com/godbus/dbus"
)

// NewObjectManager create a new Device1 client
func NewObjectManager(name string, path string) *ObjectManager {
	om := new(ObjectManager)
	om.client = bluez.NewClient(
		&bluez.Config{
			Name:  name,
			Iface: "org.freedesktop.DBus.ObjectManager",
			Path:  path,
			Bus:   bluez.SystemBus,
		},
	)

	return om
}

// ObjectManager manges the list of all available objects
type ObjectManager struct {
	client *bluez.Client
}

// Close the connection
func (o *ObjectManager) Close() {
	o.client.Disconnect()
}

// GetManagedObjects return a list of all available objects registered
func (o *ObjectManager) GetManagedObjects() (map[dbus.ObjectPath]map[string]map[string]dbus.Variant, error) {
	var objs map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	err := o.client.Call("GetManagedObjects", 0).Store(&objs)
	return objs, err
}

//Register watch for signal events
func (o *ObjectManager) Register() (chan *dbus.Signal, error) {
	path := o.client.Config.Path
	iface := o.client.Config.Iface
	return o.client.Register(path, iface)
}

//Unregister watch for signal events
func (o *ObjectManager) Unregister(signal chan *dbus.Signal) error {
	path := o.client.Config.Path
	iface := o.client.Config.Iface
	return o.client.Unregister(path, iface, signal)
}
