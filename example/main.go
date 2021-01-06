package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeeEirc/elfinder"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./elf/"))))
	mux.Handle("/connector", http.HandlerFunc(elfHandler))
	fmt.Println("Listen on :80")
	log.Fatal(http.ListenAndServe(":00", mux))
}

func elfHandler(w http.ResponseWriter, r *http.Request) {

	con := elfinder.NewElFinderConnector([]elfinder.Volume{&elfinder.DefaultVolume})
	con.ServeHTTP(w, r)
}
