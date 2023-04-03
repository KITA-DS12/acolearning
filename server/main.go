package main

import (
	"log"
	"net/http"

	"github.com/KITA-DS12/acolearning.git/controllers"
)

func main() {
	router := controllers.Router()
	log.Fatal(http.ListenAndServe(":8888", router))
}
