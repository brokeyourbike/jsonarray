package jsonarray_test

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brokeyourbike/jsonarray"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
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
	dsn := os.Getenv("GORM_DSN")
	dialect := os.Getenv("GORM_DIALECT")

	log.Printf("testing %s...", strings.ToLower(dsn))

	switch dialect {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	}

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}

func TestSlice(t *testing.T) {
	type UserWithJSON struct {
		gorm.Model
		Name string
		Tags jsonarray.Slice[int]
	}

	DB.Migrator().DropTable(&UserWithJSON{})
	assert.NoError(t, DB.Migrator().AutoMigrate(&UserWithJSON{}))

	// create
	users := []UserWithJSON{{
		Name: "json-1",
		Tags: jsonarray.Slice[int]{1, 2},
	}, {
		Name: "json-2",
		Tags: jsonarray.Slice[int]([]int{3, 4, 5}),
	}, {
		Name: "json-3",
		Tags: jsonarray.Slice[int]{},
	}}
	assert.NoError(t, DB.Create(&users).Error)

	// select
	var result UserWithJSON
	assert.NoError(t, DB.First(&result, users[0].ID).Error)
	assert.Equal(t, users[0].Name, result.Name)
	assert.Equal(t, users[0].Tags[0], result.Tags[0])

	// first or create
	assign := UserWithJSON{
		Tags: jsonarray.Slice[int]{8, 9},
	}
	assert.NoError(t, DB.Where(&UserWithJSON{Name: "json-1"}).Assign(assign).FirstOrCreate(&UserWithJSON{}).Error)

	// update
	update := UserWithJSON{
		Tags: jsonarray.Slice[int]{7, 6},
	}
	var result3 UserWithJSON
	result3.ID = 1
	assert.NoError(t, DB.Model(&result3).Updates(update).Error)
}

func TestDummySlice(t *testing.T) {
	type Dummy struct {
		Values []string
	}

	a1 := jsonarray.Slice[int]{1, 2, 3}
	a1 = append(a1, 4)
	assert.Len(t, a1, 4)

	a2 := jsonarray.Slice[string]{"a", "b", "c"}
	a2 = append(a2, "d")
	assert.Len(t, a2, 4)

	d := Dummy{Values: a2}
	assert.Len(t, d.Values, 4)
}
