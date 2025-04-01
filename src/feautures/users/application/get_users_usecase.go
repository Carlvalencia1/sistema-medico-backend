package application

import (
	"smartvitals/src/feautures/users/domain"
	"smartvitals/src/feautures/users/domain/entities"
)

type GetUsersUseCase struct {
	userRepository domain.IUser
}

func NewGetUsersUseCase(userRepository domain.IUser) *GetUsersUseCase {
	return &GetUsersUseCase{
		userRepository: userRepository,
	}
}

// GetAll retrieves all users from the repository and returns them as a slice of UserResponse.
func (uc *GetUsersUseCase) GetAll() ([]entities.UserResponse, error) {
	return uc.userRepository.GetAll()
}
