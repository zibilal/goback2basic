package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyHandler struct {}
func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My first handler")
}

type HelloHandler struct{}
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gengs")
}

func HelloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My first handler")
}

func logger(h http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		uri := request.RequestURI
		fmt.Println("Handler function called - " + uri)
		h(writer, request)
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/my", logger(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "My first handler")
	}) )
	http.HandleFunc("/hello", logger(HelloHandlerFunc) )

	log.Fatal(server.ListenAndServe())
}