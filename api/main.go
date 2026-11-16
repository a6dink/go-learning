package main
import (
	"fmt"
	"net/http"
	"log"
	"html"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, bar is me: %s", html.EscapeString(r.URL.Path))
		fmt.Println("/bar endpoint was called")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}