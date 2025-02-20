package main

import (
	"fmt"
	"net/http"
)

const port_num string = ":8080"

func welcome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go web server!")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func start_server() {

	err := http.ListenAndServe(port_num, nil)
	if err != nil {
		fmt.Println("Error starting the server - ", err)
	}

}

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/headers", headers)

	fmt.Println("Server is running on http://localhost" + port_num)
	fmt.Println("Press CTRL+C to stop the server")
	start_server()
}
