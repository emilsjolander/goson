package goson

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Args map[string]interface{}

const TOKEN_COMMENT = 0     //token a comment
const TOKEN_OPEN_BRACE = 1  //token representing opening brace
const TOKEN_CLOSE_BRACE = 2 //token representing closing brace
const TOKEN_KEY = 3         //token representing a json key
const TOKEN_STRING = 4      //token representing a string literal
const TOKEN_FLOAT = 5       //token representing a float literal
const TOKEN_INT = 6         //token representing a int literal
const TOKEN_BOOL = 7        //token representing a bool literal
const TOKEN_INCLUDE = 8     //token representing a bool literal
const TOKEN_ALIAS = 9       //token representing an alias/new variable declaration
const TOKEN_LOOP = 10       //token representing a loop variable decleration
const TOKEN_ARGUMENT = 11   //token representing a argument from the args hash

func init() {
	RegisterTokenPattern(TOKEN_COMMENT, "\\/\\/.*[\\n\\r]?")     //one line comment
	RegisterTokenPattern(TOKEN_COMMENT, "\\/\\*[\\s\\S]*\\*\\/") //multi-line comment
	RegisterTokenPattern(TOKEN_OPEN_BRACE, "{")
	RegisterTokenPattern(TOKEN_CLOSE_BRACE, "}")
	RegisterTokenPattern(TOKEN_KEY, "[A-Za-z_]+ *:")
	RegisterTokenPattern(TOKEN_STRING, "\".*\"")
	RegisterTokenPattern(TOKEN_FLOAT, "[0-9]+\\.[0-9]")
	RegisterTokenPattern(TOKEN_INT, "[0-9]+")
	RegisterTokenPattern(TOKEN_BOOL, "true|false")
	RegisterTokenPattern(TOKEN_INCLUDE, "include\\( *[A-Za-z0-9_-]+ *, *[A-Za-z\\.]+ *\\)") //include(file_name, argument)
	RegisterTokenPattern(TOKEN_ALIAS, "[A-Za-z\\.]+ +as +[A-Za-z_]+")
	RegisterTokenPattern(TOKEN_LOOP, "[A-Za-z_]+ +in +[A-Za-z\\.]+")
	RegisterTokenPattern(TOKEN_ARGUMENT, "[A-Za-z\\.]+")
}

func Render(templateName string, args Args) (result []byte, err error) {
	template, err := ioutil.ReadFile(templateName + ".goson")

	//probably cannot find the template file
	if err != nil {
		return
	}

	//recover from any panics and return them are errors instead
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			default:
				err = errors.New(fmt.Sprint(r))
			case error:
				err = r
			}
		}
	}()

	lastPathSegmentStart := strings.LastIndex(templateName, "/")
	var workingDir string
	if lastPathSegmentStart >= 0 {
		workingDir = templateName[0 : lastPathSegmentStart+1]
	}

	tokens := Tokenize(template)
	p := &parser{workingDir: workingDir, tokens: tokens, args: args, result: []byte{'{'}}
	p.parse()
	result = append(p.result, '}')
	return
}
