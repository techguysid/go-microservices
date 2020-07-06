package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hello Handler
type Hello struct {
	l *log.Logger
}

func newHello(l *log.Logger) *Hello {
	return &Hello{l}
}

//ServeHTTP Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
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
	h.l.Printf("%s\n", d)
	h.l.Printf("%T", d)
	fmt.Fprintf(rw, "Hello %s", d)
}
