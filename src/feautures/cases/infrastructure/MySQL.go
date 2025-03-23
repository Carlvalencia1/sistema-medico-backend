package infrastructure

import (
    "database/sql"
    "errors"
    "smartvitals/src/feautures/cases/domain"
    "fmt"
    "time"
)

type MySQL struct {
    db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
    return &MySQL{db: db}
}

func (m *MySQL) Create(expediente domain.MedicalCase) (domain.MedicalCase, error) {
    result, err := m.db.Prepare("INSERT INTO expedientes (id_usuario, temperatura, peso, estatura, ritmo_cardiaco, fecha_registro) VALUES (?, ?, ?, ?, ?, ?)")
    if err != nil {
        return domain.MedicalCase{}, err
    }
    defer result.Close()

    res, err := result.Exec(expediente.IDUsuario, expediente.Temperatura, expediente.Peso, expediente.Estatura, expediente.RitmoCardiaco, expediente.FechaRegistro)
    if err != nil {
        return domain.MedicalCase{}, err
    }

    id, err := res.LastInsertId()
    if err != nil {
        return domain.MedicalCase{}, err
    }

    expediente.IDExpediente = int(id)
    return expediente, nil
}

func (m *MySQL) GetAll() ([]domain.MedicalCase, error) {
    rows, err := m.db.Query("SELECT id_expediente, id_usuario, temperatura, peso, estatura, ritmo_cardiaco, fecha_registro FROM expedientes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var expedientes []domain.MedicalCase
    for rows.Next() {
        var expediente domain.MedicalCase
        var fechaRegistroBytes []byte // Para almacenar la fecha como []byte

        if err := rows.Scan(
            &expediente.IDExpediente,
            &expediente.IDUsuario,
            &expediente.Temperatura,
            &expediente.Peso,
            &expediente.Estatura,
            &expediente.RitmoCardiaco,
            &fechaRegistroBytes,
        ); err != nil {
            return nil, err
        }

        // Convertir `fechaRegistroBytes` a `time.Time`
        fechaRegistroStr := string(fechaRegistroBytes)
        expediente.FechaRegistro, err = time.Parse("2006-01-02 15:04:05", fechaRegistroStr)
        if err != nil {
            return nil, fmt.Errorf("error parsing fecha_registro: %w", err)
        }

        expedientes = append(expedientes, expediente)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return expedientes, nil
}


func (m *MySQL) GetById(id int) (domain.MedicalCase, error) {
    var expediente domain.MedicalCase

    err := m.db.QueryRow("SELECT id_expediente, id_usuario, temperatura, peso, estatura, ritmo_cardiaco, fecha_registro FROM expedientes WHERE id_expediente = ?", id).
        Scan(&expediente.IDExpediente, &expediente.IDUsuario, &expediente.Temperatura, &expediente.Peso, &expediente.Estatura, &expediente.RitmoCardiaco, &expediente.FechaRegistro)

    if err != nil {
        if err == sql.ErrNoRows {
            return domain.MedicalCase{}, errors.New("expediente not found")
        }
        return domain.MedicalCase{}, err
    }

    return expediente, nil
}

func (m *MySQL) Update(expediente domain.MedicalCase) (domain.MedicalCase, error) {
    stmt, err := m.db.Prepare("UPDATE expedientes SET id_usuario = ?, temperatura = ?, peso = ?, estatura = ?, ritmo_cardiaco = ?, fecha_registro = ? WHERE id_expediente = ?")
    if err != nil {
        return domain.MedicalCase{}, err
    }
    defer stmt.Close()

    _, err = stmt.Exec(expediente.IDUsuario, expediente.Temperatura, expediente.Peso, expediente.Estatura, expediente.RitmoCardiaco, expediente.FechaRegistro, expediente.IDExpediente)
    if err != nil {
        return domain.MedicalCase{}, err
    }

    return expediente, nil
}

func (m *MySQL) Delete(id int) error {
    stmt, err := m.db.Prepare("DELETE FROM expedientes WHERE id_expediente = ?")
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
        return errors.New("no se encontró el ID")
    }

    return nil
}

func (m *MySQL) GetByMedicalCase(medicalCaseId int) ([]domain.MedicalCase, error) {
    query := "SELECT id_expediente, id_usuario, temperatura, peso, estatura, ritmo_cardiaco, fecha_registro FROM expedientes WHERE id_expediente = ?"
    rows, err := m.db.Query(query, medicalCaseId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var medicalCases []domain.MedicalCase
    for rows.Next() {
        var medicalCase domain.MedicalCase
        if err := rows.Scan(
            &medicalCase.IDExpediente,
            &medicalCase.IDUsuario,
            &medicalCase.Temperatura,
            &medicalCase.Peso,
            &medicalCase.Estatura,
            &medicalCase.RitmoCardiaco,
            &medicalCase.FechaRegistro,
        ); err != nil {
            return nil, err
        }
        medicalCases = append(medicalCases, medicalCase)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    if len(medicalCases) == 0 {
        return nil, errors.New("no medical cases found")
    }

    return medicalCases, nil
}