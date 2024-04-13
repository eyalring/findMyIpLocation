package main

import (
	"encoding/csv"
	"fmt"
	"os"
)


var ipMap = make(map[string]Location)

type csvDB struct {
	db DatabaseConnector
}


func newCsvDB(connString string) (*DatabaseConnector, error) {
    csvDBInstance := &csvDB{db: DatabaseConnector{DB: connString}}
    db := &DatabaseConnector{
        DB: connString,
        Finder: csvDBInstance,
    }
    return db, nil
}




func (conn *csvDB) FindLocation(ip string) LocationFound {
    conn.readCsvFile()
    if loc, ok := ipMap[ip]; ok {
        return LocationFound{loc: loc, found: true}
    }
    return LocationFound{loc: Location{country: "Unknown", city: "Unknown"}, found: false}
}



func (conn *csvDB) readCsvFile(){

	if len(ipMap) == 0 {
		conn.initDataFromCSV()
	}
}

func (conn *csvDB) initDataFromCSV() {
	var fileName = conn.db.DB
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:",fileName, err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) == 3 {
			ipMap[record[0]] = Location{
				country: record[1],
				city:    record[2],
			}
		}
	}
}
