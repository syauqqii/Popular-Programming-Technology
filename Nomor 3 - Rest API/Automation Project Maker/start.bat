@ECHO OFF
cls
chdir /d D:\
mkdir Rest-API
pushd Rest-API
type nul > main.go
mkdir Controllers
pushd Controllers
type nul > Logic-CRUD.go
pushd ..
go mod init Rest-API
go get github.com/fatih/color
go get github.com/gorilla/mux
go mod tidy
