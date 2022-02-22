package main

import (
<<<<<<< HEAD
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
=======
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

  fmt.Println("Server is running on port 3000")
  http.ListenAndServe(":3000", nil)
}
>>>>>>> 77911b9 (implemented handler for /time)
