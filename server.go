package main

import (
	"github.com/aodin/antisocial/network"
	"github.com/aodin/aspect"
	_ "github.com/lib/pq"
)

// http://localhost:8001/api/?lat=39.739167&lng=-104.984722

func main() {
	// Parse the config
	c := network.Parse()

	// Connect to the database
	db, err := aspect.Connect(
		"postgres",
		"host=localhost port=5432 dbname=antisocial user=postgres password=gotest",
	)
	if err != nil {
		panic(err)
	}

	// Start the server
	s := network.NewServer(c, db)
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
