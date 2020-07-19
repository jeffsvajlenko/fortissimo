package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Song holds the schema definition for the Song entity.
type Song struct {
	ent.Schema
}

// Fields of the Song.
func (Song) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
		field.String("hash"),
		field.String("title").
			Optional(),
		field.String("title_sort").
			Optional(),
		field.JSON("artists", []string{}).
			Optional(),
		field.String("first_artist").
			Optional(),
		field.String("first_artist_sort").
			Optional(),
		field.String("first_album_artist").
			Optional(),
		field.String("first_album_artist_sort").
			Optional(),
		field.String("album_artist").
			Optional(),
		field.String("album").
			Optional(),
		field.String("publisher").
			Optional(),
		field.String("first_composer").
			Optional(),
		field.String("composers").
			Optional(),
		field.String("conductor").
			Optional(),
		field.String("genre").
			Optional(),
		field.String("grouping").
			Optional(),
		field.Uint32("year").
			Optional(),
		field.Uint32("track_number").
			Optional(),
		field.Uint32("of_track_number").
			Optional(),
		field.Uint32("disk_number").
			Optional(),
		field.Uint32("of_disk_number").
			Optional(),
		field.Uint32("duration").
			Optional(),
		field.Uint32("play_count").
			Default(0),
		field.Uint32("skipped_count").
			Default(0),
		field.String("comment").
			Optional(),
		field.Uint32("beats_per_minute").
			Optional(),
		field.String("copyright").
			Optional(),
		field.Time("date_tagged").
			Optional(),
		field.String("description").
			Optional(),
		field.String("first_composer_sort").
			Optional(),
		field.String("artists_sort").
			Optional(),
		field.String("lyrics").
			Optional(),
		field.String("initial_key").
			Optional(),
		field.String("isrc").
			Optional(),
		field.String("subtitle").
			Optional(),
		field.String("music_brainz_artist_id").
			Optional(),
		field.String("music_brainz_disc_id").
			Optional(),
		field.String("music_brainz_release_artist_id").
			Optional(),
		field.String("music_brainz_release_country").
			Optional(),
		field.String("music_brainz_release_group_id").
			Optional(),
		field.String("music_brainz_release_id").
			Optional(),
		field.String("music_brainz_release_status").
			Optional(),
		field.String("music_brainz_release_type").
			Optional(),
		field.String("music_brainz_track_id").
			Optional(),
		field.String("music_ip_id").
			Optional(),
		field.String("remixed_by").
			Optional(),
		field.Float("replay_gain_album_gain").
			Optional(),
		field.Float("replay_gain_album_peak").
			Optional(),
		field.Float("replay_gain_track_gain").
			Optional(),
		field.Float("replay_gain_track_peak").
			Optional(),
		field.String("mime_type").
			Optional(),
		field.Time("created_date").
			Optional(),
		field.Time("modified_date").
			Optional(),
	}
}

// Edges of the Song.
func (Song) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
	}
}
