package jsonarray_test

import (
	"testing"

	"github.com/brokeyourbike/jsonarray"
	"github.com/stretchr/testify/assert"
)

type Dummy struct {
	Values []string
}

func TestJSONArray(t *testing.T) {
	a1 := jsonarray.JSONArray[int]{1, 2, 3}
	a1 = append(a1, 4)
	assert.Len(t, a1, 4)

	a2 := jsonarray.JSONArray[string]{"a", "b", "c"}
	a2 = append(a2, "d")
	assert.Len(t, a2, 4)

	d := Dummy{Values: a2}
	assert.Len(t, d.Values, 4)
}
