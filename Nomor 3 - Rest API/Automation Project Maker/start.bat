@ECHO OFF
cls

:: -------------------------- [ file processing ] --------------------------
chdir /d D:\
mkdir Rest-API
pushd Rest-API
type nul > main.go

:: ------------------------------ [ main.go ] ------------------------------
:: KOSONG - NEXT UPDATE SAJA

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
:: KOSONG - NEXT UPDATE SAJA

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

:: -------------------------- [ init project ] --------------------------
go work init
echo.>> go.work
echo use (>> go.work
echo  ./Controllers>> go.work
echo  ./Utils>> go.work
echo )>> go.work

pause
