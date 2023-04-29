:: Creatornya siafh lgi klo bukan syh :)
@ECHO OFF
cls

:: -------------------------- [ file processing ] --------------------------
:: chdir /d D:\
mkdir Rest-API
pushd Rest-API
type nul > main.go

:: ------------------------------ [ main.go ] ------------------------------
echo package main>> main.go
echo.>> main.go
echo import (>> main.go
echo 	"fmt">> main.go
echo 	"encoding/json">> main.go
echo 	"github.com/fatih/color">> main.go
echo 	"github.com/gorilla/mux">> main.go
echo 	"log">> main.go
echo 	"net/http">> main.go
echo 	"Rest-API/Controllers">> main.go
echo 	"Rest-API/Utils">> main.go
echo 	"os">> main.go
echo 	"os/signal">> main.go
echo 	"syscall">> main.go
echo 	"io/ioutil">> main.go
echo )>> main.go
echo.>> main.go
echo func main() {>> main.go
echo 	Utils.ClearScreen()>> main.go
echo.>> main.go
echo 	serv := color.New(color.FgCyan, color.Underline)>> main.go
echo 	const port = ":5050">> main.go
echo.>> main.go
echo 	fmt.Print("\n > Server running on: ")>> main.go
echo 	serv.Printf("http://localhost%%s", port)>> main.go
echo 	fmt.Println("\n\n -------------------- [ LOG History ] --------------------\n")>> main.go
echo.>> main.go
echo 	router := mux.NewRouter()>> main.go
echo.>> main.go
echo 	router.HandleFunc("/nama", Controllers.CreateData).Methods("POST")>> main.go
echo 	router.HandleFunc("/semuadata", Controllers.ReadData).Methods("GET")>> main.go
echo.>> main.go
echo 	router.NotFoundHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request){>> main.go
echo 		response.Header().Set("Content-Type", "application/json")>> main.go
echo         response.WriteHeader(http.StatusNotFound)>> main.go
echo.>> main.go
echo 	    dataL := map[string]string{>> main.go
echo 	    	"get": "http://localhost:5050/semuadata",>> main.go
echo 	        "post": "http://localhost:5050/nama",>> main.go
echo 	    }>> main.go
echo.>> main.go
echo 	    jsonData, _ := json.Marshal(map[string]interface{}{>> main.go
echo 	        "status":  "Error",>> main.go
echo 	        "message": dataL,>> main.go
echo 	    })>> main.go
echo.>> main.go
echo 	    fmt.Fprint(response, string(jsonData))>> main.go
echo 	    Utils.Logger(2, "Bad Request 'Not Found Router' -> Status Not Found")>> main.go
echo 	})>> main.go
echo.>> main.go
echo 	server := ^&http.Server{Addr: port, Handler: router}>> main.go
echo.>> main.go
echo 	sigintChan := make(chan os.Signal, 1)>> main.go
echo 	signal.Notify(sigintChan, syscall.SIGINT)>> main.go
echo.>> main.go
echo 	go func() {>> main.go
echo 		^<-sigintChan>> main.go
echo 		Utils.Logger(3, "Shutting Down Server Success")>> main.go
echo.>> main.go
echo 		if err := server.Close(); err != nil {>> main.go
echo 			Utils.Logger(2, "err")>> main.go
echo 		}>> main.go
echo 		os.Exit(0)>> main.go
echo 	}()>> main.go
echo.>> main.go
echo 	log.Fatal(server.ListenAndServe())>> main.go
echo 	log.SetOutput(ioutil.Discard)>> main.go
echo }>> main.go

:: -------------------------- [ file processing ] --------------------------
mkdir Controllers
pushd Controllers
type nul > Logic-CRUD.go

:: ---------------------------- [ Logic-CRUD.go ] ----------------------------
echo package Controllers>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo import (>> Logic-CRUD.go
echo 	"encoding/json">> Logic-CRUD.go
echo 	"fmt">> Logic-CRUD.go
echo 	"net/http">> Logic-CRUD.go
echo 	"Rest-API/Utils">> Logic-CRUD.go
echo )>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo type Mahasiswa struct {>> Logic-CRUD.go
echo 	Nama   string `json:"Nama"`>> Logic-CRUD.go
echo 	NIM    string `json:"Nim"`>> Logic-CRUD.go
echo 	Alamat string `json:"Alamat"`>> Logic-CRUD.go
echo }>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo var data []Mahasiswa>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo func ReadData(response http.ResponseWriter, request *http.Request) {>> Logic-CRUD.go
echo 	response.Header().Set("Content-Type", "application/json")>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	if len(data) == 0 {>> Logic-CRUD.go
echo 		response.WriteHeader(http.StatusNotFound)>> Logic-CRUD.go
echo 		fmt.Fprint(response, `{"status":"Error", "message":"Data Not Found"}`)>> Logic-CRUD.go
echo 		Utils.Logger(2, "func 'ReadData()' -> Data not found")>> Logic-CRUD.go
echo 		return>> Logic-CRUD.go
echo 	}>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	extractData, err := json.Marshal(data)>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	if err != nil {>> Logic-CRUD.go
echo 		response.WriteHeader(http.StatusInternalServerError)>> Logic-CRUD.go
echo 		fmt.Fprint(response, `{"status":"Error", "message":"Internal Server Error"}`)>> Logic-CRUD.go
echo 		Utils.Logger(2, "func 'ReadData()' -> Internal Server Error")>> Logic-CRUD.go
echo 		return>> Logic-CRUD.go
echo 	}>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	fmt.Fprint(response, string(extractData))>> Logic-CRUD.go
echo 	Utils.Logger(1, "func 'ReadData()' -> Success Show All Data")>> Logic-CRUD.go
echo }>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo func CreateData(response http.ResponseWriter, request *http.Request) {>> Logic-CRUD.go
echo 	response.Header().Set("Content-Type", "application/json")>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	var person Mahasiswa>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	person.Nama   = request.FormValue("Nama")>> Logic-CRUD.go
echo 	person.NIM    = request.FormValue("NIM")>> Logic-CRUD.go
echo 	person.Alamat = request.FormValue("Alamat")>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	if person.Nama == "" ^|^| person.NIM == "" ^|^| person.Alamat == "" {>> Logic-CRUD.go
echo 		response.WriteHeader(http.StatusBadRequest)>> Logic-CRUD.go
echo 		fmt.Fprint(response, `{"status":"Error", "message":"Bad Request"}`)>> Logic-CRUD.go
echo 		Utils.Logger(2, "func 'CreateData()' -> Bad Request")>> Logic-CRUD.go
echo 		return>> Logic-CRUD.go
echo 	}>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	data = append(data, person)>> Logic-CRUD.go
echo 	extractData, err := json.Marshal(person)>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	if err != nil {>> Logic-CRUD.go
echo 		response.WriteHeader(http.StatusInternalServerError)>> Logic-CRUD.go
echo 		fmt.Fprint(response, `{"status":"Error", "message":"Internal Server Error"}`)>> Logic-CRUD.go
echo 		Utils.Logger(2, "func 'CreateData()' -> Internal Server Error")>> Logic-CRUD.go
echo 		return>> Logic-CRUD.go
echo 	}>> Logic-CRUD.go
echo.>> Logic-CRUD.go
echo 	fmt.Fprint(response, string(extractData))>> Logic-CRUD.go
echo 	Utils.Logger(1, "func 'CreateData()' -> Success Add Data")>> Logic-CRUD.go
echo }>> Logic-CRUD.go

