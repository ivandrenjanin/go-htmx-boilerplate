package user

import "database/sql"

func NewUserService(db *sql.DB) UserService {
	return &userService{newUserRepository(db)}
}

func newUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}
