go work init
set "ver="
for /f "delims=" %%i in ('type go.work ^| findstr /r /c:".*"') do (
    set "ver=%%i"
    goto :done
)
:done
echo.>> go.work
echo use ./mahasiswa>> go.work
go mod init json-response
echo module json-response/mahasiswa> go.mod
echo.>> go.mod
echo %ver%>> go.mod
go get github.com/gorilla/mux
go mod tidy
move go.mod mahasiswa/
