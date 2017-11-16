package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	c := http.Client{}
	t := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-t.C:
			_, err := c.Get("http://server:8080")
			if err != nil {
				log.Println(err)
			}
		}
	}
}
