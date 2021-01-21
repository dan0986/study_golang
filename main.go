package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func main() {
	if currentTime, error := ntp.Time("localhost"); error != nil {
		fmt.Printf("Current time is: %s", currentTime)
	} else {
		log.Fatal(error)
	}
}
