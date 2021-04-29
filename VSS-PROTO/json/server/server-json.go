package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	json2 "github.com/tbocek/VSS-PROTO/json"
	"log"
	"net/http"
	"time"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/user", user).Methods("GET")
	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func user(w http.ResponseWriter, r *http.Request) {
	user := json2.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	user.Created = time.Now().UTC()
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		panic(err)
	}
}
