package main

import (
	"log"
	"os"
	"time"
)

func main() {
	for {
		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			log.Fatal("DB_NAME environment variable not set")
		}

		log.Printf("DB_NAME: %s", dbName)

		waitDuration := 60 * time.Minute
		log.Printf("Waiting for %s...", waitDuration)
		time.Sleep(waitDuration)

		log.Println("Done waiting!")
	}
}

