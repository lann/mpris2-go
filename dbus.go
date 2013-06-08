package mpris2

import (
	"fmt"
	"reflect"
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

func (obj *Object) getValue(name string, kind reflect.Kind) (value reflect.Value, err error) {
	var data dbus.Variant
	err = obj.getProp(name).Store(&data)
	if err == nil {
		value = reflect.ValueOf(data.Value())
		if value.Kind() != kind {
			err = fmt.Errorf("wrong kind %s, expected %s", value.Kind(), kind)
		}
	}
	return
}

func (obj *Object) getBool(name string) (result bool, err error) {
	value, err := obj.getValue(name, reflect.Bool)
	if err == nil {
		result = value.Bool()
	}
	return
}

func (obj *Object) getString(name string) (result string, err error) {
	value, err := obj.getValue(name, reflect.String)
	if err == nil {
		result = value.String()
	}
	return
}

func (obj *Object) getStringArray(name string) (results []string, err error) {
	value, err := obj.getValue(name, reflect.Slice)
	if err == nil {
		var ok bool
		results, ok = value.Interface().([]string)
		if !ok {
			err = fmt.Errorf("unexpected type %s", value.Type())
		}
	}
	return
}

func (obj *Object) getStringMap(name string) (results map[string]interface{}, err error) {
	value, err := obj.getValue(name, reflect.Map)
	if err == nil {
		data, ok := value.Interface().(map[string]dbus.Variant)
		if !ok {
			err = fmt.Errorf("unexpected type %s", value.Type())
		}

		results = make(map[string]interface{})
		for k, v := range data {
			results[k] = v.Value()
		}
	}
	return
}