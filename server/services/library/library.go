package library

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	ent "github.com/jeffsvajlenko/fortissimo/server/ent"
	song "github.com/jeffsvajlenko/fortissimo/server/ent/song"
	"io"
	"os"
	"time"
	taglib "github.com/dhowden/tag"
)

type Service interface {
	ImportSong(ctx context.Context, path string) (*ent.Song, error)
	CreateSong(ctx context.Context, path string, hash string) (*ent.Song, error)
	GetSong(ctx context.Context, id int64) (*ent.Song, error)
	GetSongs(ctx context.Context) ([]*ent.Song, error)
	RemoveSong(ctx context.Context, id int64) error
}

type EntLibraryService struct{
	db *ent.Client
}

func New(client *ent.Client) *EntLibraryService {
	return &EntLibraryService{
		db: client,
	}
}

func (s *EntLibraryService) ImportSong(ctx context.Context, path string) (*ent.Song, error) {
	var rsong *ent.Song

	hash, err := hash(path)
	if err != nil {
		return rsong, err
	}

	// try to find existing import
	rsong, err = s.db.Song.Query().Where(song.Path(path)).First(ctx)
	if ent.IsNotFound(err) {
		rsong, err = s.db.Song.Query().Where(song.Hash(hash)).First(ctx)
		if ent.IsNotFound(err) {
			rsong, err = s.CreateSong(ctx, path, hash)
			if err != nil {
				return rsong, err
			}
		} else if err != nil {
			return rsong, err
		}
	} else if err != nil {
		return rsong, err
	}

	f, err := os.Open(path)
	if err != nil {
		return rsong, err
	}

	finfo, err := f.Stat()
	if err != nil {
		return rsong, err
	}

	// Get tags
	if rsong.ModifiedDate.IsZero() || rsong.ModifiedDate.Before(finfo.ModTime()) {
		md, err := taglib.ReadFrom(f)
		if err != nil {
			return rsong, err
		}

		// TODO: Improve library or switch to taglib-sharp for better tag parsing
		rsong.Title = md.Title()
		rsong.Album = md.Album()
		rsong.Artists = []string{md.Artist()}
		rsong.AlbumArtist = md.AlbumArtist()
		rsong.Composers = md.Composer()
		rsong.Year = uint32(md.Year())
		rsong.Genre = md.Genre()
		tn, otn := md.Track()
		rsong.TrackNumber = uint32(tn)
		rsong.OfTrackNumber = uint32(otn)
		dn, odn := md.Disc()
		rsong.DiskNumber = uint32(dn)
		rsong.OfDiskNumber = uint32(odn)
		rsong.Lyrics = md.Lyrics()
		rsong.Comment = md.Comment()

		return s.db.Song.UpdateOne(rsong).Save(ctx)
	} else {
		return rsong, nil
	}

}

func (s *EntLibraryService) CreateSong(ctx context.Context, path string, hash string) (*ent.Song, error) {
	return s.db.Song.Create().
		SetPath(path).
		SetHash(hash).
		SetCreatedDate(time.Now()).
		SetPlayCount(0).
		SetSkippedCount(0).
		Save(ctx)
}

func (s *EntLibraryService) GetSong(ctx context.Context, id int64) (*ent.Song, error) {
	return s.db.Song.Get(ctx, id)
}

func (s *EntLibraryService) GetSongs(ctx context.Context) ([]*ent.Song, error) {
	return s.db.Song.Query().All(ctx)
}

func (s *EntLibraryService) AddSong(ctx context.Context) (*ent.Song, error) {
	return s.db.Song.Create().Save(ctx)
}

func (s *EntLibraryService) RemoveSong(ctx context.Context, id int64) error {
	return s.db.Song.DeleteOneID(id).Exec(ctx)
}

func hash(path string) (string, error) {
	var md5str string

	file, err := os.Open(path)
	if err != nil {
		return md5str, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return md5str, err
	}
	md5str = hex.EncodeToString(hash.Sum(nil))

	return md5str, nil
}
