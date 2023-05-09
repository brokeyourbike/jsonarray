package jsonarray_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/brokeyourbike/jsonarray"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error
	if DB, err = openTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v\n", err)
		os.Exit(1)
	}
}

func openTestConnection() (db *gorm.DB, err error) {
	dialect := os.Getenv("GORM_DIALECT")

	switch dialect {
	default:
		log.Println("testing sqlite3...")
		db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	}

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}

func TestJSONArray(t *testing.T) {
	type Dummy struct {
		Values []string
	}

	a1 := jsonarray.JSONArray[int]{1, 2, 3}
	a1 = append(a1, 4)
	assert.Len(t, a1, 4)

	a2 := jsonarray.JSONArray[string]{"a", "b", "c"}
	a2 = append(a2, "d")
	assert.Len(t, a2, 4)

	d := Dummy{Values: a2}
	assert.Len(t, d.Values, 4)
}

func TestJSONArrayI(t *testing.T) {
	type UserWithJSON struct {
		gorm.Model
		Name string
		Tags jsonarray.JSONArray[int]
	}

	DB.Migrator().DropTable(&UserWithJSON{})
	assert.NoError(t, DB.Migrator().AutoMigrate(&UserWithJSON{}))

	users := []UserWithJSON{{
		Name: "json-1",
		Tags: jsonarray.JSONArray[int]{1, 2},
	}, {
		Name: "json-2",
		Tags: jsonarray.JSONArray[int]([]int{3, 4, 5}),
	}}
	assert.NoError(t, DB.Create(&users).Error)

	var result UserWithJSON
	assert.NoError(t, DB.First(&result, users[0].ID).Error)
	assert.Equal(t, users[0].Name, result.Name)
	assert.Equal(t, users[0].Tags[0], result.Tags[0])
}
