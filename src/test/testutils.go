// testutils/testutils.go
package testutils

import (
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func RunTestsWithDatabase(m *testing.M, migrations ...interface{}) {
	Db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = Db.AutoMigrate(migrations...)
	if err != nil {
		panic("failed to migrate models")
	}

	// Run the tests
	code := m.Run()

	// Teardown: Clean up after all tests have run
	sqlDB, err := Db.DB()
	if err != nil {
		panic("failed to close database connection")
	}
	sqlDB.Close()

	// Exit with the status code returned by m.Run()
	os.Exit(code)
}
