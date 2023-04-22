package main

import (
	"github.com/cezarine/API-GO-GIN/database"
	"github.com/cezarine/API-GO-GIN/routes"
)

func main() {
	database.ConectaBanco()
	routes.HandRequests()
}
