package main

type postgresDB struct {
	db DatabaseConnector
}

func newPostgresDB(connString string) (*DatabaseConnector, error) {
	postgresDBInstance := &postgresDB{db: DatabaseConnector{DB: connString}}
	db := &DatabaseConnector{
		DB:     connString,
		Finder: postgresDBInstance,
	}
	return db, nil
}

func (conn *postgresDB) FindLocation(ip string) LocationFound {
	return LocationFound{loc: Location{country: "Unknown", city: "Unknown"}, found: false}
}