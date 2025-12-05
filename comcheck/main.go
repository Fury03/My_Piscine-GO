package main

import (
	"fmt"
	"os"
)

func main() {
	fot _,arg := range os.Args[1:] {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
