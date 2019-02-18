package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/time", showCurrentSystemTime)
	http.ListenAndServe(":8080", nil)
}

func showCurrentSystemTime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
