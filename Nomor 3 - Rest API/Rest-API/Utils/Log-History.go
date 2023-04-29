package Utils

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func Logger(options int, msg string){
	now := time.Now()	

	grn := color.New(color.FgHiGreen)
	cyn := color.New(color.FgHiCyan)
	red := color.New(color.FgHiRed)
	yel := color.New(color.FgHiYellow)

	fmt.Print(" [")
	yel.Printf("%s", now.Format("15:04:05 2006/01/02"))
	fmt.Print("] ")
	if options == 1 {
		cyn.Print("INFO")
	} else if options == 2 {
		red.Print("ERROR")
	} else if options == 3 {
		grn.Print("EXIT")
	}
	
	fmt.Printf(": %s\n", msg)
}