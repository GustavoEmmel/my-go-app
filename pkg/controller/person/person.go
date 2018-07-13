package person

import (
	"net/http"
	"fmt"
)

func Person(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET PERSON")
	case "PUT":
		fmt.Fprintf(w, "PUT PERSON")
	default:
		http.Error(w, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
	}
}
