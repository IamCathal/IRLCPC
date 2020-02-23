package main

import (
	"log"
	"os"
)

func main() {

	switch os.Args[1] {
	case "question10":
		question10()
		break
	case "question1":
		question1()
		break
	default:
		log.Fatal("Invalid qustion")
	}

}
