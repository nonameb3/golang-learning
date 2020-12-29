package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	proto := "http://"
	links := []string{
		proto + "google.com",
		proto + "facebook.com",
		proto + "twitter.com",
		proto + "linkdein.com",
	}

	channels := make(chan string)
	for _, link := range links {
		go checkLink(link, channels)
	}

	for link := range channels {
		go func(l string) {
			time.Sleep(2 * time.Second)
			checkLink(l, channels)
		}(link)
	}
	// fmt.Println(<-channels)
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, " is down!!")
		c <- l
		return
	}
	fmt.Println(l, "is up!!")
	c <- l
}
