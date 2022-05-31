package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	servs := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://youtube.com",
		"http://instagram.com",
	}
	for _, serv := range servs {
		checkServer(serv)
	}
	time := time.Since(start)
	fmt.Println("Elapsed time:", time)
}

func checkServer(serv string) {
	_, err := http.Get(serv)
	if err != nil {
		fmt.Println(serv, "is down")
	}
	fmt.Println(serv, "is up")
}
