package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("Error getting a file for the provided form key : ", err)
		return
	}

	defer file.Close()
	out, pathError := os.Create("tmp/uploadedFile")
	if pathError != nil {
		log.Println("error creating a file for writing: ", pathError)
	}
	defer out.Close()
	_, copyFileError := io.Copy(out, file)
	if copyFileError != nil {
		log.Println("error occured while file copy: ", copyFileError)
	}
	fmt.Fprint(w, "File uploadedFile successfully: "+header.Filename)

}
