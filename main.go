package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)

	servs := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://youtube.com",
		"http://instagram.com",
	}
	for _, serv := range servs {
		go checkServer(serv, channel)
	}

	for i := 0; i < len(servs); i++ {
		fmt.Println(<-channel)
	}
	time := time.Since(start)
	fmt.Println("Elapsed time:", time)
}

func checkServer(serv string, channel chan string) {
	_, err := http.Get(serv)
	if err != nil {
		//fmt.Println(serv, "is down")
		channel <- serv + " is down"
	}
	//fmt.Println(serv, "is up")
	channel <- serv + " is up"
}
