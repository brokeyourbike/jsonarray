package jsonarray

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// Slice defined JSON data type, need to implements driver.Valuer, sql.Scanner interface
type Slice[T any] []T

// Value return json value, implement driver.Valuer interface
func (m Slice[T]) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	ba, err := m.MarshalJSON()
	return string(ba), err
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (m *Slice[T]) Scan(val interface{}) error {
	if val == nil {
		*m = make(Slice[T], 0)
		return nil
	}
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return fmt.Errorf("unsupported type: %T", val)
	}
	t := []T{}
	if err := json.Unmarshal(ba, &t); err != nil {
		return fmt.Errorf("failed to unmarshal JSON value: %w", err)
	}
	*m = t
	return nil
}

// MarshalJSON to output non base64 encoded []byte
func (m Slice[T]) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	t := ([]T)(m)

	data, err := json.Marshal(t)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON value: %w", err)
	}

	return data, nil
}

// UnmarshalJSON to deserialize []byte
func (m *Slice[T]) UnmarshalJSON(b []byte) error {
	t := []T{}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSONB value: %w", err)
	}

	*m = Slice[T](t)
	return nil
}

// GormDataType gorm common data type
func (m Slice[T]) GormDataType() string {
	return "jsonarray"
}

// GormDBDataType gorm db data type
func (Slice[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	case "sqlserver":
		return "NVARCHAR(MAX)"
	}
	return ""
}

func (m Slice[T]) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	data, _ := m.MarshalJSON()
	if db.Dialector.Name() == "mysql" {
		if v, ok := db.Dialector.(*mysql.Dialector); ok && !strings.Contains(v.ServerVersion, "MariaDB") {
			return gorm.Expr("CAST(? AS JSON)", string(data))
		}
	}
	return gorm.Expr("?", string(data))
}
