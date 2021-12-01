package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("E:\\06-go")))
	log.Fatal(http.ListenAndServe(":3030", nil))
}
