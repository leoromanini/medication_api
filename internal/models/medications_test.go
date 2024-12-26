package models

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// TODO: Currently just a simple integration test is implemented here in order to demonstrate the capability.
// In a real scenario this would be more complex.

func TestMedicationModelGet(t *testing.T) {
	tests := []struct {
		name           string
		medicationID   int
		medicationName string
		expectError    bool
	}{
		{
			name:           "Valid ID 1",
			medicationID:   1,
			medicationName: "Amoxicillin",
			expectError:    false,
		},
		{
			name:           "Valid ID 2",
			medicationID:   2,
			medicationName: "Ozempic",
			expectError:    false,
		},
		{
			name:           "Valid ID 3",
			medicationID:   3,
			medicationName: "Citalopram",
			expectError:    false,
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
			assert.Nil(t, err)

		})
	}
}
