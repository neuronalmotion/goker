package main

import (
	"github.com/neuronalmotion/goker"
	"log"
	"net/http"
)

func main() {
	// database
	goker.HandleArgs()
	h := goker.HttpHandler()
	defer goker.DBClose()

	// start the engine!
	log.Printf("Server listen on address %s...", goker.GokerCtx.Cfg.Addr())
	err := http.ListenAndServe(goker.GokerCtx.Cfg.Addr(), h)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
