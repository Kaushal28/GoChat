package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var allUsers []string

//Subscribe for subscribing new user
func Subscribe(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	if !userAlreadyExists(user.Username) {
		allUsers = append(allUsers, user.Username)
		userJSON, err := json.Marshal(allUsers)
		if err != nil {
			log.Fatal("Can't encode to JSON ", err)
		} else {
			w.Write(userJSON)
		}
	} else {
		errorMsg, err := json.Marshal("User already exists.")
		if err != nil {
			log.Fatal("Can't encode to JSON ", err)
		}
		w.Write(errorMsg)
	}
}

func userAlreadyExists(uname string) bool {
	for _, username := range allUsers {
		if username == uname {
			return true
		}
	}
	return false
}
