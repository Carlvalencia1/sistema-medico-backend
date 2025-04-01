package domain

import "smartvitals/src/feautures/users/domain/entities"

// IUser is the interface for user domain logic.
type IUser interface {
	Save(user *entities.User) (*entities.UserResponse, error)
	LogIn(userLog *entities.UserLogIn) (*entities.User, error)
	Update(user *entities.User) (*entities.UserResponse, error)
	Delete(user *entities.User) (*entities.UserResponse, error)
	GetAll() ([]entities.UserResponse, error)
	GetByID(id int32) (*entities.UserResponse, error)
	GetByUsername(username string) (*entities.UserResponse, error)
}
