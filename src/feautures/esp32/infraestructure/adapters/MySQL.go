package adapters

import (
	"database/sql"
	"log"
	"smartvitals/src/feautures/esp32/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{
		conn: conn,
	}
}

func (misql *MySQL) Save(esp32 *entities.Esp32) (*entities.Esp32, error) {
	// Prepare the SQL statement for inserting the ESP32 data
	stmt, err := misql.conn.Prepare("INSERT INTO esp32 (id, id_propietario) VALUES (?, ?)")
	if err != nil {
		log.Printf("error preparing statement: %v", err)
		return &entities.Esp32{}, err
	}
	defer stmt.Close()

	// Execute the SQL statement with the ESP32 data
	_, errInsert := stmt.Exec(esp32.Id, esp32.IdPropietario)
	if errInsert != nil {
		log.Printf("error inserting into esp32: %v", errInsert)
		return &entities.Esp32{}, errInsert
	}

	return esp32, nil
}

func (misql *MySQL) GetByPropietario(id int) ([]entities.Esp32, error) {
	rows, err := misql.conn.Query("SELECT id, id_propietario FROM esp32 WHERE id_propietario = ?", id)
	if err != nil {
		log.Printf("error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var devices []entities.Esp32
	for rows.Next() {
		var Esp32 entities.Esp32
		if err := rows.Scan(&Esp32.Id, &Esp32.IdPropietario); err != nil {
			log.Printf("error scanning row: %v", err)
			return nil, err
		}
		devices = append(devices, Esp32)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error iterating rows: %v", err)
		return nil, err
	}
	return devices, nil
}

func (misql *MySQL) GetByID(id string) (*entities.Esp32, error) {
	row := misql.conn.QueryRow("SELECT id, id_propietario FROM esp32 WHERE id = ?", id)

	var Esp32 entities.Esp32
	err := row.Scan(&Esp32.Id, &Esp32.IdPropietario)
	if err != nil {
		log.Printf("error scanning row: %v", err)
		return &entities.Esp32{}, err
	}

	return &Esp32, nil
}

func (misql *MySQL) Delete(id string) error {
	_, err := misql.conn.Exec("DELETE FROM esp32 WHERE id = ?", id)
	if err != nil {
		log.Printf("error deleting esp32 with id %s: %v", id, err)
		return err
	}
	return nil
}
