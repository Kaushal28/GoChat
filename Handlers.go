package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var allUsers []string

//Create a hashmap to store the messages of offline users
var messageQ = make(map[string][]Message)

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
		messageJSON, err := json.Marshal(messageQ[user.Username])
		if err != nil {
			log.Fatal("Can't encode to JSON ", err)
		} else {
			w.Write(messageJSON)
			//remove from map once delivered
			delete(messageQ, user.Username)
		}
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

//SendMsg for sending messages
func SendMsg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var message Message
	err := decoder.Decode(&message)
	if err != nil {
		panic(err)
	} else {
		messageQ[message.To] = append(messageQ[message.To], message)
	}
}
