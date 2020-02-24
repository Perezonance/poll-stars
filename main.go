package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", VoteHandler)
	http.ListenAndServe(":8080", nil)
}
