package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		fmt.Println("for loop is running")
		go func(l string) {
			time.Sleep(2 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "- error")
		c <- link
		return
	}
	fmt.Println(link, "- OK")
	c <- link
}
