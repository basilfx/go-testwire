package main

import (
	"math/rand"
	"strconv"

	"github.com/basilfx/go-testwire"
)

type Hero struct {
	name   string
	health int
}

func NewHero(name string) *Hero {
	h := Hero{
		name:   name,
		health: rand.Intn(100),
	}

	// This code is only enabled if compiled with the testwire tag set.
	testwire.AddProbe("hero.name", func() string {
		return h.name
	})
	testwire.AddProbe("hero.health", func() string {
		return strconv.Itoa(h.health)
	})

	return &h
}
