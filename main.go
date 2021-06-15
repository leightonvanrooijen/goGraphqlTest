package main

import (
	// "fmt"

	// "github.com/graphql-basics/gqldb"
	"github.com/graphql-basics/gqldb"
	graph "github.com/graphql-basics/graphql"
	"github.com/graphql-basics/logger"
)

func main() {
	db := gqldb.ConnectDB()
	logger.Create()
	// db.Seed()
	graph.ConnectGraphqQL(db)
}