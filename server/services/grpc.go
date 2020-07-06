package services

import "github.com/jeffsvajlenko/fortissimo/api/go/fortissimo"
import "context"

type FortissimoGrpcApiServer struct{}

func (s *FortissimoGrpcApiServer) AddSong(ctx context.Context, request *fortissimo.AddSongRequest) (*fortissimo.AddSongResponse, error) {
	panic("implement me")
}

func (s *FortissimoGrpcApiServer) RemoveSong(ctx context.Context, request *fortissimo.RemoveSongRequest) (*fortissimo.RemoveSongResponse, error) {
	panic("implement me")
}

func (s *FortissimoGrpcApiServer) GetSong(ctx context.Context, req *fortissimo.GetSongRequest) (*fortissimo.GetSongResponse, error) {
	panic("implement me")
}

func (s *FortissimoGrpcApiServer) GetSongs(*fortissimo.GetSongsRequest, fortissimo.Fortissimo_GetSongsServer) error {
	panic("implement me")
}
