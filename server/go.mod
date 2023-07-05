module github.com/jeffsvajlenko/fortissimo/server

go 1.14

replace github.com/jeffsvajlenko/fortissimo/api/go => ../api/go

require (
	github.com/dhowden/tag v0.0.0-20200412032933-5d76b8eaae27
	github.com/facebookincubator/ent v0.2.7
	github.com/golang/protobuf v1.5.2
	github.com/jeffsvajlenko/fortissimo/api/go v0.0.0
	github.com/lib/pq v1.2.0
	github.com/mattn/go-sqlite3 v1.13.0
	google.golang.org/grpc v1.53.0
)
