/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Set package program
 * Untuk memberitaukan bahwa ini program ini dari package main. */
package main

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Import Package / Library
 * Untuk memberikan fitur / fungsi tambahan yang lebih fleksible untuk program ini. */
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"Rest-API/Controllers"
	"Rest-API/Utils"
	"syscall"
)

/* Main program
 * disinilah tempat pemanggilan beberapa function untuk dijalankan pada poin awal / entry poin. */
func main() {
	// Panggil function ClearScreen() dari folder Utils/
	Utils.ClearScreen()
	
	// Set port untuk menjalankan program berdasarkan port yang telah diset!
	const port = ":5050"

	// Menampilkan output variasi, supaya terlihat lebih bagus (output header)
	fmt.Printf("\n > Server running on: %s", Utils.Serv.Sprintf("http://localhost%s", port))
	fmt.Println("\n\n -------------------- [ LOG History ] --------------------\n")

	// Set variable router untuk menampung object dari NewRouter
	router := mux.NewRouter()

	/* Set route program saya sesuai dengan SOAL
	 * pada /nama menggunakan method POST 		=> berfungsi untuk POST data (Nama NIM Alamat)
	 * untuk /semuadata menggunakan method GET	=> berfungsi untuk GET data (tampilkan semua data dari array of struct)
	 *
	 * /nama akan meneruskan request ke function CreateData() yang berada di folder Controllers
	 * hasilnya berupa response yang akan di teruskan ke user
	 *
	 * /semuadata sama seperti /nama, hanya beda fungsi yaitu menggunakan fungsi ReadData()
	 */
	router.HandleFunc("/nama", Controllers.CreateData).Methods("POST")
	router.HandleFunc("/semuadata", Controllers.ReadData).Methods("GET")

	/* Untuk handle jika route tidak ditemukan, misal :
	 * ketika user request root domain / selain '/nama' & '/semuadata'
	 * maka akan dihandle disini, dengan cara : memanggil function yang ada di folder Controllers/
	 */
	router.NotFoundHandler = http.HandlerFunc(Controllers.HandleNotFound)

	
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
