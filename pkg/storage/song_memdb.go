package storage

import (
	"maps"
	"slices"
)

type SongMemDB map[string]*SongModel

func (db SongMemDB) GetAll() []*SongModel {
	return slices.Collect(maps.Values(db))
}

func (db SongMemDB) Get(key string) *SongModel {
	if value, ok := db[key]; ok {
		return value
	}

	return nil
}

func (db SongMemDB) Set(key string, value *SongModel) error {
	db[key] = value

	return nil
}
