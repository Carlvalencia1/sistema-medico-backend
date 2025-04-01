package application

import (
	"errors"
	"smartvitals/src/feautures/users/domain"
	"smartvitals/src/feautures/users/domain/entities"
)

type UpdateUserUseCase struct {
	userRepository domain.IUser
}

func NewUpdateUserUseCase(userRepository domain.IUser) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserUseCase) Run(user *entities.User) (*entities.UserResponse, error) {
	// No permitir que se convierta un usuario normal en admin si ya existe uno
	if user.Rol == "admin" {
		existingAdmin, _ := uc.userRepository.GetByUsername("admin")
		if existingAdmin != nil && existingAdmin.ID != user.ID {
			return nil, errors.New("no puedes asignar otro Super Admin")
		}
	}

	return uc.userRepository.Update(user)
}