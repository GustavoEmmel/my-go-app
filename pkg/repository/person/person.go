package person

import (
	"my-go-app/pkg/entity/person"
	"my-go-app/pkg/service/db-connector"
)

func FindAllPerson() ([]person.Person) {

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
	defer db.Close()

	return res
}
