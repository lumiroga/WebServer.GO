package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	init := time.Now()
	channel := make(chan string)
	servidores := []string{
		"http://yahoo.com",
		"http://google.com",
		"http://reforma.com",
	}

	i := 0

	for {
		if i > 5 {
			break
		}

		for _, servidor := range servidores {
			go checkServer(servidor, channel)

		}
		time.Sleep(time.Second)
		fmt.Println(<-channel)
		i++
	}

	timeLapse := time.Since(init)

	fmt.Printf("Execution time %s", timeLapse)
}

func checkServer(servidor string, channel chan string) {
	_, err := http.Get(servidor)
	if err != nil {

		channel <- servidor + "No se encuentra disponible"
	} else {

		channel <- servidor + "Functiona correctamente"
	}
}
