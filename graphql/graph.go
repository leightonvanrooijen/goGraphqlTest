package graph

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graphql-basics/gqldb"
)

// Reads and parses the schema from file
// Associates root resolver, checks for errors along the way
func parseSchema(path string, resolver interface{}) *graphql.Schema {
	bstr, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Couldn't get the Graphql schema file", err)
	}

	schemaString := string(bstr)
	parsedSchema, err := graphql.ParseSchema(
		schemaString,
		resolver,
	)

	if err != nil {
		log.Fatal("Couldn't parse the Graphql schema", err)
	}

	return parsedSchema
}

// Serves GraphQL Playground on root
// Serves GraphQL endpoint at /graphql
func ConnectGraphqQL(db gqldb.BasicDb) {
	playground := http.FileServer(http.Dir("graphql/graphqlPlayground"))

	http.Handle("/", playground)
	http.Handle("/graphql", &relay.Handler{
		Schema: parseSchema("./graphql/schema.graphql", &RootResolver{db: db}),
	})

	fmt.Println("serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
