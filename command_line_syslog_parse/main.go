package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	path := flag.String("path", "/var/log/syslog", "The path to the log that should be analized")
	level := flag.String("level", "ERROR", "The log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL")

	flag.Parse()

	// Open a resource
	f, err := os.Open(*path)
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
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
