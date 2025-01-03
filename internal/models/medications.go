package models

import (
	"database/sql"
	"errors"
	"time"
)

type Medications struct {
	ID         int
	Name       string
	Dosage     string
	Form       string
	Created    time.Time
	LastUpdate time.Time
}

type MedicationsModelInterface interface {
	Create(name string, dosage string, form string) (int, error)
	Get(id int) (*Medications, error)
	List() ([]*Medications, error)
	Update(id int, name string, dosage string, form string) error
	Delete(id int) error
}

type MedicationModel struct {
	DB *sql.DB
}

func (model *MedicationModel) List() ([]*Medications, error) {
	stmt := `SELECT id, name, dosage, form, created, last_update
			 FROM medications
			 WHERE is_active = 1`

	rows, err := model.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	medications := []*Medications{}

	for rows.Next() {
		m := &Medications{}

		err := rows.Scan(&m.ID, &m.Name, &m.Dosage, &m.Form, &m.Created, &m.LastUpdate)
		if err != nil {
			return nil, err
		}

		medications = append(medications, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return medications, nil
}

func (model *MedicationModel) Get(id int) (*Medications, error) {

	stmt := `SELECT id, name, dosage, form, created, last_update
			 FROM medications
			 WHERE is_active = 1 AND id = ?`

	row := model.DB.QueryRow(stmt, id)

	m := &Medications{}

	err := row.Scan(&m.ID, &m.Name, &m.Dosage, &m.Form, &m.Created, &m.LastUpdate)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return m, nil
}

func (model *MedicationModel) Create(name string, dosage string, form string) (int, error) {
	stmt := `INSERT INTO medications (name, dosage, form)
    		 VALUES(?, ?, ?)`

	result, err := model.DB.Exec(stmt, name, dosage, form)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (model *MedicationModel) Update(id int, name string, dosage string, form string) error {
	stmt := `UPDATE medications SET name=?, dosage=?, form=?
			 WHERE is_active = 1 AND id = ?`

	_, err := model.DB.Exec(stmt, name, dosage, form, id)
	if err != nil {
		return err
	}

	return nil
}

func (model *MedicationModel) Delete(id int) error {
	stmt := `UPDATE medications SET is_active=0
			 WHERE is_active = 1 AND id = ?`

	_, err := model.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
