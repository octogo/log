package main

import "github.com/octogo/log"

func main() {
	log.Init()
	log.Println("Hello world!")
	log.Fatal("FATALITY")
}
