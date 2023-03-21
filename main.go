package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s! From Develop! Favourite sauce: %s", r.URL.Path[1:], os.Getenv("sauce"))
}
