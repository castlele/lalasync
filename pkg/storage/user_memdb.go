package storage

import (
	"maps"
	"slices"
)

type UserMemDB map[string]*UserModel

func (db UserMemDB) GetAll() []*UserModel {
	return slices.Collect(maps.Values(db))
}

func (db UserMemDB) Get(key string) *UserModel {
	if value, ok := db[key]; ok {
		return value
	}

	return nil
}

func (db UserMemDB) Set(key string, value *UserModel) error {
	db[key] = value

	return nil
}
