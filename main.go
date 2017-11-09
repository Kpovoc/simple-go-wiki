package main

import (
	"net/http"
	"github.com/Kpovoc/simple-go-wiki/src/handlers"
)

func main() {
	handlers.Init()
	http.ListenAndServe(":8080", nil)
}