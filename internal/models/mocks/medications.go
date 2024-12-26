package mocks

import (
	"time"

	"github.com/leoromanini/medication_api/internal/models"
)

var mockMedication = &models.Medications{
	ID:         1,
	Name:       "Ibuprofen",
	Dosage:     "400 mg",
	Form:       "Capsule",
	Created:    time.Now(),
	LastUpdate: time.Now(),
}

type MedicationModel struct{}

func (model *MedicationModel) Create(name string, dosage string, form string) (int, error) {
	return 2, nil
}

func (model *MedicationModel) Get(id int) (*models.Medications, error) {
	switch id {
	case 1:
		return mockMedication, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (model *MedicationModel) Update(id int, name string, dosage string, form string) error {
	// TODO: To be implemented
	return nil
}

func (model *MedicationModel) Delete(id int) error {
	// TODO: To be implemented
	return nil
}

func (model *MedicationModel) List() ([]*models.Medications, error) {
	// TODO: To be implemented
	return []*models.Medications{}, nil
}
