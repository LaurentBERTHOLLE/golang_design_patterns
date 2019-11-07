package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleton(t *testing.T) {
	assert.NotNil(t, GetInstance())
	assert.True(t, GetInstance() == GetInstance())
}
