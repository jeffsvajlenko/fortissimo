package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/jeffsvajlenko/fortissimo/api/go/fortissimo"
)

const port = ":50000"

type server struct{}

func (s *server) AddSong(ctx context.Context, request *fortissimo.AddSongRequest) (*fortissimo.AddSongResponse, error) {
	panic("implement me")
}

func (s *server) RemoveSong(ctx context.Context, request *fortissimo.RemoveSongRequest) (*fortissimo.RemoveSongResponse, error) {
	panic("implement me")
}

func (s *server) GetSong(ctx context.Context, req *fortissimo.GetSongRequest) (*fortissimo.GetSongResponse, error) {
	panic("implement me")
}

func (s *server) GetSongs(*fortissimo.GetSongsRequest, fortissimo.Fortissimo_GetSongsServer) error {
	panic("implement me")
}

func main() {
	fmt.Println("--Fortissimo Server--")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fortissimo.RegisterFortissimoServer(s, &server{})
	s.Serve(lis)
}
