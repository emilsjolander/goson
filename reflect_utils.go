package goson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

//passes result of getArg() to encoding/json.Marshal()
func valueForKey(args Args, key []byte) []byte {
	arg := getArg(args, key)
	result, err := json.Marshal(arg)
	if err != nil {
		panic(err)
	}
	return result
}

//uses getArg() but adds validation that the value is indeed representable as a json object
func objectForKey(args Args, key []byte) interface{} {
	arg := getArg(args, key)
	t := reflect.TypeOf(arg)
	if isTypeObject(t) {
		return arg
	}
	panic(fmt.Sprintf("Argument error: Value was not of type struct/*struct/map[string], was type %s", t))
	return nil
}

//uses getArg() but adds validation that the value is indeed a collection of some kind
func collectionForKey(args Args, key []byte) Collection {
	arg := getArg(args, key)

	switch arg := arg.(type) {
	case Collection:
		return arg
	default:
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Array, reflect.Slice:
			return &reflectArrayWrapper{value: reflect.ValueOf(arg)}
		}
	}

	panic(fmt.Sprintf("Argument error: Value was not of type array/slice/goson.Collection, was type %s", reflect.TypeOf(arg)))
	return nil
}

//check of the type represents a json object (map[string], struct or *struct)
func isTypeObject(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Map:
		if t.Key().Kind() == reflect.String {
			return true
		}
	case reflect.Ptr:
		if t.Elem().Kind() == reflect.Struct {
			return true
		}
	case reflect.Struct:
		return true
	}
	return false
}

//get the value of a possible nested attribute inside the args map. nested attributes are represented by dot notation
func getArg(args Args, key []byte) interface{} {
	expParts := strings.Split(string(key), ".")
	rootValue, ok := args[expParts[0]]
	if !ok {
		panic(fmt.Sprintf("Argument error: %s not found", expParts[0]))
	}
	if rootValue == nil {
		return nil
	}

	value := reflect.ValueOf(rootValue)
	for i := 1; i < len(expParts); i++ {
		value, ok = getReflectValue(value, expParts[i])
		if !ok {
			panic(fmt.Sprintf("Argument error: %s not found in %s", expParts[i], expParts[i-1]))
		}
	}

	if value.Kind() == reflect.Func {
		value = getFuncSingleReturnValue(value)
	}

	return value.Interface()
}

//get the value with with the matching name inside v.
//This value can be a struct field, a method attached to a struct or a value in a map
func getReflectValue(v reflect.Value, valueName string) (value reflect.Value, ok bool) {

	// first check if input was a map and handle that.
	// otherwise input was a struct or pointer to a struct
	if v.Kind() == reflect.Map {
		value = v.MapIndex(reflect.ValueOf(valueName))
		if value.IsValid() {
			if value.Kind() == reflect.Func {
				value = getFuncSingleReturnValue(value)
			}
			ok = true
			return value, true
		}
	}

	value = reflect.Indirect(v).FieldByName(valueName)
	if value.IsValid() {
		if value.Kind() == reflect.Func {
			value = getFuncSingleReturnValue(value)
		}
		ok = true
		return
	}

	value = v.MethodByName(valueName)
	if value.IsValid() {
		value = getFuncSingleReturnValue(value)
		ok = true
		return
	}

	return
}

//take all of the field/methods/key-value pairs from val and add them as args. Valid input is struct, *struct and map[string]
func explodeIntoArgs(val interface{}) (args Args) {
	v := reflect.Indirect(reflect.ValueOf(val))
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Struct:
		args = Args{}
		for i := 0; i < t.NumField(); i++ {
			args[t.Field(i).Name] = v.Field(i).Interface()
		}
		for i := 0; i < t.NumMethod(); i++ {
			args[t.Method(i).Name] = v.Method(i).Interface()
		}
	case reflect.Map:
		args = Args{}
		if t.Key().Kind() == reflect.String {
			for _, key := range v.MapKeys() {
				args[key.String()] = v.MapIndex(key).Interface()
			}
		} else {
			panic("Maps used as arguments must have string keys")
		}
	default:
		panic(fmt.Sprintf("Variables must be of type map or struct/*struct to be used as arguments, was type %s", reflect.TypeOf(t)))
	}
	return
}

//validate that the reflect.Value represents a function with no arguments and a single return value.
//Return that value
func getFuncSingleReturnValue(fnc reflect.Value) reflect.Value {
	if fnc.Type().NumIn() != 0 {
		panic("Functions in template must be no arg functions")
	}
	if fnc.Type().NumOut() != 1 {
		panic("Functions in template must have exactly 1 return parameter")
	}
	if fnc.Type().Out(0).Kind() == reflect.Func {
		panic("Functions in template may not have a function return type")
	}
	return fnc.Call(nil)[0]
}
