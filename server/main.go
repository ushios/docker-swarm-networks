package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	name = flag.String("name", "", "server name")
	port = flag.Int("port", 8080, "server port")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, *name)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil))
}
