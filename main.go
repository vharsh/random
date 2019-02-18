package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", showCurrentSystemTime)
	http.ListenAndServe(":8080", nil)
}

func showCurrentSystemTime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
	w.Header().Set("X-API-Time", time.Now().String())
	fmt.Println("%#v", w)
}
