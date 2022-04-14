package reader_test

import (
	parser "github.com/nestorov88/rss_reader/pkg/reader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReader(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		r, err := parser.Parse([]string{
			"https://www.blog.google/rss/",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, r)
	})

	t.Run("DuplicateUrls", func(t *testing.T) {
		r, err := parser.Parse([]string{
			"https://www.blog.google/rss/",
			"https://www.blog.google/rss/",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, r)
	})

	t.Run("EmptyUrls", func(t *testing.T) {
		r, err := parser.Parse([]string{})

		assert.Error(t, err)
		assert.Empty(t, r)
	})

	t.Run("MixedValidAndNonValidUrls", func(t *testing.T) {
		r, err := parser.Parse([]string{
			"htxtpxs://www.blog.google/rss/",
			"https://www.blog.google/rss/",
		})
		assert.Error(t, err)
		assert.NotEmpty(t, r)
	})
}
