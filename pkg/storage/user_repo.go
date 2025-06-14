package storage

type UserDB DB[string, UserModel]

type UserRepo struct {
	db UserDB
}

func NewUserRepo(db UserMemDB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetUserByUserName(userName string) *UserModel {
	return r.db.Get(userName)
}

func (r *UserRepo) SetUser(user *UserModel) error {
	return r.db.Set(user.Name, user)
}
