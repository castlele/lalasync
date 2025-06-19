package storage

import (
	"reflect"
	"testing"
)

func TestSongRepo(t *testing.T) {
	t.Run("GetSongByName returns song when present", func(t *testing.T) {
		db := SongMemDB{
			"track1": {Name: "track1"},
		}
		sut := NewSongRepo(db)

		result := sut.GetSongByName("track1")

		if result == nil || result.Name != "track1" {
			t.Errorf("expected song 'track1', got %+v", result)
		}
	})

	t.Run("GetSongByName returns nil when song not found", func(t *testing.T) {
		db := SongMemDB{}
		sut := NewSongRepo(db)

		result := sut.GetSongByName("nonexistent")

		if result != nil {
			t.Errorf("expected nil, got %+v", result)
		}
	})

	t.Run("GetUserSongs returns only user's songs", func(t *testing.T) {
		db := SongMemDB{
			"track1": {Name: "track1", UserName: "alice"},
			"track2": {Name: "track2", UserName: "bob"},
			"track3": {Name: "track3", UserName: "alice"},
		}
		sut := NewSongRepo(db)

		result := sut.GetUserSongs("alice")

		expected := []*SongModel{
			{Name: "track1", UserName: "alice"},
			{Name: "track3", UserName: "alice"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %+v, got %+v", expected, result)
		}
	})

	t.Run("GetUserSongs returns empty slice when no songs found", func(t *testing.T) {
		db := SongMemDB{
			"track1": {Name: "track1", UserName: "bob"},
		}
		sut := NewSongRepo(db)

		result := sut.GetUserSongs("alice")

		if len(result) != 0 {
			t.Errorf("expected empty slice, got %+v", result)
		}
	})

	t.Run("SetSongForUser sets song and assigns username", func(t *testing.T) {
		db := SongMemDB{}
		sut := NewSongRepo(db)

		err := sut.SetSongForUser("alice", &SongModel{Name: "track1"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		saved := db["track1"]
		if saved == nil || saved.UserName != "alice" {
			t.Errorf("expected song with UserName 'alice', got %+v", saved)
		}
	})

	t.Run("SetSongForUser overwrites existing song", func(t *testing.T) {
		db := SongMemDB{
			"track1": {Name: "track1", UserName: "bob"},
		}
		sut := NewSongRepo(db)

		err := sut.SetSongForUser("alice", &SongModel{Name: "track1"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		saved := db["track1"]
		if saved.UserName != "alice" {
			t.Errorf("expected UserName to be 'alice', got '%s'", saved.UserName)
		}
	})

	t.Run("SetSongsForUser sets multiple songs with user", func(t *testing.T) {
		db := SongMemDB{}
		sut := NewSongRepo(db)

		songs := []*SongModel{
			{Name: "song1"},
			{Name: "song2"},
		}

		err := sut.SetSongsForUser("alice", songs)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, name := range []string{"song1", "song2"} {
			saved := db[name]
			if saved == nil || saved.UserName != "alice" {
				t.Errorf("expected song '%s' with UserName 'alice', got %+v", name, saved)
			}
		}
	})
}
