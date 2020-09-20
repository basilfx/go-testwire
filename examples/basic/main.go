package main

import (
	"github.com/basilfx/go-testwire"
	"github.com/imroc/req"

	log "github.com/sirupsen/logrus"
)

func main() {
	NewHero("Jack")

	if !testwire.IsEnabled() {
		log.Fatalf("Application not compiled with the testwire tag enabled.")
		return
	}

	go testwire.Serve(9000)

	// This would normally be part of an automated testing suite.
	name, _ := req.Get("http://localhost:9000/probes/hero.name")
	log.Infof("Hero's name is '%s'.", name.String())

	health, _ := req.Get("http://localhost:9000/probes/hero.health")
	log.Infof("Hero's health is '%s'.", health.String())
}
