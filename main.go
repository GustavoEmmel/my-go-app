package main

import (
	"my-go-app/pkg/entity/person"
	"my-go-app/pkg/repository/hello"
	"my-go-app/pkg/service/db-connector"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"my-go-app/routes"
)

func main() {
	t := hello.Hello()
	fmt.Println(t)

	listAllPerson()

	routes.Routes()
}

func listAllPerson() {
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
