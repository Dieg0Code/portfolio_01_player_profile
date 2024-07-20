package testutils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestModel es un modelo de GORM para fines de prueba.
type TestModel struct {
	gorm.Model
	Name string
}

func TestSetupTestDB(t *testing.T) {
	// Initialize test database.
	db := SetupTestDB(&TestModel{})
	defer func() {
		sqlDB, _ := db.DB()
		err := sqlDB.Close()
		if err != nil {
			t.Errorf("Error closing database connection: %v", err)
		}
	}()

	// Try to create a record and verify if it was created.
	testRecord := &TestModel{Name: "Test Name"}
	err := db.Create(testRecord).Error
	require.NoError(t, err, "Failed to create test record")

	// Try to retrieve the record and verify if it was retrieved.
	var retrievedRecord TestModel
	err = db.First(&retrievedRecord, "name = ?", "Test Name").Error
	require.NoError(t, err, "Failed to retrieve test record")

	require.Equal(t, "Test Name", retrievedRecord.Name, "Retrieved record name does not match")
}
