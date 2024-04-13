package main

import (
	"fmt"
	"os"
	"strings"
)



func parseConnectionStrings() (map[string]string, error) {
	dbs := os.Getenv("ACTIVE_DATA_STORE_CONNECTION_STRINGS")
	dbEntries := strings.Split(dbs, ",")
	connStrings := make(map[string]string)

	for _, entry := range dbEntries {
		parts := strings.Split(entry, "~")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid connection string format: %s", entry)
		}
		connStrings[parts[0]] = parts[1]
	}

	return connStrings, nil
}

func initDb() (*DatabaseConnector, error) {
	connStrings, parseErr := parseConnectionStrings()
	if parseErr != nil {
		return nil, parseErr
	}

	activeDataStore := os.Getenv("ACTIVE_DATA_STORE")

	connString, ok := connStrings[activeDataStore]
	if !ok {
		return nil, fmt.Errorf("no connection string found for db type: %s", activeDataStore)
	}

	var err error
	var dataStore *DatabaseConnector
	switch activeDataStore {
	case "mysql":
		dataStore, err = newMySqlDB(connString)
	case "postgres":
		dataStore, err = newPostgresDB(connString)
	case "csv":
		dataStore, err = newCsvDB(connString)
	default:
		return nil, fmt.Errorf("unsupported active data store: %s", activeDataStore)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to active data store: %v", err)
	}
	return dataStore, nil
}
