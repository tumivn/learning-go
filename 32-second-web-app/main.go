package main

import (
	"fmt"
	"github.com/tumivn/go-booking-app/pkg/handlers"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
