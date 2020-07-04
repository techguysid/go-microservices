package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			/* This could be done but HTTP has got us covered with Error method
			rw.WriteHeader(http.StatusBadRequest) //WriteHeader lets us specify the status code
			rw.Write([]byte("Oops"))
			return
			*/
			http.Error(rw, "Oops Something bad occurred.", http.StatusBadRequest)
			return //We need this return since http Error doesn't terminate the app
		}
		log.Printf("%s\n", d)
		log.Printf("%T", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})
	http.ListenAndServe(":9090", nil) //DefaultServeMux is the default handler for ListenAndServe, ServeMux is a HTTP request handler(multiplexer)
}
