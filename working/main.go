package main

import (
	"log"
	"net/http"
	"os",
	"github.com/techguysid/working/handlers"
)

func main() {
	//http Handle Func converts the function you are passing in parameter into an http handler and registers it under Default ServeMux
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.newHello()

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", nil) //DefaultServeMux is the default handler for ListenAndServe, ServeMux is a HTTP request handler(multiplexer)
}
