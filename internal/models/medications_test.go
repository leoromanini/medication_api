package models

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// TODO: Currently just a single/simple integration test is implemented here in order to demonstrate the capability.
// Missing implement tests for others model functions.

func TestMedicationModelGet(t *testing.T) {
	tests := []struct {
		name             string
		medicationID     int
		medicationName   string
		medicationDosage string
		medicationForm   string
		expectError      bool
	}{
		{
			name:             "Valid ID 1",
			medicationID:     1,
			medicationName:   "Amoxicillin",
			medicationDosage: "250 mg",
			medicationForm:   "Tablet",
			expectError:      false,
		},
		{
			name:             "Valid ID 2",
			medicationID:     2,
			medicationName:   "Ozempic",
			medicationDosage: "0.25 mg",
			medicationForm:   "Pen",
			expectError:      false,
		},
		{
			name:             "Valid ID 3",
			medicationID:     3,
			medicationName:   "Citalopram",
			medicationDosage: "30 mg",
			medicationForm:   "Capsule",
			expectError:      false,
		},
		{
			name:         "Zero ID",
			medicationID: 0,
			expectError:  true,
		},
		{
			name:         "Non-existent ID",
			medicationID: 10,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			model := MedicationModel{db}

			medication, err := model.Get(tt.medicationID)

			if tt.expectError {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, medication.Name, tt.medicationName)
			assert.Equal(t, medication.Dosage, tt.medicationDosage)
			assert.Equal(t, medication.Form, tt.medicationForm)
			assert.Nil(t, err)

		})
	}
}
