package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Failed to load environment configuration file")
	}

	t := os.Getenv("TICKERS")
	topics := strings.Split(t, ",")
	fmt.Println(topics)
}
