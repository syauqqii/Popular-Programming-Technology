/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Set package program
 * Untuk memberitaukan bahwa ini program ini dari package Utils. */
package Utils

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Import Package / Library
 * Untuk memberikan fitur / fungsi tambahan yang lebih fleksible untuk program ini. */
import (
	/* library color dari github.com/fatih ini berfungsi untuk pewarnaan pada console
	 * berfungsi untuk mengkelompokkan supaya mudah, log apa saja yang error / success / dll. */
        "github.com/fatih/color"
	
	/* library OS disini saya gunakan untuk exit -> os.Exit(0) dengan return code 0 
	 * Karena ini log, jadi error sudah di print, tidak perlu return code 1. */
	"os"
	
	/* Library TIME disini saya gunakan untuk mendapatkan waktu sekarang yang valid.
	 * pada penerapannya saya menggunakan variable now untuk menampung nilai waktu saya. */
        "time"
)

/* Created by Dimas Syauqi Syafa - 2502004405. */

/* Function Logger
 * Fungsi ini saya buat untuk menampilkan custom log history. 
 *
 * @Params opt int    
 *	-> 1. INFO | 2. ERROR | 3. SUCCESS | 4. FATAL
 *	  (HiCyan) | (HiRed)  | (HiGreen)  | (Red)
 *
 * @Params msg string
 *	-> {pesan} / {message}
 */
func Logger(opt int, msg string) {
	// Set variable now untuk mendapatkan waktu valid saat variable di set.
        now := time.Now()

	// Set variable warna, untuk mempermudah pemanggilan.
        white := color.New(color.FgHiWhite); red    := color.New(color.FgHiRed)
        green := color.New(color.FgHiGreen); yellow := color.New(color.FgHiYellow)
        cyan  := color.New(color.FgHiCyan);  fatal  := color.New(color.FgRed)

	/* Keluarkan output berformat [hh:mm:ss yyyy/mm/dd] dengan '[' & ']' berwarna putih
	 * sedangkan string format dari variable 'now' saya outputkan dengan warna kuning. */
        white.Printf(" [%s] ", yellow.SprintFunc()(now.Format("15:04:05 2006/01/02")))

	/* Logic Pewarnaan
	 * 1. INFO berwanra 'cyan' terang
	 * 2. ERROR berwarna 'merah' terang
	 * 3. SUCCESS berwarna 'hijau' terang
	 * 4. FATAL berwanra 'merah'
	 *	-> untuk FATAL (@params opt int -> 4),
	 *	   maka program akan menampilkan pesan & exit.
	 */
        if opt == 1 {
                cyan.Printf("INFO")
        } else if opt == 2{
                red.Printf("ERROR")
        } else if opt == 3 {
                green.Printf("SUCCESS")
        } else if opt == 4 {
                fatal.Printf("FATAL")
                white.Println(": "+msg)
                os.Exit(0)
        }

	/* Menampilkan pesan dengan format  : (": {pesan}") ini bertujuan untuk 
	 * warna pesan tidak ikut berubah dengan variable 'opt' diatas (tetap putih). */
        white.Println(": "+msg)
}
