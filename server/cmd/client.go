package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jeffsvajlenko/fortissimo/api/go/fortissimo"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	// Input Parameters
	port := flag.Int("port", 50000, "The port the gRPC server is listening on.")
	server := flag.String("server", "localhost", "The url of the server.")
	flag.Parse()

	var conn *grpc.ClientConn

	// creds
	//creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	//if err != nil {
	//	log.Fatalf("could not load tls cert: %s", err)
	//}

	// connect
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", *server, *port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// client
	client := fortissimo.NewFortissimoClient(conn)

	// Try to get songs
	songsClient, err := client.GetSongs(context.Background(), &fortissimo.GetSongsRequest{})
	if err != nil {
		log.Fatalf("Error2: %v\n", err)
	}
	for {
		msg, err := songsClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to recieve: %v", err)
		}
		fmt.Printf("Received %v\n", msg.Song)
	}
}