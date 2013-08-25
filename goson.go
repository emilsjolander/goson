package goson

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Args map[string]interface{}

const (
	TOKEN_COMMENT     = iota //token a comment
	TOKEN_OPEN_BRACE         //token representing opening brace
	TOKEN_CLOSE_BRACE        //token representing closing brace
	TOKEN_KEY                //token representing a json key
	TOKEN_STRING             //token representing a string literal
	TOKEN_FLOAT              //token representing a float literal
	TOKEN_INT                //token representing a int literal
	TOKEN_BOOL               //token representing a bool literal
	TOKEN_INCLUDE            //token representing a bool literal
	TOKEN_ALIAS              //token representing an alias/new variable declaration
	TOKEN_LOOP               //token representing a loop variable decleration
	TOKEN_ARGUMENT           //token representing a argument from the args hash
)

var tokenCache = make(map[string][]token)

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

	tokens, exists := tokenCache[templateName]

	if !exists {
		var template []byte
		template, err = ioutil.ReadFile(templateName + ".goson")

		//probably cannot find the template file
		if err != nil {
			return
		}

		tokens = Tokenize(template)
		tokenCache[templateName] = tokens
	}

	lastPathSegmentStart := strings.LastIndex(templateName, "/")
	var workingDir string
	if lastPathSegmentStart >= 0 {
		workingDir = templateName[0 : lastPathSegmentStart+1]
	}

	p := &parser{workingDir: workingDir, tokens: tokens, args: args, result: []byte{'{'}}
	p.parse()
	result = append(p.result, '}')
	return
}
