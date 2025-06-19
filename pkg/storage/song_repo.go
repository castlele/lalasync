package storage

type SongDB DB[string, SongModel]

type SongRepo interface {
	GetSongByName(string) *SongModel
	GetUserSongs(string) []*SongModel
	SetSongForUser(string, *SongModel) error
	SetSongsForUser(string, []*SongModel) error
}

type songRepoImpl struct {
	db SongDB
}

func NewSongRepo(db SongDB) SongRepo {
	return &songRepoImpl{
		db: db,
	}
}

func (r *songRepoImpl) GetSongByName(name string) *SongModel {
	return r.db.Get(name)
}

func (r *songRepoImpl) GetUserSongs(userName string) []*SongModel {
	songs := r.db.GetAll()
	var userSongs []*SongModel

	for _, song := range songs {
		if song.UserName == userName {
			userSongs = append(userSongs, song)
		}
	}

	return userSongs
}

func (r *songRepoImpl) SetSongForUser(name string, song *SongModel) error {
	song.UserName = name

	return r.db.Set(song.Name, song)
}

func (r *songRepoImpl) SetSongsForUser(name string, songs []*SongModel) error {
	var err error

	for _, song := range songs {
		if err = r.SetSongForUser(name, song); err != nil {
			return err
		}
	}

	return nil
}
