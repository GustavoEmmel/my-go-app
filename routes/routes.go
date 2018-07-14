package routes

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"log"
	"my-go-app/pkg/controller/person"
)

/*
func Routes() {
	http.HandleFunc("/person", person.Person)
	http.ListenAndServe(":8080", nil)
}
*/

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to my golang app!\n" +
		"Version: 1.0")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Routes() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/person/", person.GetAllPerson)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8282", router))
}
