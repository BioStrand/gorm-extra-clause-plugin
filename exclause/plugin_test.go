package exclause

import (
	"testing"

	extraClausePlugin "github.com/BioStrand/gorm-extra-clause-plugin"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	db, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}))
	err = db.Use(extraClausePlugin.New())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when registering the plugin", err)
	}

	_, ok := db.Plugins["ExtraClausePlugin"]
	if !ok {
		t.Errorf("Could not find ExtraClausePlugin after registration")
	}
}
