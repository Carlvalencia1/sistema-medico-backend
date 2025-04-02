package infraestructure

import (
	"errors"
	"database/sql"
    "smartvitals/src/feautures/users/domain/entities"
)

type MYSQL struct {
	db *sql.DB
}

func NewMysql(db *sql.DB) *MYSQL {
	return &MYSQL{
		db: db,
	}
}

func (mysql *MYSQL) Save(user *entities.User) (*entities.UserResponse, error) {
	result, err := mysql.db.Exec("INSERT INTO users (name, password, rol, email, username) VALUES (?,?,?,?,?)",
		user.Name, user.Password, user.Rol, user.Email, user.Username)
	if err != nil {
		return nil, err
	}
	id, errId := result.LastInsertId()
	if errId != nil {
		return nil, errId
	}
	userResponse := &entities.UserResponse{
		ID:       int32(id),
		Name:     user.Name,
		Rol:      user.Rol,
		Email:    user.Email,
		Username: user.Username,
	}
	return userResponse, nil
}

func (mysql *MYSQL) LogIn(userLog *entities.UserLogIn) (*entities.User, error) {
	var user entities.User

	err := mysql.db.QueryRow("SELECT id, name, password, rol, email, username FROM users WHERE username = ?",
		userLog.Username).Scan(&user.ID, &user.Name, &user.Password, &user.Rol, &user.Email, &user.Username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (mysql *MYSQL) Update(user *entities.User) (*entities.UserResponse, error) {
	result, err := mysql.db.Exec("UPDATE users SET name = ?, password = ?, rol = ?, email = ?, username = ? WHERE id = ?",
		user.Name, user.Password, user.Rol, user.Email, user.Username, user.ID)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	userResponse := &entities.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Rol:      user.Rol,
		Email:    user.Email,
		Username: user.Username,
	}
	return userResponse, nil
}

func (mysql *MYSQL) Delete(user *entities.User) (*entities.UserResponse, error) {
	result, err := mysql.db.Exec("DELETE FROM users WHERE id = ?", user.ID)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	userResponse := &entities.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Rol:      user.Rol,
		Email:    user.Email,
		Username: user.Username,
	}
	return userResponse, nil
}

func (mysql *MYSQL) GetAll() ([]entities.UserResponse, error) {
	rows, err := mysql.db.Query("SELECT id, name, rol, email, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.UserResponse
	for rows.Next() {
		var user entities.UserResponse

		if err := rows.Scan(&user.ID, &user.Name, &user.Rol, &user.Email, &user.Username); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (mysql *MYSQL) GetByID(id int32) (*entities.UserResponse, error) {
	var user entities.UserResponse

	err := mysql.db.QueryRow("SELECT id, name, rol, email, username FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Rol, &user.Email, &user.Username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (mysql *MYSQL) GetByUsername(username string) (*entities.UserResponse, error) {
    var user entities.UserResponse

    err := mysql.db.QueryRow("SELECT id, name, rol, email, username FROM users WHERE username = ?", username).
        Scan(&user.ID, &user.Name, &user.Rol, &user.Email, &user.Username)

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (mysql *MYSQL) GetAdmin() (*entities.UserResponse, error) {
    var user entities.UserResponse

    // Consulta para obtener el usuario con rol "admin"
    err := mysql.db.QueryRow("SELECT id, name, rol, email, username FROM users WHERE rol = 'admin' LIMIT 1").
        Scan(&user.ID, &user.Name, &user.Rol, &user.Email, &user.Username)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("no se encontró un usuario con rol de admin")
        }
        return nil, err
    }

    return &user, nil
}