package virtual_constructor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	exemplar, present := Create("dummy field")
	assert.True(t, present)
	assert.IsType(t, Dummy{}, exemplar)

	assert.Equal(t, "field", exemplar.(Dummy).Field)
}
