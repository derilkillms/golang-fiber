package unit_test

import (
	"fmt"
	"golang-fiber/database"
	"testing"
)

func TestConnecting(t *testing.T) {
	//INITIAL DATABASE
	database.DatabaseInit()
	fmt.Println("test")
}
