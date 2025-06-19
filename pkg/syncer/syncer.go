package syncer

import "github.com/castlele/lalasync/pkg/storage"

type Syncer interface {
	Save(userName string, songs []*storage.SongModel) error
	Load(userName string) ([]*storage.SongModel, error)
}

type syncerImpl struct {
	repo storage.SongRepo
}

func NewSyncer(repo storage.SongRepo) Syncer {
	return &syncerImpl{
		repo: repo,
	}
}

func (s *syncerImpl) Save(userName string, songs []*storage.SongModel) error {
	return nil
}

func (s *syncerImpl) Load(userName string) ([]*storage.SongModel, error) {
	return nil, nil
}
