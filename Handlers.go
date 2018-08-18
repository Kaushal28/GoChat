package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Subscribe for subscribing new user
func Subscribe(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Username)
	fmt.Fprintf(w, "Subscribed!")
}
