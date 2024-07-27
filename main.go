package main

import "os"

func main() {

	portString := os.Getenv("PORT")
	println(portString)
}
