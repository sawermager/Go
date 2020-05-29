package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Open a resource
	f, err := os.Open("myapp.log")
	if err != nil {
		log.Fatal(err)
	}
	// When main() exits, defer will close the file descirptor that we opened
	defer f.Close()

	// Open a "reader"
	r := bufio.NewReader(f)

	// Infinite for loop (no expelicit conditions used)
	for {

		// Grab next line of file (parsed by newline char)
		s, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		if strings.Contains(s, "ARP") {
			fmt.Println(s)
		}
	}
	panic(fmt.Sprintf("%v", err))
}
