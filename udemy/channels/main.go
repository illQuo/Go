package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, w := range websites {
		go checkStatus(w, c)
	}

	for w := range c {
		go func(w string) {
			time.Sleep(time.Second * 5)
			checkStatus(w, c)
		}(w)
	}
}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%v is offline\n", link)
		c <- link
		return
	}

	fmt.Printf("%v is online\n", link)
	c <- link
}
