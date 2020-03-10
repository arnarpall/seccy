package secret

import (
	"testing"

	"github.com/arnarpall/seccy/internal/store/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const encryptionKey = "arnar-rocks"

func TestSetSecret(t *testing.T) {
	t.Run("Successful", func(tt *testing.T) {
		c := NewClient(memory.InMemory())

		err := c.Set("favorite-band", "Aqua")
		require.NoError(tt, err)
	})
	t.Run("CanSetMultipleTimes", func(tt *testing.T) {
		c := NewClient(memory.InMemory())

		err := c.Set("favorite-band", "Aqua")
		require.NoError(tt, err)
		err = c.Set("favorite-band", "Aqua1")
		require.NoError(tt, err)
		err = c.Set("favorite-band", "Aqua2")
		require.NoError(tt, err)
	})
}

func TestGetSecret(t *testing.T) {
	t.Run("NoneExistent", func(tt *testing.T) {
		c := NewClient(memory.InMemory())

		_, err := c.Get("none-existent")
		require.Error(tt, err)
	})
	t.Run("Successfully", func(tt *testing.T) {
		c := NewClient(memory.InMemory())

		err := c.Set("favorite-band", "Aqua")
		require.NoError(tt, err)

		val, err := c.Get("favorite-band")
		require.NoError(tt, err)
		assert.Equal(tt, "Aqua", val)
	})
}