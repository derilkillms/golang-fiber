package unit_test

import (
	"fmt"
	"golang-fiber/database"
	"golang-fiber/database/migration"
	"testing"
)

func TestConnecting(t *testing.T) {
	//INITIAL DATABASE
	database.DatabaseInit()
	fmt.Println("test")
}

func TestRunMigration(t *testing.T) {
	//INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigration()
}
