package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEntry(t *testing.T) {
	key := "test-key"
	value := "test-value"

	AddEntry(key, value, 1*time.Second)
	assert.Equal(t, value, GetEntry(key))
	time.Sleep(1 * time.Second)
	assert.Nil(t, GetEntry(key))

	AddEntry(key, value, 1*time.Second)
	DelEntry(key)
	assert.Nil(t, GetEntry(key))
}
