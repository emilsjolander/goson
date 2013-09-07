package tests

import (
	"encoding/json"
	"github.com/emilsjolander/goson"
	"testing"
)

type Repo struct {
	Name  string
	URL   string
	Stars int
	Forks int
}

type User struct {
	Name  string
	Repos []Repo
}

var user = &User{
	Name: "Emil Sjölander",
	Repos: []Repo{
		Repo{
			Name:  "goson",
			URL:   "https://github.com/emilsjolander/goson",
			Stars: 71,
			Forks: 1,
		},
		Repo{
			Name:  "StickyListHeaders",
			URL:   "https://github.com/emilsjolander/StickyListHeaders",
			Stars: 722,
			Forks: 197,
		},
		Repo{
			Name:  "android-FlipView",
			URL:   "https://github.com/emilsjolander/android-FlipView",
			Stars: 157,
			Forks: 47,
		},
	},
}

func BenchmarkGosonSerialization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goson.Render("templates/user", goson.Args{"User": user})
	}
}

func BenchmarkStdlibSerialization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(user)
	}
}
