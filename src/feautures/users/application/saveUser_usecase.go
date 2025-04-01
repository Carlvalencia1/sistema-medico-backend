package application

import (
	"errors"
	"smartvitals/src/core/middlewares"
	"smartvitals/src/feautures/users/domain"
	"smartvitals/src/feautures/users/domain/entities"
)

type SaveUserUseCase struct {
	userRepository domain.IUser
}

func NewSaveUser(userRepository domain.IUser) *SaveUserUseCase {
	return &SaveUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *SaveUserUseCase) Run(user *entities.User) (*entities.UserResponse, error) {
	// Si el usuario que se intenta crear es un admin, verificar que no exista otro
	if user.Rol == "admin" {
		existingAdmin, _ := uc.userRepository.GetByUsername("admin") // Buscar si ya hay un admin
		if existingAdmin != nil {
			return nil, errors.New("ya existe un Super Admin en el sistema")
		}
	}

	// Hashear la contraseña antes de guardarla
	hashedPassword, err := middlewares.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	return uc.userRepository.Save(user)
}
