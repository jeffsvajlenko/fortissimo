package main

import (
	"context"
	"fmt"
	"github.com/jeffsvajlenko/fortissimo/server/ent"
	"google.golang.org/grpc"
	"log"
	"net"
	"flag"

	"github.com/jeffsvajlenko/fortissimo/api/go/fortissimo"
	"github.com/jeffsvajlenko/fortissimo/server/services"
	_ "github.com/lib/pq"
)

func main() {
// Input Parameters
	port := flag.Int("port", 50000, "The port the server should listen on (gRPC).")
	dbConnStr := flag.String("dbconn", "localhost", "The postgresSQL connection string for the database.")
	flag.Parse()
	// todo: add config file option

	fmt.Println("--Fortissimo Server--")
	fmt.Printf("\tListening on port %v\n", *port)
	fmt.Printf("\tUsing database connection: %v\n", *dbConnStr)

	// Establish Connection
	dbclient, err := ent.Open("postgres", *dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer dbclient.Close()
	ctx := context.Background()

	// Run Database Setup/Migrations
	if err := dbclient.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources : %v", err)
	}

	// Start gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fortissimo.RegisterFortissimoServer(s, &services.FortissimoGrpcApiServer{})
	s.Serve(lis)
}
