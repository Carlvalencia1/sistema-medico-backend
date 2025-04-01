package application

import (
	"errors"
	"fmt"
	"smartvitals/src/core/middlewares"
	"smartvitals/src/feautures/users/domain"
	"smartvitals/src/feautures/users/domain/entities"
)

type LogInUseCase struct {
	userRepository domain.IUser
}

func NewLogInUseCase(userRepository domain.IUser) *LogInUseCase {
	return &LogInUseCase{
		userRepository: userRepository,
	}
}

func (uc *LogInUseCase) Run(userLog *entities.UserLogIn) (*entities.Claims, error) {
	fmt.Printf("user: %v\n", userLog)

	// Obtener el usuario de la base de datos
	user, err := uc.userRepository.LogIn(userLog)
	if err != nil {
		return nil, err
	}

	// Verificar la contraseña
	err = middlewares.VerifyPassword(userLog.Password, user.Password)
	if err != nil {
		return nil, errors.New("redenciales inválidas")
	}

	// Crear los claims con el rol del usuario
	claims := &entities.Claims{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Rol:      user.Rol, // Si es "admin", tendrá control total
		Email:    user.Email,
	}

	return claims, nil
}
