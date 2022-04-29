package cruder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyTypeInt32(t *testing.T) {
	val, err := AnyTypeInt32(interface{}(1))
	assert.Nil(t, err)
	assert.Equal(t, val, int32(1))

	val1, err := AnyTypeUint32(interface{}(1))
	assert.Nil(t, err)
	assert.Equal(t, val1, uint32(1))
}
