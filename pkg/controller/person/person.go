package person

import (
	"net/http"
	"fmt"
	"my-go-app/pkg/repository/person"
	"my-go-app/pkg/service/file_writer"
)

func Person(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		listAllPerson(w, r)
	case "PUT":
		fmt.Fprintf(w, "PUT PERSON")
	default:
		http.Error(w, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
	}
}

func listAllPerson(w http.ResponseWriter, r *http.Request) {
	res := person.FindAllPerson()
	file_writer.WriteJsonFile(res, "teste2.json")
	fmt.Println(res)
}
