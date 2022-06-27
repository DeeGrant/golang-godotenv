package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	fmt.Println("")

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}

	// Check environment variable
	token, ok := os.LookupEnv("APP_TOKEN")
	if !ok {
		fmt.Printf("Token not found")
	}
	fmt.Println("App Token: ", token)

	// Get environment variable
	fmt.Println("DB String: ", os.Getenv("DB_STRING"))

	fmt.Println("")

	// Set an environment variable
	err = os.Setenv("NEW_ENV_VAR", "SomethingNew")
	fmt.Println("New War: ", os.Getenv("NEW_ENV_VAR"))
	fmt.Println("")

	// Lead environment variables into a map
	envMap, err := godotenv.Read(".env")
	if err != nil {
		fmt.Printf("Could not load .env file into map[string]string\n")
		os.Exit(1)
	}

	fmt.Println("App Token: ", envMap["APP_TOKEN"])
	fmt.Println("Token:     ", envMap["DB_STRING"])
	fmt.Println("")

	// Print out all environmental variables
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s\n", pair[0])
	}
}
