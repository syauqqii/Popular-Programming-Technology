/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Set package program
 * Untuk memberitaukan bahwa ini program ini dari package Utils. */
package Utils

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Import Package / Library
 * Untuk memberikan fitur / fungsi tambahan yang lebih fleksible untuk program ini. */
import (
	"os"
	
	/* Runtime bawaan go lang yang dapat membuat sebuah instance 'Cmd' dan 'Run'
	 * Library ini saya gunakan untuk membuat object dari function Command() serta Run(). */
	"os/exec"
	
	/* Library runtime memiliki banyak function, salah satunya saya gunakan untuk mendeteksi os yang digunakan
	 * dengan cara runtime.GOOS, maka akan mengembalikan string nama OS nya. */
	"runtime"
)

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Function Logger
 * Fungsi ini saya buat untuk membersihkan console baik di CMD / Terminal.
 */
func ClearScreen(){
	/* Logic Program
	 * mendaptkan tipe OS dengan menjalankan perintah: runtime.GOOS
	 * 	-> jika return "windows" :
	 *		- run command di cmd: cmd (run cmd) -> /c (argument run cmd [cls]) -> cls (bersihkan console)
	 * 	-> jika return "linux" / "darwin" :
	 *		- run command di terminal: clear (membersihkan dengan cara seperti di newline sampai bersih, beda dengan windows)
	 */
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
  	} else {
		return
	}
}
