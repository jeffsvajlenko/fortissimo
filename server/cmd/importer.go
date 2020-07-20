package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jeffsvajlenko/fortissimo/server/ent"
	"github.com/jeffsvajlenko/fortissimo/server/services/library"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"log"
	"os"
	pathlib "path"
	"path/filepath"
)



func main() {
	dbDriver := flag.String("database", "sqlite3", "Database driver to use, sqlite3 or postgres.")
	dbConnStr := flag.String("dbconn", "file:db.s3db?_fk=1", "The postgresSQL connection string for the database.")
	flag.Parse()
	paths := flag.Args()

	fmt.Println("--Fortissimo Importer--")
	fmt.Printf("\tUsing database connection: %v\n", *dbConnStr)

	// Initialize database
	dbclient, err := Database(*dbConnStr, *dbDriver, context.Background())
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}
	defer dbclient.Close()

	// Library service
	library := library.New(dbclient)

	supportedFormats := map[string]bool{
		".aa": true,
		".aax": true,
		".aac": true,
		".aiff": true,
		".ape": true,
		".dsf": true,
		".flac": true,
		".m4a": true,
		".m4b": true,
		".m4p": true,
		".mp3": true,
		".mpc": true,
		".mpp": true,
		".ogg": true,
		".oga": true,
		".wav": true,
		".wma": true,
		".wv": true,
		".webm": true,
	}

	// Process Files
	for _, path := range paths {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			ext := pathlib.Ext(path)

			if _, ok := supportedFormats[ext]; ok {
				song, err := library.ImportSong(context.Background(), path)
				if err != nil {
					log.Fatalf("failed song import : %v", err)
				}
				log.Printf("Imported: %v\n", *song)
			}

			return nil
		})
	}
}

func Database(dbConnStr string, dbDriver string, ctx context.Context) (*ent.Client, error) {
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