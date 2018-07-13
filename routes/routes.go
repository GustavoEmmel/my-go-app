package routes

import (
	"net/http"
	"my-go-app/pkg/controller/person"
)

func Routes() {
	http.HandleFunc("/person", person.Person)
	http.ListenAndServe(":8080", nil)

}
