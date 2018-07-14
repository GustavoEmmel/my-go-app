package file_writer

import (
	"encoding/json"
	"io/ioutil"
)

func WriteJsonFile(list interface{}, file_name string) {

	listJson, _ := json.Marshal(list)
	_ = ioutil.WriteFile(file_name, listJson, 0644)

}
