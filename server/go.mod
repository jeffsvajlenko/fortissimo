module github.com/jeffsvajlenko/fortissimo/server

go 1.14

replace github.com/jeffsvajlenko/fortissimo/api/go => ../api/go

require (
	github.com/facebookincubator/ent v0.2.5
	github.com/jeffsvajlenko/fortissimo/api/go v0.0.0
	google.golang.org/grpc v1.30.0
)
