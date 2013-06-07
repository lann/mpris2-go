package mpris2

import (
	"fmt"
	
	"github.com/lann/go-dbus"
)

// http://specifications.freedesktop.org/mpris-spec/latest/

const (
	dbusObject        = "org.freedesktop.DBus"
	dbusPath          = "/org/freedesktop/DBus"
	dbusInterface     = "org.freedesktop.DBus"
	propertyInterface = "org.freedesktop.DBus.Properties"
	listMethod        = "ListNames"
)

type Connection struct {
	*dbus.Connection
}

func Connect() (*Connection, error) {
	conn, err := dbus.Connect(dbus.SessionBus)
	if err != nil {
		return nil, err
	}

	if err = conn.Authenticate(); err != nil {
		return nil, err
	}

	return &Connection{conn}, nil
}

type Object struct {
	*dbus.Object
	conn *Connection
}

func (conn *Connection) getObject(name, path string) (*Object, error) {
	obj := conn.Object(name, path)
	if obj == nil {
		return nil, fmt.Errorf("couldn't find %s %s", name, path)
	}

	return &Object{Object: obj, conn: conn}, nil
}

type Interface struct {
	*dbus.Interface
    obj  *Object
	name string
}

func (obj *Object) getInterface(name string) (*Interface, error) {
	iface := obj.Interface(name)
	if iface == nil {
		return nil, fmt.Errorf("couldn't get interface %s", name)
	}
	return &Interface{Interface: iface, obj: obj, name: name}, nil
}

func (iface *Interface) call(methodName string, args ...interface{}) ([]interface{}, error) {
	method, err := iface.Method(methodName)
	if err != nil {
		return nil, err
	}

	return iface.obj.conn.Call(method, args...)
}

func (iface *Interface) callVoid(methodName string, args ...interface{}) (error) {
	_, err := iface.call(methodName, args...)
	return err
}

func (iface *Interface) getProp(name string) ([]interface{}, error) {
	propIface, err := iface.obj.getInterface(propertyInterface)
	if err != nil {
		return nil, err
	}
	return propIface.call("Get", iface.name, name)
}

func (iface *Interface) getValue(name string) (interface{}, error) {
	out, err := iface.getProp(name)
	if err != nil {
		return nil, err
	}
	if len(out) != 1 {
		return nil, fmt.Errorf("unexpected return value %v", out)
	}

	return out[0], nil
}

func (iface *Interface) getBool(name string) (result bool, err error) {
	out, err := iface.getValue(name)
	if err != nil {
		return
	}
	
	result, ok := out.(bool)
	if !ok {
		err = fmt.Errorf("could not convert result to boolean")
	}
	return
}

func (iface *Interface) getString(name string) (result string, err error) {
	out, err := iface.getValue(name)
	if err != nil {
		return
	}
	
	result, ok := out.(string)
	if !ok {
		err = fmt.Errorf("could not convert result to string")
	}
	return
}

func (iface *Interface) getArray(name string) (result []interface{}, err error) {
	out, err := iface.getValue(name)
	if err != nil {
		return
	}
	
	result, ok := out.([]interface{})
	if !ok {
		err = fmt.Errorf("could not convert result to array")
	}
	return
}

func (iface *Interface) getStringArray(name string) (results []string, err error) {
	items, err := iface.getArray(name)
	if err != nil {
		return
	}

	for _, item := range items {
		result, ok := item.(string)
		if !ok {
			err = fmt.Errorf("unexpected return value %v", item)
			return
		}
		results = append(results, result)
	}
	return
}

