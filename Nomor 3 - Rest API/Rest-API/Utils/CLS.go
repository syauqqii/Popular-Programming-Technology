package Utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen(){
	_os := runtime.GOOS

	if _os == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if _os == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
  	} else {
		fmt.Print()
	}
}
