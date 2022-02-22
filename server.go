package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
