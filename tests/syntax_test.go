package tests

import (
	"github.com/emilsjolander/goson"
	"testing"
)

// Test rendering of constants in the template. string, int, float, bool and object literals
func TestConstants(t *testing.T) {
	result, err := goson.Render("templates/constants", goson.Args{})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"object\":{\"string\":\"hej\",\"int\":123,\"float\":1.23,\"bool\":true}}" {
		t.Error("json did not match")
	}
}

// Test commenting out lines with both single and multiline syntax
func TestComments(t *testing.T) {
	result, err := goson.Render("templates/comments", goson.Args{})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"first\":1,\"third\":3,\"fifth\":5}" {
		t.Error("json did not match")
	}
}

// Test aliasing a struct
func TestAlias(t *testing.T) {
	type InnerStruct struct{ MyString string }
	type MyStruct struct{ MyInnerStruct InnerStruct }
	myStruct := MyStruct{InnerStruct{"hej!"}}
	result, err := goson.Render("templates/alias", goson.Args{"MyStruct": myStruct})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"object\":{\"string\":\"hej!\"}}" {
		t.Error("json did not match")
	}
}

// Test looping over a slice
func TestLoop(t *testing.T) {
	type Item struct{ Id int }
	items := []Item{Item{1}, Item{12}, Item{123}}
	result, err := goson.Render("templates/loop", goson.Args{"Items": items})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"items\":[{\"id\":1},{\"id\":12},{\"id\":123}]}" {
		t.Error("json did not match ")
	}
}

// Test including a template with arguments
func TestInclude(t *testing.T) {
	type MyStruct struct{ MyString string }
	myStruct := MyStruct{"hej!"}
	result, err := goson.Render("templates/include", goson.Args{"MyStruct": myStruct})
	if err != nil {
		t.Error(err)
	} else if string(result) != "{\"a_string\":\"rendering a include\",\"my_included_string\":\"hej!\"}" {
		t.Error("json did not match")
	}
}
