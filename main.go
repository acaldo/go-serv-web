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

	i := 0

	for i < 2 {

		for _, servidor := range servs {
			go checkServer(servidor, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		i++
	}
	finish := time.Since(start)
	fmt.Println("Tiempo de ejecucion: ", finish)
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
