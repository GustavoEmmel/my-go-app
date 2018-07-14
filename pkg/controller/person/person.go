package person

import (
	"net/http"
	"fmt"
	"my-go-app/pkg/repository/person"
	"my-go-app/pkg/service/file_writer"
	"github.com/julienschmidt/httprouter"
)

func GetAllPerson(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := person.FindAllPerson()
	file_writer.WriteJsonFile(res, "teste2.json")
	fmt.Println(res)
}
