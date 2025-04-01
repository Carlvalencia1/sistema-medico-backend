package application

import (
	"smartvitals/src/feautures/users/domain"
	"smartvitals/src/feautures/users/domain/entities"
)

type GetUserByIDUseCase struct {
	userRepository domain.IUser
}

func NewGetUserByIDUseCase(userRepository domain.IUser) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{userRepository: userRepository}
}

func (uc *GetUserByIDUseCase) Run(id int32) (*entities.UserResponse, error) {
	return uc.userRepository.GetByID(id)
}