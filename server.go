package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().UTC()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_ = json.NewEncoder(w).Encode(currentTime.Format("2006-01-02T15:04:05Z07:00"))
	})

	fmt.Println("Server is running on port 8795")
	http.ListenAndServe(":8795", nil)
}
//