package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//CheckAuth simulates an authentication
func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := false
			fmt.Println("Checking Authentication")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

//Login authenticates an user
func Login() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			f(w, r)
		}
	}
}
