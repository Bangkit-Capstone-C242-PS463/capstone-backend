package main

import (
	"flag"
	"fmt"
	"log"

	"capstone-backend/internal/server"
	"capstone-backend/utils"
)

func main() {
	envPtr := flag.String("env", "dev", "Specify the environment: dev or prod")
	flag.Parse()

	envName := utils.DetermineEnvName(*envPtr)
	projectEnvName := fmt.Sprintf(".env.%s", envName)
	if envName == "prod" {
		fmt.Println("Running in production mode")
	} else {
		fmt.Println("Running in development mode")
	}

	// Load environment variables
	isEnvLoaded := utils.LoadEnvFile(projectEnvName)
	if isEnvLoaded {
		log.Printf("Environment file %s loaded successfully", projectEnvName)
	} else {
		fmt.Printf("Environment file %s not found\n", projectEnvName)
	}

	// Additional setup
	server.Setup(true, nil)
}
