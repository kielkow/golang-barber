package main

import (
	"log"
	"net/http"
)

const apiBasePath = "/api"

func main() {
	log.Fatal(http.ListenAndServe(":3333", nil))
}
