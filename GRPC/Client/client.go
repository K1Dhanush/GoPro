package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	//Establishing a Connection with the gRPC(STUB) server
	clientCon, err := grpc.Dial(":5050")
	if err != nil {
		log.Fatalf("Not connected to the server %v", err)
	}

	//Creating a client for the Example -- service -using the connection
	client := generated.NewExampleClient(clientCon)

}
