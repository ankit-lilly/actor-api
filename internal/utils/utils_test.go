package utils

import "testing"
import "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	assert.Equal(t, 3, result)
}
