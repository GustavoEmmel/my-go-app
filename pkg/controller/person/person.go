package person

import (
	"net/http"
	"fmt"
	"my-go-app/pkg/service/db-connector"
	"my-go-app/pkg/entity/person"
	"encoding/json"
	"io/ioutil"
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
	db := db_connector.Conn()

	results, err := db.Query("SELECT id, firstname FROM person")
	if err != nil {
		panic(err.Error())
	}

	per := person.Person{}
	res := []person.Person{}

	for results.Next() {

		var id string
		var firstname string

		err = results.Scan(&id, &firstname)
		if err != nil {
			panic(err.Error())
		}

		per.ID = id
		per.Name = firstname

		res = append(res, per)

	}

	personJson, _ := json.Marshal(res)
	err = ioutil.WriteFile("output.json", personJson, 0644)

	fmt.Println(res)
	defer db.Close()
}