:: -------------------------- [ file processing ] --------------------------
pushd ..
mkdir Utils
pushd Utils
type nul > CLS.go
type nul > Log-History.go

:: ---------------------------- [ CLS.go ] ----------------------------
echo package Utils>> CLS.go
echo.>> CLS.go
echo import (>> CLS.go
echo 	"fmt">> CLS.go
echo 	"os">> CLS.go
echo 	"os/exec">> CLS.go
echo 	"runtime">> CLS.go
echo )>> CLS.go
echo.>> CLS.go
echo func ClearScreen(){>> CLS.go
echo 	_os := runtime.GOOS>> CLS.go
echo.>> CLS.go
echo 	if _os == "windows" {>> CLS.go
echo 		cmd := exec.Command("cmd", "/c", "cls")>> CLS.go
echo 		cmd.Stdout = os.Stdout>> CLS.go
echo 		cmd.Run()>> CLS.go
echo 	} else if _os == "linux" {>> CLS.go
echo 		cmd := exec.Command("clear")>> CLS.go
echo 		cmd.Stdout = os.Stdout>> CLS.go
echo 		cmd.Run()>> CLS.go
echo   	} else {>> CLS.go
echo 		fmt.Print()>> CLS.go
echo 	}>> CLS.go
echo }>> CLS.go

:: ---------------------------- [ Log-History.go ] ----------------------------
echo package Utils>> Log-History.go
echo.>> Log-History.go
echo import (>> Log-History.go
echo 	"fmt">> Log-History.go
echo 	"github.com/fatih/color">> Log-History.go
echo 	"time">> Log-History.go
echo )>> Log-History.go
echo.>> Log-History.go
echo func Logger(options int, msg string){>> Log-History.go
echo 	now := time.Now()>> Log-History.go
echo.>> Log-History.go
echo 	grn := color.New(color.FgHiGreen)>> Log-History.go
echo 	cyn := color.New(color.FgHiCyan)>> Log-History.go
echo 	red := color.New(color.FgHiRed)>> Log-History.go
echo 	yel := color.New(color.FgHiYellow)>> Log-History.go
echo.>> Log-History.go
echo 	fmt.Print(" [")>> Log-History.go
echo 	yel.Printf("%%s", now.Format("15:04:05 2006/01/02"))>> Log-History.go
echo 	fmt.Print("] ")>> Log-History.go
echo 	if options == 1 {>> Log-History.go
echo 		cyn.Print("INFO")>> Log-History.go
echo 	} else if options == 2 {>> Log-History.go
echo 		red.Print("ERROR")>> Log-History.go
echo 	} else if options == 3 {>> Log-History.go
echo 		grn.Print("EXIT")>> Log-History.go
echo 	}>> Log-History.go
echo.>> Log-History.go
echo 	fmt.Printf(": %%s\n", msg)>> Log-History.go
echo }>> Log-History.go

:: -------------------------- [ init project ] --------------------------
pushd ..
go work init
set "ver="
for /f "delims=" %%i in ('type go.work ^| findstr /r /c:".*"') do (
    set "ver=%%i"
    goto :done
)
:done
echo.>> go.work
echo use (>> go.work
echo  ./Controllers>> go.work
echo  ./Utils>> go.work
echo )>> go.work

:: -------------------------- [ init project ./Controllers ] --------------------------
go mod init Rest-API
echo module Rest-API/Controllers> go.mod
echo.>> go.mod
echo %ver%>> go.mod
go get github.com/fatih/color
go get github.com/gorilla/mux
go mod tidy
move go.mod Controllers/

:: -------------------------- [ init project ./Utils ] --------------------------
go mod init Rest-API
echo module Rest-API/Utils> go.mod
echo.>> go.mod
echo %ver%>> go.mod
move go.mod Utils/

pause
