package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jeffsvajlenko/fortissimo/server/ent"
	fortissimoGrpc "github.com/jeffsvajlenko/fortissimo/server/services/grpc"
	"github.com/jeffsvajlenko/fortissimo/server/services/library"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	"github.com/jeffsvajlenko/fortissimo/api/go/fortissimo"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
	// Input Parameters
	port := flag.Int("port", 50000, "The port the server should listen on (gRPC).")
	dbDriver := flag.String("database", "sqlite3", "Database driver to use, sqlite3 or postgres.")
	dbConnStr := flag.String("dbconn", "file:db.s3db?_fk=1", "The postgresSQL connection string for the database.")
	certFile := flag.String("tlscert", "cert/server.crt", "Certificate file for TLS.")
	keyFile := flag.String("tlskey", "cert/server.key", "Key file for TLS.")
	flag.Parse()

	// Output
	fmt.Println("--Fortissimo Server--")
	fmt.Printf("\tListening on port %v\n", *port)
	fmt.Printf("\tUsing database connection: %v\n", *dbConnStr)
	fmt.Printf("\tCert file: %v\n", *certFile)
	fmt.Printf("\tCert file: %v\n", *keyFile)

	// create a TCP listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	// Initialize database
	dbclient, err := database(*dbConnStr, *dbDriver, context.Background())
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}
	defer dbclient.Close()

	// create instance of service
	s := fortissimoGrpc.New(library.New(dbclient))

	// Create TLS credentials
	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	// create a gRPC server object
	opts := []grpc.ServerOption{
		grpc.Creds(creds),
	}
	grpcServer := grpc.NewServer(opts...)

	// attach service to the server
	fortissimo.RegisterFortissimoServer(grpcServer, s)

	// start the server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	fmt.Println("Fortissimo has ended.  Goodbye!")
}

func database(dbConnStr string, dbDriver string, ctx context.Context) (*ent.Client, error) {
	dbclient, err := ent.Open(dbDriver, dbConnStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Run Database Setup/Migrations
	if err := dbclient.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources : %v", err)
		return nil, err
	}

	return dbclient, nil
}