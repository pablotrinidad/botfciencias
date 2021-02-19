package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pablotrinidad/botfciencias/crawler/lib"
)

func main() {
	majors, err := lib.GetMajors()
	if err != nil {
		fmt.Printf("OOOPS; %v", err)
		os.Exit(1)
	}

	for _, m := range majors {
		fmt.Printf("%s (%d)\n", m.Name, m.ID)
		for _, p := range m.Plans {
			fmt.Printf("\t%s (%d)\n", p.Name, p.ID)
		}
	}

	if _, err := lib.GetCourses("20211", majors[1].Plans[0].ID); err != nil {
		log.Fatal(err)
	}
}
