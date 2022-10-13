package controller

import (
	"fmt"
	"github.com/bdxygy/exception"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path"
)

func UploadController(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := fmt.Fprintf(w, "Hello Upload!")
	exception.Throw(err)
	// 1. Parse input, multipart/form-data
	err = r.ParseMultipartForm(20 << 10)
	exception.Throw(err)

	// 2. Retrieve file posted on form-data
	file, handler, err := r.FormFile("media")
	exception.Throw(err)
	defer file.Close()

	fmt.Println("File Upload Name: %v\n", handler.Filename)
	fmt.Println("File Upload Size: %v\n", handler.Size)
	fmt.Println("File Upload Mime Header: %v\n", handler.Header)

	// 3. Write temporary file on server
	base := path.Base(handler.Filename)
	created, err := os.Create(base)
	exception.Throw(err)
	defer created.Close()
	_, err = io.Copy(created, file)
	exception.Throw(err)

	//	4. return success or not for file uploaded
	fmt.Fprintf(w, "Upload succes")
}
