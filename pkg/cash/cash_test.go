package cash

import (
	"image/jpeg"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("EmptyCache", func(t *testing.T) {
		c := NewLru(10)

		_, ok := c.Get("abc.jpg")
		require.False(t, ok)

		_, ok = c.Get("zyx.jpeg")
		require.False(t, ok)
	})

	t.Run("SimpleCache", func(t *testing.T) {
		c := NewLru(5)

		file, _ := os.Open("./kartinochki/pkg/cash/testDb/test1.jpg")
		dbimg, _ := jpeg.Decode(file)
		file2, _ := os.Open("./kartinochki/pkg/cash/testDb/abc1.jpg")
		dbimg2, _ := jpeg.Decode(file2)
		file3, _ := os.Open("./kartinochki/pkg/cash/testDb/xyz.jpg")
		dbimg3, _ := jpeg.Decode(file3)

		wasInCache := c.Set("test1", dbimg)
		require.False(t, wasInCache)

		wasInCache = c.Set("abc1", dbimg2)
		require.False(t, wasInCache)

		value, ok := c.Get("test1")
		require.True(t, ok)
		require.Equal(t, dbimg, value)

		value, ok = c.Get("abc1")
		require.True(t, ok)
		require.Equal(t, dbimg2, value)

		wasInCache = c.Set("test1", dbimg3)
		require.True(t, wasInCache)

		value, ok = c.Get("test1")
		require.True(t, ok)
		require.Equal(t, dbimg3, value)
	})

	t.Run("Purge and movetoDb logic", func(t *testing.T) {
		c := NewLru(2)

		file, _ := os.Open("./kartinochki/pkg/cash/testDb/test1.jpg")
		dbimg, _ := jpeg.Decode(file)
		file2, _ := os.Open("./kartinochki/pkg/cash/testDb/abc1.jpg")
		dbimg2, _ := jpeg.Decode(file2)
		file3, _ := os.Open("./kartinochki/pkg/cash/testDb/xyz.jpg")
		dbimg3, _ := jpeg.Decode(file3)

		c.Set("First", dbimg)
		c.Set("Second", dbimg2)

		value, ok := c.Get("First")
		require.Equal(t, dbimg, value)
		require.True(t, ok)

		value, ok = c.Get("Second")
		require.Equal(t, dbimg2, value)
		require.True(t, ok)

		c.Set("Third", dbimg3)
		value, ok = c.Get("Third")
		require.Equal(t, dbimg3, value)
		require.True(t, ok)

		_, ok = c.Get("First")
		require.False(t, ok)
	})
}
