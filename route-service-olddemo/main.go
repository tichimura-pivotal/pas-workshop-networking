package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is old path....")
}
