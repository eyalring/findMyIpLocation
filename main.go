package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var keepGapOfSecondsBetweenCalls int64

func main() {
	ensureEnvVars()
	dbConnector, dbError := initDb()
	
    keepGapOfSecondsBetweenCalls = initializeTimeGapsForService()
	if dbError != nil || dbConnector == nil {
		fmt.Println("Error initializing database:", dbError)
		return
	}

	router := gin.Default()
	router.Use(maxCallsCheck(keepGapOfSecondsBetweenCalls))
	router.GET("/v1/find-country", ipInputValidation, gin.HandlerFunc(func(c *gin.Context) {
		calculateCountryFromIP(c, dbConnector)
	}))
	router.Run("localhost:" + os.Getenv("PORT"))
}



