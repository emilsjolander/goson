package main

import (
	"github.com/davecheney/profile"
	"github.com/emilsjolander/goson"
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
	defer profile.Start(profile.CPUProfile).Stop()

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
