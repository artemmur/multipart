package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getFiles(w http.ResponseWriter, r *http.Request) {
	log.Println("HERE")
	reader, err := r.MultipartReader()
	if err != nil {
		log.Println(err)
		return
	}
LOOP:
	for {
		part, err := reader.NextPart()
		if err != nil {
			switch {
			case errors.Is(err, io.EOF):
				break LOOP
			default:
				log.Println(err)
			}
		}
		defer part.Close()

		fmt.Println(part)
	}
}

func main() {
	http.HandleFunc("/", getFiles)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalln("cannot serve", err)
	}
}
