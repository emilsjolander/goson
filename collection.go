package goson

import (
	"reflect"
)

type Collection interface {
	Get(index int) interface{}
	Len() int
}

//data structure to make a reflect.Value representing a arrya/slice conform to the collection interface
type reflectArrayWrapper struct {
	value reflect.Value
}

func (c *reflectArrayWrapper) Get(index int) interface{} {
	return c.value.Index(index).Interface()
}

func (c *reflectArrayWrapper) Len() int {
	return c.value.Len()
}
