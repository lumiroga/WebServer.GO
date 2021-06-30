package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	init := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://reforma.com",
	}

	for _, servidor := range servidores {
		checkServer(servidor)
	}

	timeLapse := time.Since(init)

	fmt.Printf("Execution time %s", timeLapse)
}

func checkServer(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(servidor, " esta funcionando normalmente")
	}
}
