package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	json2 "github.com/tbocek/VSS-PROTO/json"
	"net/http"
	"time"
)

func main() {
	user := json2.User{"Thomas", "hsr12345", time.Time{}}
	jsonValue, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodGet, "http://10.0.2.16:3000/user", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	start := time.Now()
	for i := 0; i < 100; i++ {

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		err = json.NewDecoder(resp.Body).Decode(&user)
		if err != nil {
			panic(err)
		}
	}
	elapsed := time.Since(start)

	fmt.Printf("%+v, took %s\n", user, elapsed)
}
