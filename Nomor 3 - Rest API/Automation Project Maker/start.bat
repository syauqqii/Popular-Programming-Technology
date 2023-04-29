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
:: KOSONG - NEXT UPDATE SAJA

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
