package main

import (
	"fmt"
	"github.com/davecheney/profile"
	"github.com/emilsjolander/goson"
	"os"
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

func main() {
	profileKind := os.Args[1]
	switch profileKind {
	case "cpu":
		defer profile.Start(profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile).Stop()
	default:
		fmt.Println("only cpu, mem and block are valid profile arguments")
		return
	}

	for i := 0; i < 100000; i++ {
		user := &User{
			Name: "Emil Sjölander",
			Repos: []Repo{
				Repo{
					Name:  "goson",
					URL:   "https://github.com/emilsjolander/goson",
					Stars: 0,
					Forks: 0,
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

		goson.Render("user", goson.Args{"User": user})
	}
}
