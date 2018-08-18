package main

import (
	"fmt"
	"net/http"
)

func Subscribe(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Subscribed!")
}