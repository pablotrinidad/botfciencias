package main

import (
	"fmt"
	"os"

	"github.com/pablotrinidad/botfciencias/crawler/lib"
)

func main() {
	if err := lib.Run(); err != nil {
		fmt.Printf("UUUUPPSSSSSS HUBO UN ERROR: %v",err)
		os.Exit(1)
	}
}