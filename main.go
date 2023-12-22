package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)
	fmt.Fprintf(w, "Your IP is %s\n", ip)
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func getIP(r *http.Request) string {
	ip := r.RemoteAddr
	if index := strings.LastIndex(ip, ":"); index != -1 {
		ip = ip[:index]
	}
	return ip
}
