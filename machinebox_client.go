package main

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

// struct to store GraphQL response
type QueryResponseStruct struct {
	User struct {
		Publications struct {
			TotalDocuments int `json:"totalDocuments"`
		}
		Badges []struct {
			Name string `json:"name"`
		}
	}
}

func main() {

	// Create a GraphQL client to connect to the gql.hashnode server
	var client *graphql.Client = graphql.NewClient("https://gql.hashnode.com")

	// Define GraphQL query to fetch required data
	var query string = `
	{
		user (username: "hashnode"){
			publications (first: 3) {
				totalDocuments
			}
			badges {
				name
			}
		} 
	}
	`

	// Create a new GraphQL request with the query
	var request *graphql.Request = graphql.NewRequest(query)
	var response QueryResponseStruct

	// Send the request to the server and store the response
	var err error = client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err) // Stop execution if error occured in communication with the server
	}

	log.Print(response)
}

// About package
// - Simple to use
// Why need json here?
// Why response struct with public fields
