package tests

import (
	"github.com/emilsjolander/goson"
	"testing"
)

// Test rendering a string passed to Args.
func TestString(t *testing.T) {
	result, err := goson.Render("templates/string", goson.Args{"string": "a string"})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"string\":\"a string\"}" {
		t.Error("json did not match")
	}
}

// Test rendering an int passed to Args. Test all sizes of ints
func TestInt(t *testing.T) {
	result, err := goson.Render("templates/int", goson.Args{"int": int(-1), "int8": int8(2), "int16": int16(3), "int32": int32(-4), "int64": int64(5)})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"int\":-1,\"int8\":2,\"int16\":3,\"int32\":-4,\"int64\":5}" {
		t.Error("json did not match")
	}
}

// Test rendering an uint passed to Args. Test all sizes of uints
func TestUint(t *testing.T) {
	result, err := goson.Render("templates/uint", goson.Args{"uint": uint(10), "uint8": uint8(20), "uint16": uint16(30), "uint32": uint32(40), "uint64": uint64(50)})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"uint\":10,\"uint8\":20,\"uint16\":30,\"uint32\":40,\"uint64\":50}" {
		t.Error("json did not match")
	}
}

// Test rendering a float passed to Args. Test all sizes of floats
func TestFloat(t *testing.T) {
	result, err := goson.Render("templates/float", goson.Args{"float32": float32(32.32), "float64": float64(64.64)})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"float32\":32.32,\"float64\":64.64}" {
		t.Error("json did not match")
	}
}

// Test rendering a bool passed to Args.
func TestBool(t *testing.T) {
	result, err := goson.Render("templates/bool", goson.Args{"bool": true})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"bool\":true}" {
		t.Error("json did not match")
	}
}

// Test rendering a struct passed to Args.
func TestStruct(t *testing.T) {
	myStruct := struct{ MyString string }{MyString: "MyString"}
	result, err := goson.Render("templates/struct", goson.Args{"MyStruct": myStruct})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"my_struct\":{\"my_string\":\"MyString\"}}" {
		t.Error("json did not match")
	}
}

// Test rendering a pointer to a struct passed to Args.
func TestStructPtr(t *testing.T) {
	myStruct := struct{ MyString string }{MyString: "MyString"}
	result, err := goson.Render("templates/struct", goson.Args{"MyStruct": &myStruct})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"my_struct\":{\"my_string\":\"MyString\"}}" {
		t.Error("json did not match")
	}
}

// Test rendering a map passed to Args.
func TestMap(t *testing.T) {
	myMap := map[string]string{"key1": "key 1!", "key2": "key 2!", "key3": "key 3!"}
	result, err := goson.Render("templates/map", goson.Args{"MyMap": myMap})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"my_map\":{\"key1\":\"key 1!\",\"key2\":\"key 2!\",\"key3\":\"key 3!\"}}" {
		t.Error("json did not match")
	}
}

// Test rendering a function passed to Args.
func TestFunction(t *testing.T) {
	result, err := goson.Render("templates/function", goson.Args{"MyFunc": func() int { return 1337 }})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"func_result\":1337}" {
		t.Error("json did not match")
	}
}

//used for TestMethod test
type MethodTester struct{}

//used for TestMethod test
func (mt *MethodTester) TestMethod() string { return "method test" }

// Test rendering a method attached to a struct passed to Args.
func TestMethod(t *testing.T) {
	result, err := goson.Render("templates/method", goson.Args{"MyStruct": new(MethodTester)})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"my_struct\":{\"method\":\"method test\"}}" {
		t.Error("json did not match")
	}
}

// Test rendering a slice of structs.
func TestSliceOfStructs(t *testing.T) {
	type Object struct{ StringField string }
	objects := []Object{Object{"hej"}, Object{"jag"}, Object{"heter"}, Object{"emil"}}
	result, err := goson.Render("templates/slice_of_objects", goson.Args{"Objects": objects})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"objects\":[{\"string\":\"hej\"},{\"string\":\"jag\"},{\"string\":\"heter\"},{\"string\":\"emil\"}]}" {
		t.Error("json did not match")
	}
}

// Test rendering a slice of maps.
func TestSliceOfMaps(t *testing.T) {
	type Object map[string]string
	objects := []Object{Object{"StringField": "hej"}, Object{"StringField": "jag"}, Object{"StringField": "heter"}, Object{"StringField": "emil"}}
	result, err := goson.Render("templates/slice_of_objects", goson.Args{"Objects": objects})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"objects\":[{\"string\":\"hej\"},{\"string\":\"jag\"},{\"string\":\"heter\"},{\"string\":\"emil\"}]}" {
		t.Error("json did not match")
	}
}

// Test rendering a slice of single return value functions.
func TestSliceOfFunctions(t *testing.T) {
	i := 1
	intFunc := func() int {
		defer func() { i *= 2 }()
		return i
	}
	ints := []func() int{intFunc, intFunc, intFunc, intFunc}
	result, err := goson.Render("templates/slice_of_ints", goson.Args{"Ints": ints})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"ints\":[{\"int\":1},{\"int\":2},{\"int\":4},{\"int\":8}]}" {
		t.Error("json did not match")
	}
}

// Test rendering a slice of primites where the primitives are looped over and wrapped in json objects.
func TestSliceOfPrimitives(t *testing.T) {
	ints := []int{1, 2, 4, 8}
	result, err := goson.Render("templates/slice_of_ints", goson.Args{"Ints": ints})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"ints\":[{\"int\":1},{\"int\":2},{\"int\":4},{\"int\":8}]}" {
		t.Error("json did not match")
	}
}

// Test rendering a slice of primites passed as a json argument, should not be wrapped in json objects.
func TestSliceOfAnonymousPrimitives(t *testing.T) {
	person := struct {
		Name      string
		Nicknames []string
	}{"Emil", []string{"mardox", "odjuret", "kodarn"}}
	result, err := goson.Render("templates/slice_of_anonymous_strings", goson.Args{"Person": person})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"person\":{\"name\":\"Emil\",\"nicknames\":[\"mardox\",\"odjuret\",\"kodarn\"]}}" {
		t.Error("json did not match")
	}
}

type CollectionTester struct {
	ints []int
}

func (ct *CollectionTester) Add(item int) {
	ct.ints = append(ct.ints, item)
}

func (ct *CollectionTester) Get(index int) interface{} {
	return ct.ints[index]
}

func (ct *CollectionTester) Len() int {
	return len(ct.ints)
}

// Test rendering a collection.
func TestCollection(t *testing.T) {
	ints := new(CollectionTester)
	ints.Add(1)
	ints.Add(1)
	ints.Add(2)
	ints.Add(3)
	ints.Add(5)
	ints.Add(8)
	result, err := goson.Render("templates/slice_of_ints", goson.Args{"Ints": ints})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"ints\":[{\"int\":1},{\"int\":1},{\"int\":2},{\"int\":3},{\"int\":5},{\"int\":8}]}" {
		t.Error("json did not match")
	}
}
