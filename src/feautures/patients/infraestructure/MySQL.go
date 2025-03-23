package infraestructure

import (
	"database/sql"
	//"errors"
	"fmt"
	"smartvitals/src/feautures/patients/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (m *MySQL) Create(paciente domain.Patients) (domain.Patients, error) {
	stmt, err := m.db.Prepare("INSERT INTO pacientes (nombre, apellido, edad, genero, numero_contacto) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return domain.Patients{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Edad, paciente.Genero, paciente.NumeroContacto)
	if err != nil {
		return domain.Patients{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Patients{}, err
	}

	paciente.IDUsuario = int(id)
	return paciente, nil
}

func (m *MySQL) GetAll() ([]domain.Patients, error) {
	rows, err := m.db.Query("SELECT id_usuario, nombre, apellido, edad, genero, numero_contacto FROM pacientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pacientes []domain.Patients
	for rows.Next() {
		var paciente domain.Patients
		if err := rows.Scan(&paciente.IDUsuario, &paciente.Nombre, &paciente.Apellido, &paciente.Edad, &paciente.Genero, &paciente.NumeroContacto); err != nil {
			return nil, err
		}
		pacientes = append(pacientes, paciente)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return pacientes, nil
}

func (m *MySQL) GetById(id int) (domain.Patients, error) {
	var expediente domain.Patients
	err := m.db.QueryRow("SELECT id_usuario, nombre, apellido, edad, genero, numero_contacto FROM pacientes WHERE id_usuario = ?", id).
		Scan(&expediente.IDUsuario, &expediente.Nombre, &expediente.Apellido, &expediente.Edad, &expediente.Genero, &expediente.NumeroContacto)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Patients{}, fmt.Errorf("no se encontró el paciente con ID %d", id)
		}
		return domain.Patients{}, err
	}

	return expediente, nil
}

func (m *MySQL) Update(paciente domain.Patients) (domain.Patients, error) {
	stmt, err := m.db.Prepare("UPDATE pacientes SET nombre = ?, apellido = ?, edad = ?, genero = ?, numero_contacto = ? WHERE id_usuario = ?")
	if err != nil {
		return domain.Patients{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Edad, paciente.Genero, paciente.NumeroContacto, paciente.IDUsuario)
	if err != nil {
		return domain.Patients{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return domain.Patients{}, err
	}

	if rowsAffected == 0 {
		return domain.Patients{}, fmt.Errorf("no se encontró el paciente con ID %d para actualizar", paciente.IDUsuario)
	}

	return paciente, nil
}

func (m *MySQL) Delete(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM pacientes WHERE id_usuario = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el paciente con ID %d para borrarlo", id)
	}

	return nil
}

func (m *MySQL) GetByPatient(patientId int) ([]domain.Patients, error) {
	rows, err := m.db.Query("SELECT id_usuario, nombre, apellido, edad, genero, numero_contacto FROM pacientes WHERE id_usuario = ?", patientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pacientes []domain.Patients
	for rows.Next() {
		var paciente domain.Patients
		if err := rows.Scan(&paciente.IDUsuario, &paciente.Nombre, &paciente.Apellido, &paciente.Edad, &paciente.Genero, &paciente.NumeroContacto); err != nil {
			return nil, err
		}
		pacientes = append(pacientes, paciente)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(pacientes) == 0 {
		return []domain.Patients{}, nil
	}
	return pacientes, nil
}
