package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":4000"

type application struct {}

func main()  {
	// app := application {}
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Dinh!!")	
	})

	fmt.Println("Starting web application on port", port)
	
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic(err)
	}
}