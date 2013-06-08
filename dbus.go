package mpris2

import (
	"fmt"
	"github.com/guelfey/go.dbus"
)

type Conn struct {
	*dbus.Conn
}

func Connect() (*Conn, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	return &Conn{conn}, nil
}

func (conn *Conn) ListNames() (names []string, err error) {
	err = conn.BusObject().Call("org.freedesktop.DBus.ListNames", 0).Store(&names)
	return
}

type Object struct {
	*dbus.Object
	interfaceName string
}

func (conn *Conn) getObject(name, path, interfaceName string) *Object {
	return &Object{conn.Object(name, dbus.ObjectPath(path)), interfaceName}
}

func (obj *Object) call(method string, args ...interface{}) *dbus.Call {
	return obj.Call(fmt.Sprintf("%s.%s", obj.interfaceName, method), 0, args...)
}

func (obj *Object) callVoid(method string, args ...interface{}) error {
	return obj.call(method, args...).Store()
}
	
func (obj *Object) getProp(name string) *dbus.Call {
	return obj.Call("org.freedesktop.DBus.Properties.Get", 0, obj.interfaceName, name)
}

func (obj *Object) getBool(name string) (result bool, err error) {
	err = obj.getProp(name).Store(&result)
	return
}

func (obj *Object) getString(name string) (result string, err error) {
	err = obj.getProp(name).Store(&result)
	return
}

func (obj *Object) getStringArray(name string) (results []string, err error) {
	err = obj.getProp(name).Store(&results)
	return
}

func (obj *Object) getStringMap(name string) (results map[string]interface{}, err error) {
	var data dbus.Variant
	err = obj.getProp(name).Store(&data)
	if err != nil {
		return
	}
	if sig := data.Signature().String(); sig != "a{sv}" {
		err = fmt.Errorf("Unexpected result signature %s", sig)
		return
	}

	results = make(map[string]interface{})
	
	for k, v := range data.Value().(map[string]dbus.Variant) {
		results[k] = v.Value()
	}
	
	return
}