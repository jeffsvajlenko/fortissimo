package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jeffsvajlenko/fortissimo/server/ent"
	"google.golang.org/grpc"
	"log"
	"net"

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

	// create a TCP listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	// Initialize database
	dbclient, err := database(*dbConnStr, context.Background())
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	// create instance of service
	s := services.FortissimoGrpcApiServer{
		DbClient: dbclient,
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach service to the server
	fortissimo.RegisterFortissimoServer(grpcServer, &s)

	// start the server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	fmt.Println("Fortissimo has ended.  Goodbye!")
}

func database(dbConnStr string, ctx context.Context) (*ent.Client, error) {
	dbclient, err := ent.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer dbclient.Close()

	// Run Database Setup/Migrations
	if err := dbclient.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources : %v", err)
		return nil, err
	}

	return dbclient, nil
}