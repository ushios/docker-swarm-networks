package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	name string
	port = flag.Int("port", 8080, "server port")
)

func main() {
	flag.Parse()

	name = os.Getenv("HOSTNAME")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		log.Printf("server(%s) client(%s)\n", name, v.Get("c"))
		fmt.Fprintf(w, name)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil))
}
