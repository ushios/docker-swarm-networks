package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{}
	t := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-t.C:
			_, err := c.Get(fmt.Sprintf("http://server:8080?c=%s", os.Getenv("HOSTNAME")))
			if err != nil {
				log.Println(err)
			}
		}
	}
}
