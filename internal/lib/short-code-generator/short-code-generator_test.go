package shortcodegenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Generate(t *testing.T) {
	t.Run("successfully generate short code", func(t *testing.T) {
		t.Parallel()

		length := 5
		generator := New(length)

		shortCode, err := generator.Generate()
		assert.NoError(t, err)
		assert.NotEmpty(t, shortCode)
		assert.Equal(t, length, len(shortCode))
	})

	t.Run("throws an error when length is negative value", func(t *testing.T) {
		t.Parallel()

		generator := New(-1)
		shortCode, err := generator.Generate()
		assert.Error(t, err)
		assert.Empty(t, shortCode)
	})
}
