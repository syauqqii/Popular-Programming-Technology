/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Set package program
 * Untuk memberitaukan bahwa ini program ini dari package Controllers. */
package Controllers

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Import Package / Library
 * Untuk memberikan fitur / fungsi tambahan yang lebih fleksible untuk program ini. */
import (
	"encoding/json"
	"fmt"
	"net/http"
	"Rest-API/Utils"
)

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Struct dengan nama Mahasiswa serta field data:
 * 	1. Nama   : untuk menyimpan nilai Nama 
 * 	2. NIM    : untuk menyimpan nilai NIM
 * 	3. Alamat : untuk menyimpan nilai Alamat
 *
 * NOTE : 3 field data saya memiliki tipe data yang sama semua yaitu string. 
 *
 * setetlah tipe data string saya mendeklarasikan `json:*` dengan backtick
 * fungsinya untuk serialize supaya output nya nanti sama dengan yang saya deklarasikan. 
 * contoh :
 *   "Nama": "Dimas Syauqi Syafa"
 *   "Nim": "2502004405"
 *   "Alamat": "Jombang, Jawa Timur"
 *
 * Perlu diingat, bahwa itu case sensitiv. 
 * Saya menggunakan format sesuai di SOAL.
 */
type Mahasiswa struct {
	Nama   string `json:"Nama"`
	NIM    string `json:"Nim"`
	Alamat string `json:"Alamat"`
}

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Deklarasi array of struct data dengan mendapatkan format dari struct Mahasiswa
 * variable ini bertujuan menampung beberapa 'data' yang telah saya kelompokkan dari input.
 *
 * Ini menggunakan penyimpanan sederhana, tidak menggunakan database dinamis / statis
 * Jadi, ketika program di hentikan, maka data akan hilang. */
var data []Mahasiswa

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Function CreateData
 * Fungsi ini saya buat untuk menambahkan data ke array of struct (line code: 50). 
 *
 * @Params response http.ResponseWriter
 *	-> params ini akan mengembalikan proses dari handle request user
 *
 * @Params request *http.Request
 *	-> params ini akan menghandle request body user, (formvalue Nama, NIM, Alamat)
 */
func CreateData(response http.ResponseWriter, request *http.Request) {
	// Set respon MIME type json, bertujuan untuk memberi tau user bahwa output ini merupakan json.
	response.Header().Set("Content-Type", "application/json")

	/* Variable person dengan tipe data struct Mahasiswa ini memiliki fungsi
	 * untuk menampung data sementara, kemduain data akan di forward ke 
	 * var data []Mahasiswa (line code: 50) untuk ditampung dijadikan 1
	 */
	var person Mahasiswa

	/* Disini kita menentukan request bodynya
	 * contoh: 
	 *	(benar) input Nama menggunakan 'Nama=Dimas Syauqi Syafa'
	 *	(salah) input Nama menggunakan 'nama=Dimas Syauqi Syafa'
	 *
	 * bisa dilihat bahwa terdapat perbedaan request bodynya, yaitu case sensitive.
	 * Nama dan nama di sini berbeda, yang pertama menggunakan Kapital, yang kedua tidak menggunakan Kapital.
	 *
	 * Begitu juga seterusnya.. disini saya menggunakan format Nama, NIM, Alamat.
	 */
	person.Nama   = request.FormValue("Nama")
	person.NIM    = request.FormValue("NIM")
	person.Alamat = request.FormValue("Alamat")

	/* Apapun alasannyam jika ada salah satu dari (Nama, NIM, Alamt) dengan request body kosong,
	 * maka akan return bad request 400
	 * semua data harus diisi.
	 */
	if person.Nama == "" || person.NIM == "" || person.Alamat == "" {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(response, `{"status":"error", "message":"request body (Nama, NIM, Alamat) tidak boleh kosong!"}`)
		Utils.Logger(2, "func 'CreateData()' -> Pengisian data masih ada yang kosong.")
		return
	} else {
		/* Jika semua data terisi (Nama, NIM, Alamat) selanjutnya kita akan melakukan penambahan data
		 * kita akan menambahkan data person (line code 71) sebagai penyimpanan sementara untuk struct Mahasiwa 
		 * ke dalam array of struct data (line code 50).
		 */
		data = append(data, person)
		
		/* Selanjutnya, saya membuat code untuk mengekstrak data ke dalam array (json)
		 * dengan menggunakan function marshal yang memiliki 2 return variable
		 * 1. return untuk datanya, saya menggunakan extractData
		 * 2. return untuk error, saya menggunakan err
		 *
		 * pada logicnya, ketika err ! nil (kosong) / maksudnya variable error terisi,
		 * maka program akan lanjut ke section IF yaitu handle error
		 * dengan logic : return code 500 - Internal Server Error
		 *		  mengembalikan respon dalam bentuk JSON ke user / client
		 *		  mengeluarkan output error dalam bentuk log di console
		 *
		 * jika proses marshal aman, lanjut masuk ke section ELSE
		 * dengan logic : saya tidak set return code, karena secara default sudah 200 OK
		 * 		  saya tampilkan response (data yang telah di request) dalam bentuk json,
		 * 		  saya tampilkan output di console untuk log history
		 */
		if extractData, err := json.Marshal(person); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(response, `{"status":"error", "message":"terjadi kesalahan pada server."}`)
			Utils.Logger(2, "func 'CreateData()' -> Terjadi error saat marshal data (data->array json).")
			return
		} else {
			fmt.Fprint(response, string(extractData))
			Utils.Logger(3, "func 'CreateData()' -> Berhasil menambahkan data.")
		}
	}
}

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Function CreateData
 * Fungsi ini saya buat untuk menampilakan semua data yang ada di array of struct 'data' (line code: 50). 
 *
 * @Params response http.ResponseWriter
 *	-> params ini akan mengembalikan proses dari handle request user
 *
 * @Params request *http.Request
 *	-> params ini akan menghandle request body user, (formvalue Nama, NIM, Alamat)
 */
