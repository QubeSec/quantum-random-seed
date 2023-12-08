package main

import (
	"embed"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed temp_hex.txt
var file embed.FS

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a route for the root endpoint
	r.GET("/", quantumSeed)

	// Run the server on port 80
	r.Run(":80")
}

func quantumSeed(c *gin.Context) {
	// Seed the random number generator with the current time

	desiredBytes, err := strconv.Atoi(c.Query("bytes"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rand.NewSource(time.Now().UnixNano())

	// Read the entire hex content of the file
	hexContent, err := file.ReadFile("temp_hex.txt")
	if err != nil {
		fmt.Println("Error reading hex content:", err)
		return
	}

	// Generate a random starting position in the byte slice
	randomStart := rand.Intn(len(hexContent) - desiredBytes)

	// Extract the desired number of bytes from the original slice
	randomBytes := hexContent[randomStart : randomStart+desiredBytes]

	c.String(200, string(randomBytes))
}
