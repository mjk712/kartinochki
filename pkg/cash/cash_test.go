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

		wasInCache := c.Set("test1", dbimg)
		require.False(t, wasInCache)

		wasInCache = c.Set("abc1", dbimg2)
		require.False(t, wasInCache)

	})
}
