package grpc

import (
	timestamppb "github.com/golang/protobuf/ptypes"
	"github.com/jeffsvajlenko/fortissimo/server/ent"
	"github.com/jeffsvajlenko/fortissimo/server/services/library"
)
import "github.com/jeffsvajlenko/fortissimo/api/go/fortissimo"
import "context"

type Server struct{
	libraryService library.Service
}

func New(libraryService library.Service) *Server {
	return &Server{
		libraryService: libraryService,
	}
}

func (s *Server) RemoveSong(ctx context.Context, request *fortissimo.RemoveSongRequest) (*fortissimo.RemoveSongResponse, error) {
	err := s.libraryService.RemoveSong(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &fortissimo.RemoveSongResponse{}, err
}

func (s *Server) GetSong(ctx context.Context, request *fortissimo.GetSongRequest) (*fortissimo.GetSongResponse, error) {
	song, err := s.libraryService.GetSong(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &fortissimo.GetSongResponse{
		Song: encodeSong(song),
	}, err
}

func (s *Server) GetSongs(request *fortissimo.GetSongsRequest, server fortissimo.Fortissimo_GetSongsServer) error {
	songs, err := s.libraryService.GetSongs(server.Context())
	if err != nil {
		return err
	}

	for _, element := range songs {
		server.Send(&fortissimo.GetSongsResponse{
			Song: encodeSong(element),
		})
	}

	return nil
}

func encodeSong(song *ent.Song) *fortissimo.Song {
	dateTagged, err := timestamppb.TimestampProto(song.DateTagged)
	if err != nil {
		dateTagged = nil
	}

	createdDate, err := timestamppb.TimestampProto(song.CreatedDate)
	if err != nil {
		createdDate = nil
	}

	modifiedDate, err := timestamppb.TimestampProto(song.ModifiedDate)
	if err != nil {
		modifiedDate = nil
	}

	return &fortissimo.Song{
		Id:                         song.ID,
		Title:                      song.Title,
		TitleSort:                  song.TitleSort,
		Artists:                    song.Artists,
		FirstArtist:                song.FirstArtist,
		FirstArtistSort:            song.FirstArtistSort,
		FirstAlbumArtist:           song.FirstAlbumArtist,
		FirstAlbumArtistSort:       song.FirstAlbumArtistSort,
		AlbumArtist:                song.AlbumArtist,
		Album:                      song.Album,
		Publisher:                  song.Publisher,
		FirstComposer:              song.FirstComposer,
		Composers:                  song.Composers,
		Conductor:                  song.Conductor,
		Genre:                      song.Genre,
		Grouping:                   song.Grouping,
		Year:                       song.Year,
		TrackNumber:                song.TrackNumber,
		OfTrackNumber:              song.OfTrackNumber,
		DiskNumber:                 song.DiskNumber,
		OfDiskNumber:               song.OfDiskNumber,
		Duration:                   song.Duration,
		PlayCount:                  song.PlayCount,
		SkippedCount:               song.SkippedCount,
		Comment:                   	song.Comment,
		BeatsPerMinute:             song.BeatsPerMinute,
		Copyright:                  song.Copyright,
		DateTagged:					dateTagged,
		Description:                song.Description,
		FirstComposerSort:          song.FirstComposerSort,
		ArtistsSort:                song.ArtistsSort,
		Lyrics:                     song.Lyrics,
		InitialKey:                 song.InitialKey,
		Isrc:                       song.Isrc,
		Subtitle:                   song.Subtitle,
		MusicBrainzArtistId:        song.MusicBrainzArtistID,
		MusicBrainzDiscId:          song.MusicBrainzDiscID,
		MusicBrainzReleaseArtistId: song.MusicBrainzReleaseArtistID,
		MusicBrainzReleaseCountry:  song.MusicBrainzReleaseCountry,
		MusicBrainzReleaseGroupId:  song.MusicBrainzReleaseGroupID,
		MusicBrainzReleaseId:       song.MusicBrainzReleaseID,
		MusicBrainzReleaseStatus:   song.MusicBrainzReleaseStatus,
		MusicBrainzReleaseType:     song.MusicBrainzReleaseType,
		MusicBrainzTrackId:         song.MusicBrainzTrackID,
		MusicIpId:                  song.MusicIPID,
		RemixedBy:                  song.RemixedBy,
		ReplayGainAlbumGain:        song.ReplayGainAlbumGain,
		ReplayGainAlbumPeak:        song.ReplayGainAlbumPeak,
		ReplayGainTrackGain:        song.ReplayGainTrackGain,
		ReplayGainTrackPeak:        song.ReplayGainTrackPeak,
		MimeType:                   song.MimeType,
		Path:                       song.Path,
		Hash:                       song.Hash,
		CreatedDate:                createdDate,
		ModifiedDate:               modifiedDate,
	}
}
