Goson
=====
Goson is a small simple DSL for generating json from your Go datatypes. Supports both structs and maps as well as function/method values.

Installation
------------
Make sure to first install Go and setup all the necessary environment variables. Visit http://golang.org/doc/install for more info.
After that simply `go get github.com/emilsjolander/goson`.

Getting started
---------------
This is a short demonstration of the library in use. The next sections will go into detail about the API and the templating syntax.

All the code below can be found in the sample folder of the project.

Start with setting up a small Go program, something like this:
```go
package main

import (
	"fmt"
	"github.com/emilsjolander/goson"
)

type Repo struct {
	Name  string
	Url   string
	Stars int
	Forks int
}

type User struct {
	Name  string
	Repos []Repo
}

func main() {

	user := &User{
		Name: "Emil Sjölander",
		Repos: []Repo{
			Repo{
				Name:  "goson",
				Url:   "https://github.com/emilsjolander/goson",
				Stars: 0,
				Forks: 0,
			},
			Repo{
				Name:  "StickyListHeaders",
				Url:   "https://github.com/emilsjolander/StickyListHeaders",
				Stars: 722,
				Forks: 197,
			},
			Repo{
				Name:  "android-FlipView",
				Url:   "https://github.com/emilsjolander/android-FlipView",
				Stars: 157,
				Forks: 47,
			},
		},
	}

	result, err := goson.Render("user", goson.Args{"User": user})

	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}
```
The first thing we do in the above code is to import goson as well as the fmt packages. We also define 2 data types, User and Repo. These are the data types that we want to format into json. In the main function of our small sample we create an instance of User containing 3 instances of Repo. After this come the only public function of this library, the `Render()` function. `Render` takes two arguments, first the name of the template to render excluding the templates file type which should be `.goson`. The second argument to render is a map of argument that the tempate can make use of.

Let's take a look as user.goson to see how we define out json structure.
```text
user: {
	name: User.Name
	repos: Repo in User.Repos {
		name: Repo.Name
		url: Repo.Url
		stars: Repo.Stars
		forks: Repo.Forks
	}
}
```
The Above template starts by wrapping the fields within a "user" json object, next it writes the name of the user and than itterates through the repos printing each repos name, url, stars and forks. 
The resulting json is the following:
```json
{
    "user": {
        "name": "Emil Sjölander",
        "repos": [
            {
                "name": "goson",
                "url": "https://github.com/emilsjolander/goson",
                "stars": 0,
                "forks": 0
            },
            {
                "name": "StickyListHeaders",
                "url": "https://github.com/emilsjolander/StickyListHeaders",
                "stars": 722,
                "forks": 197
            },
            {
                "name": "android-FlipView",
                "url": "https://github.com/emilsjolander/android-FlipView",
                "stars": 157,
                "forks": 47
            }
        ]
    }
}
```
As you can see the result is automatically wrapped inside a json object. This is to follow standard restfull response formats.

API
---
As i hinted at during the getting started part of this readme, the API is very small. It consists of only one function and that is
```go
goson.Render(template string, args Args)
```
The template parameter should be the relative filepath to the template. So if you are executing main.go and your template is inside the templates folder you will want to pass `"templates/my_template"` to `Render()`. This will render the your data with the my_template.goson template which is located inside the template directory.

Args is just an alias for map[string]interface{} and accepts almost anything as an argument. Complex numbers and channels are the two common data types not currently supported.

Syntax
------
Goson is a fairly powerfull templating language with support for everything you could want (Open a pull request if i've missed anything).
`TODO`

Contributing
------------

Pull requests and issues are very welcome! 
If you want a fix to happen sooner than later i suggest you make a pull request. Please make pull request early, even before you are done with a feature/fix/enhancement. This way we can discuss and help each other out :)


License
-------

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
