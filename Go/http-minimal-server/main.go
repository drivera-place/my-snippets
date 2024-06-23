package main

import (
	"log"
	"net/http"
)

const Port string = "8080"

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:"+Port, nil)
	log.Fatal(err)
}
