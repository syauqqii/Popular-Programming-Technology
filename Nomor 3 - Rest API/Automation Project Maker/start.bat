@ECHO OFF
cls
chdir /d D:\
mkdir Rest-API
pushd Rest-API
type nul > main.go

:: ------------------------------ [ main.go ] ------------------------------
echo package main> main.go
echo.>> main.go
echo import (>> main.go
echo 	"log">> main.go
echo 	"net/http">> main.go
echo 	"github.com/fatih/color">> main.go
echo 	"github.com/gorilla/mux">> main.go
echo 	"Rest-API/Controllers">> main.go
echo )>> main.go
echo.>> main.go
echo func main() {>> main.go
echo 	router := mux.NewRouter()>> main.go
echo.>> main.go
echo 	serv := color.New(color.FgCyan, color.Underline)>> main.go
echo 	port := ":5050">> main.go
echo.>> main.go
echo 	log.Print("\n > Server running on: ")>> main.go
echo 	serv.Printf("http://localhost%%s\n", port)>> main.go
echo.>> main.go
echo 	router.HandleFunc("/nama", Controllers.CreateData).Methods("POST")>> main.go
echo 	router.HandleFunc("/semuadata", Controllers.ReadData).Methods("GET")>> main.go
echo.>> main.go
echo 	log.Fatal(http.ListenAndServe(port, router)>> main.go
echo }>> main.go

mkdir Controllers
pushd Controllers
type nul > Logic-CRUD.go

:: ---------------------------- [ Logic-CRUD.go ] ----------------------------
echo package Controllers>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo import (>> Logic-CRUD.go
echo 	"encoding/json">> Logic-CRUD.go
echo 	"net/http">> Logic-CRUD.go
echo )>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo type Data struct {>> Logic-CRUD.go
echo 	Name	string `json:"name"`>> Logic-CRUD.go
echo 	NIM	string `json:"nim"`>> Logic-CRUD.go
echo 	Alamat	string `json:"alamat"`>> Logic-CRUD.go
echo }>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo var datas []data>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo func CreateData(response http.ResponseWriter, request *http.Request){>> Logic-CRUD.go
echo 	json.NewEncoder(response).Encode(datas)>> Logic-CRUD.go
echo }>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo func ReadData(response http.ResponseWriter, request *http.Request){>> Logic-CRUD.go
echo 	var data Data>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	_ = json.NewDecoder(request.Body).Decode(&data)>> Logic-CRUD.go
echo 	datas = append(datas, data)>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	json.NewEncoder(response).Encode(data)>> Logic-CRUD.go
echo }>> Logic-CRUD.go

pushd ..
go mod init Rest-API
go get github.com/fatih/color
go get github.com/gorilla/mux
go mod tidy
pause
