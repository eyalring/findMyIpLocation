package main

type Location struct {
	country string
	city    string
}

type LocationFound struct {
	loc   Location
	found bool
}


type LocationFinder interface {
    FindLocation(ip string) LocationFound
}

type DatabaseConnector struct {
    DB           string
    Finder       LocationFinder
}
