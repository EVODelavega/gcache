package gcache_test

import (
	"testing"

	"github.com/EVODelavega/gcache"
	"github.com/stretchr/testify/require"
)

func TestConstructor(t *testing.T) {
	t.Run("create new cache with types, no data", testNoData)
	t.Run("create new cache with initial data", testWithData)
}

func testNoData(t *testing.T) {
	c := gcache.New[string, string]()
	require.NotNil(t, c)
	k, v := "foo", "bar"
	require.False(t, c.Has(k))
	c.Set(k, v)
	got, ok := c.Get(k)
	require.True(t, ok)
	require.Equal(t, v, got)
}

func testWithData(t *testing.T) {
	data := map[string]int{
		"one": 1,
		"two": 2,
	}
	c := gcache.NewWithVals(data)
	require.NotNil(t, c)
	for k, v := range data {
		require.True(t, c.Has(k))
		got, _ := c.Get(k)
		require.EqualValues(t, v, got)
	}
}
