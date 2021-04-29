package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	bf "github.com/tbocek/VSS-PROTO/bloomfilter"
	"github.com/willf/bloom"
	"log"
	"math/rand"
	"net/http"
)

var a []int64

func main() {
	for i := 0; i < 1000; i++ {
		a[i] = rand.Int63n(2000)
	}
	var router = mux.NewRouter()
	router.HandleFunc("/nr", user).Methods("GET")
	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func user(w http.ResponseWriter, r *http.Request) {

	filter := bloom.New(100, 5)
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	filter.UnmarshalJSON(buf.Bytes())

	var ret []int64
	c := 0
	for i := 0; i < 1000; i++ {
		n1 := make([]byte, 8)
		binary.BigEndian.PutUint64(n1, uint64(a[i]))
		if !filter.Test(n1) {
			ret[c] = a[i]
			c++
		}
	}

	retJson := bf.Numbers{ret}
	err := json.NewEncoder(w).Encode(retJson)
	if err != nil {
		panic(err)
	}
}
