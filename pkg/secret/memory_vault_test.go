package secret

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Run("Fails on no key", func(tt *testing.T) {
		vault := InMemory()
		_, err := vault.Get("none-existent")
		require.Error(tt, err)
	})
	t.Run("ExistingKey", func(tt *testing.T) {
		vault := InMemory()
		err := vault.Set("awesome", "stuff")
		require.NoError(tt, err)

		v, err := vault.Get("awesome")
		require.NoError(tt, err)
		assert.Equal(tt, "stuff", v)
	})
}