func ReadData(response http.ResponseWriter, request *http.Request) {
	// Set respon MIME type json, bertujuan untuk memberi tau user bahwa output ini merupakan json.
	response.Header().Set("Content-Type", "application/json")

	/* section IF untuk mengecek apakah data (array of struct) masih kosong ? isEmpty -> handle
	 * maka akan return code 404 Not Found
	 * mengemblikan respon dalam bentuk json dengan pesan yang tertera di code saya. 
	 * menampilkan log di console untuk mengetahui apa yang telah terjadi (untuk developer) */
	if len(data) == 0 {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprint(response, `{"status":"error", "message":"silahkan tambah data terlebih dahulu di route /nama"}`)
		Utils.Logger(1, "func 'ReadData()' -> Data dalam array of struct masih kosong.")
		return
	}

	/* Selanjutnya, saya membuat code untuk mengekstrak data ke dalam array (json)
	 * dengan menggunakan function marshal yang memiliki 2 return variable
	 * 1. return untuk datanya, saya menggunakan extractData
	 * 2. return untuk error, saya menggunakan err
	 *
	 * pada logicnya, ketika err ! nil (kosong) / maksudnya variable error terisi,
	 * maka program akan lanjut ke section IF yaitu handle error
	 * dengan logic : return code 500 - Internal Server Error
	 *		  mengembalikan respon dalam bentuk JSON ke user / client
	 *		  mengeluarkan output error dalam bentuk log di console
	 *
	 * jika proses marshal aman, lanjut masuk ke section ELSE
	 * dengan logic : saya tidak set return code, karena secara default sudah 200 OK
	 * 		  saya tampilkan response semua data (yang ada pada array of struct) dalam bentuk json,
	 * 		  saya tampilkan output di console untuk log history
	 */
	if extractData, err := json.Marshal(data); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(response, `{"status":"error", "message":"terjadi kesalahan pada server."}`)
			Utils.Logger(2, "func 'ReadData()' -> Terjadi error saat marshal data (data->array json).")
		return
	} else {
		fmt.Fprint(response, string(extractData))
		Utils.Logger(3, "func 'ReadData()' -> Berhasil menampilkan semua data.")
	}
}
