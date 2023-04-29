@ECHO OFF
cls

:: -------------------------- [ file processing ] --------------------------
chdir /d D:\
mkdir Rest-API
pushd Rest-API
type nul > main.go

:: ------------------------------ [ main.go ] ------------------------------
echo package main> main.go
echo.>> main.go
echo import (>> main.go
echo 	"fmt">> main.go
echo 	"log">> main.go
echo 	"net/http">> main.go
echo 	"github.com/fatih/color">> main.go
echo 	"github.com/gorilla/mux">> main.go
echo 	"Rest-API/Controllers">> main.go
echo )>> main.go
echo.>> main.go
echo func main() {>> main.go
echo 	serv := color.New(color.FgCyan, color.Underline)>> main.go
echo 	port := ":5050">> main.go
echo.>> main.go
echo 	fmt.Print("\n > Server running on: ")>> main.go
echo 	serv.Printf("http://localhost%%s\n", port)>> main.go
echo.>> main.go
echo 	router := mux.NewRouter()>> main.go
echo.>> main.go
echo 	router.HandleFunc("/nama", Controllers.CreateData).Methods("POST")>> main.go
echo 	router.HandleFunc("/semuadata", Controllers.ReadData).Methods("GET")>> main.go
echo.>> main.go
echo 	log.Fatal(http.ListenAndServe(port, router))>> main.go
echo }>> main.go

:: -------------------------- [ file processing ] --------------------------
mkdir Controllers
pushd Controllers
type nul > Logic-CRUD.go

:: ---------------------------- [ Logic-CRUD.go ] ----------------------------
:: KOSONG - NEXT UPDATE SAJA

:: -------------------------- [ file processing ] --------------------------
pushd ..
mkdir Utils
pushd Utils
type nul > CLS.go
type nul > Log-History.go

:: -------------------------- [ init project ./Controllers ] --------------------------
pushd ..
go mod init Rest-API
go get github.com/fatih/color
go get github.com/gorilla/mux
go mod tidy
move go.mod Controllers/

:: -------------------------- [ init project ./Utils ] --------------------------
go mod init Rest-API
move go.mod Utils/

pause
