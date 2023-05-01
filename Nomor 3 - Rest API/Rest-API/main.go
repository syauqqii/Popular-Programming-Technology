/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Set package program
 * Untuk memberitaukan bahwa ini program ini dari package main. */
package main

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Import Package / Library
 * Untuk memberikan fitur / fungsi tambahan yang lebih fleksible untuk program ini. */
import (
	"fmt"
	"github.com/gorilla/mux"
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

	/* variable server saya set untuk membuat sebuat object / instance dari http.Server
	 * yang bertujuan untuk membuat server HTTP */
	server := &http.Server{Addr: port, Handler: router}

	/* Logic Handle CTRL + C, References : https://pkg.go.dev/os/signal
	 * disini saya membuat sebuah program untuk handle jika developer menekan CTRL + C
	 * supaya tidak ada output yang menggangu ketika kita menekan CTRL + C saat program dijalankan.
	 *
	 * SIGINT (interrupt signal) akan dihasilkan ketika user menekan CTRL + C
	 *
	 * pada variable isPressedCTRLC akan menunggu adanya sinyal CTRL+C (SIGINT)
	 * pada function signal.Notify() akan memberitahu bahwa ada sinyal yang ditangkap ketika CTRL + C ditekan
	 */
	isPressedCTRLC := make(chan os.Signal, 1)
	signal.Notify(isPressedCTRLC, syscall.SIGINT)

	/* Pada fungsi tanpa nama ini (goroutine) / anonymous function
	 * berfungsi untuk menunggu SIGINT di tekan
	 * jika key pressed, maka akan menampilakn log pada console,
	 * dilanjut untuk cek, saat server di tutup apakah terjadi error, jika tidak maka exit
	 * jika terjadi error maak akan menampilkan error di log console.
	 */
	go func() {
		<-isPressedCTRLC
		Utils.Logger(3, "Berhasil mematikan server.")

		if err := server.Close(); err != nil {
			Utils.Logger(2, err)
		}
		os.Exit(0)
	}()

	/* pada code dibawah akan menjalankan server dengan port yang telah di tentukan diatas (line code 58)
	 * serta handle jika terdapat error pada server maka akan exit (function utils.logger params 4 [FATAL])
	 */
	if err := server.ListenAndServe(); err != nil {
		Utils.Logger(4, fmt.Sprintf("%s", err))
	}
}
