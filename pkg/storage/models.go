package storage

type UserModel struct {
	Name     string
	Password string
}

type SongModel struct {
	Name string
	Artist string
	Album string
	UserName string
	Hash string
	Content []byte
}
