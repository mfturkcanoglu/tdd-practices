package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func myGreetingHandler(rw http.ResponseWriter, r *http.Request) {
	Greet(rw, "world!")
}

func main() {
	Greet(os.Stdout, "Mf")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(myGreetingHandler)))
}
