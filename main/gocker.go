package main

import (
	"github.com/neuronalmotion/gocker"
	"log"
	"net/http"
)

func main() {
	// database
	gocker.HandleArgs()
	h := gocker.HttpHandler()
	defer gocker.DBClose()

	// start the engine!
	log.Printf("Server listen on address %s...", gocker.GockerCtx.Cfg.Addr())
	err := http.ListenAndServe(gocker.GockerCtx.Cfg.Addr(), h)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
