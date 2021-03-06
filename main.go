package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	const newYorkTimeZone = "America/New_York"
	newYorkLoc, err := time.LoadLocation(newYorkTimeZone)
	if err != nil {
		log.Printf("error: %+v\n", err)
		return
	}
	ny := time.Now().In(newYorkLoc)
	log.Println("ny", ny)
	log.Println("utc", ny.UTC())
}