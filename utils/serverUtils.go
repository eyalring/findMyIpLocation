package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ensureEnvVars() {
	envVars := []string{"ACTIVE_DATA_STORE_CONNECTION_STRINGS", "ACTIVE_DATA_STORE", "PORT"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			fmt.Printf("Environment variable %s not set. Exiting.\n", envVar)
			os.Exit(1)
		}
	}
}

func initializeTimeGapsForService() int64 {
	gapString := os.Getenv("KEEP_GAP_OF_SECONDS_BETWEEN_CALLS")
	
	var err error
	keepGapOfSecondsBetweenCallsInt, err := strconv.Atoi(gapString)
	if err != nil {
		fmt.Printf("Error converting KEEP_GAP_OF_SECONDS_BETWEEN_CALLS to integer: %s. Using default value 1.\n", err)
		return int64(1)
	} else {
		return int64(keepGapOfSecondsBetweenCallsInt)
	}
	
}