package main

import (
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	println(portString)
}
