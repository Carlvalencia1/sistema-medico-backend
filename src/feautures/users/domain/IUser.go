package domain

import "smartvitals/src/feautures/users/domain/entities"

// IUser is the interface for user domain logic.
type IUser interface {
	Save(user *entities.User) (*entities.UserResponse, error)  // Crear usuario (pero no Super Admin)
	LogIn(userLog *entities.UserLogIn) (*entities.User, error) // Inicio de sesión
	Update(user *entities.User) (*entities.UserResponse, error) // Actualizar datos de usuario (sin cambiar el rol de Super Admin)
	Delete(user *entities.User) (*entities.UserResponse, error) // Eliminar usuario (excepto el Super Admin)
	GetAll() ([]entities.UserResponse, error) // Obtener todos los usuarios
	GetByID(id int32) (*entities.UserResponse, error) // Buscar usuario por ID
	GetByUsername(username string) (*entities.UserResponse, error) // Buscar usuario por username
	GetAdmin() (*entities.UserResponse, error) // Obtener el único Super Admin
}
