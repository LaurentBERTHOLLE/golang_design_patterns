package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	dummy := GetBuilder().
		Field2("field2").
		Field3(7).
		Field1("field1").
		Build()

	assert.Equal(t, "field1", dummy.Field1)
	assert.Equal(t, "field2", dummy.Field2)
	assert.Equal(t, 7, dummy.Field3)
}
