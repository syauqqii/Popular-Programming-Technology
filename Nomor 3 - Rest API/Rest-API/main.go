package main

import (
	"fmt"
	"encoding/json"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"Rest-API/Controllers"
	"Rest-API/Utils"
	"os"
	"os/signal"
	"syscall"
	"io/ioutil"
)

func main() {
	Utils.ClearScreen()

	serv := color.New(color.FgCyan, color.Underline)
	const port = ":5050"

	fmt.Print("\n > Server running on: ")
	serv.Printf("http://localhost%s", port)
	fmt.Println("\n\n -------------------- [ LOG History ] --------------------\n")

	router := mux.NewRouter()

	router.HandleFunc("/nama", Controllers.CreateData).Methods("POST")
	router.HandleFunc("/semuadata", Controllers.ReadData).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request){
		response.Header().Set("Content-Type", "application/json")
        response.WriteHeader(http.StatusNotFound)

	    dataL := map[string]string{
	    	"get": "http://localhost:5050/semuadata",
	        "post": "http://localhost:5050/nama",
	    }

	    jsonData, _ := json.Marshal(map[string]interface{}{
	        "status":  "Error",
	        "message": dataL,
	    })

	    fmt.Fprint(response, string(jsonData))
	    Utils.Logger(2, "Bad Request 'Not Found Router' -> Status Not Found")
	})

	server := &http.Server{Addr: port, Handler: router}

	sigintChan := make(chan os.Signal, 1)
	signal.Notify(sigintChan, syscall.SIGINT)

	go func() {
		<-sigintChan
		Utils.Logger(3, "Shutting Down Server Success")

		if err := server.Close(); err != nil {
			Utils.Logger(2, "err")
		}
		os.Exit(0)
	}()

	log.Fatal(server.ListenAndServe())
	log.SetOutput(ioutil.Discard)
}