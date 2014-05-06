package main

import (
	"log"
	"net/http"
	"github.com/neuronalmotion/gocker"
)

func main() {
    // database
    gocker.InitDefaultDatabaseData()
	defer gocker.DBClose()

    h := gocker.HttpHandler()

    // start the engine!
	log.Println("Server listen on port 8000...")
    err := http.ListenAndServe(":8000", h)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

}
