package application

import (
	"smartvitals/src/feautures/users/domain"
	"smartvitals/src/feautures/users/domain/entities"
)

type DeleteUserUsecase struct {
	db domain.IUser
}

func NewDeleteUserUsecase(db domain.IUser) *DeleteUserUsecase {
	return &DeleteUserUsecase{db: db}
}

func (u *DeleteUserUsecase) Run(user *entities.User) (*entities.UserResponse, error) {
	return u.db.Delete(user) // Return a response with the deleted user data
}
