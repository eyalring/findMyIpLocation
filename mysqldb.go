package main

type mysqlDB struct {
	db DatabaseConnector
}

func newMySqlDB(connString string) (*DatabaseConnector, error) {
	mySqlDBInstance := &mysqlDB{db: DatabaseConnector{DB: connString}}
	db := &DatabaseConnector{
		DB:     connString,
		Finder: mySqlDBInstance,
	}
	return db, nil
}

func (conn *mysqlDB) FindLocation(ip string) LocationFound {
	return LocationFound{loc: Location{country: "Unknown", city: "Unknown"}, found: false}
}