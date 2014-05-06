package main

import (
    "encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/neuronalmotion/gocker"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>hello ! </h1>\n")
}

func UsersListHandler(w http.ResponseWriter, req *http.Request) {
    var users []gocker.User
    db := gocker.DB

	db.Find(&users)
    js, err := json.Marshal(users)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func UserGetHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id := vars["id"]
    var user gocker.User
    query := gocker.DB.Where("id = ?", id).First(&user)
    if query.Error == gorm.RecordNotFound {
        http.Error(w, query.Error.Error(), http.StatusNotFound)
        return
    }

    js, err := json.Marshal(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func main() {
    // database
    gocker.InitDefaultDatabaseData()
	defer gocker.DBClose()

    // url routing
	r := mux.NewRouter()
	r.HandleFunc("/", HelloServer)
	r.HandleFunc("/users", UsersListHandler).Methods("GET")
    r.HandleFunc("/users/{id:[0-9]+}", UserGetHandler).Methods("GET")
	http.Handle("/", r)

    // start the engine!
	log.Println("Server listen on port 8000...")
	http.ListenAndServe(":8000", nil)

}
