package cruder

import (
	"github.com/google/uuid"
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

	val2, err := AnyTypeUUID(interface{}(uuid.UUID{}.String()))
	assert.Nil(t, err)
	assert.Equal(t, val2, uuid.UUID{})

	val3, err := AnyTypeUUID(interface{}(uuid.UUID{}))
	assert.Nil(t, err)
	assert.Equal(t, val3, uuid.UUID{})
}
