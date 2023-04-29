package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"Rest-API/Utils"
)

type Mahasiswa struct {
	Nama   string `json:"Nama"`
	NIM    string `json:"Nim"`
	Alamat string `json:"Alamat"`
}

var data []Mahasiswa

func ReadData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	if len(data) == 0 {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprint(response, `{"status":"Error", "message":"Data Not Found"}`)
		Utils.Logger(2, "func 'ReadData()' -> Data not found")
		return
	}

	extractData, err := json.Marshal(data)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(response, `{"status":"Error", "message":"Internal Server Error"}`)
		Utils.Logger(2, "func 'ReadData()' -> Internal Server Error")
		return
	}

	fmt.Fprint(response, string(extractData))
	Utils.Logger(1, "func 'ReadData()' -> Success Show All Data")
}

func CreateData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var person Mahasiswa

	person.Nama   = request.FormValue("Nama")
	person.NIM    = request.FormValue("NIM")
	person.Alamat = request.FormValue("Alamat")

	if person.Nama == "" || person.NIM == "" || person.Alamat == "" {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(response, `{"status":"Error", "message":"Bad Request"}`)
		Utils.Logger(2, "func 'CreateData()' -> Bad Request")
		return
	}

	data = append(data, person)
	extractData, err := json.Marshal(person)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(response, `{"status":"Error", "message":"Internal Server Error"}`)
		Utils.Logger(2, "func 'CreateData()' -> Internal Server Error")
		return
	}

	fmt.Fprint(response, string(extractData))
	Utils.Logger(1, "func 'CreateData()' -> Success Add Data")
}