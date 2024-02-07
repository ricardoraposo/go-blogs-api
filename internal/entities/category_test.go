package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
    category := NewCategory("test")

    assert.NotNil(t, category.ID)
    assert.NotNil(t, category.Name)
}
