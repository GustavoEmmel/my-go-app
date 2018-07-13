package routes

import (
	"my-go-app/pkg/controller/person"
	"net/http"
)

func Routes() {
	http.HandleFunc("/person", person.Person)
	http.ListenAndServe(":8090", nil)

}
