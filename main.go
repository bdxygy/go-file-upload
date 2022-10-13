package main

import (
	"github.com/bdxygy/go-file-upload/controller"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	router.POST("/upload", controller.UploadController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
