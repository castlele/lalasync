package storage

type UserDB DB[string, UserModel]

type UserRepo interface {
	GetUserByUserName(string) *UserModel
	SetUser(*UserModel) error
}

type userRepoImpl struct {
	db UserDB
}

func NewUserRepo(db UserDB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}

func (r *userRepoImpl) GetUserByUserName(userName string) *UserModel {
	return r.db.Get(userName)
}

func (r *userRepoImpl) SetUser(user *UserModel) error {
	return r.db.Set(user.Name, user)
}
