package number_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vietta-net/agokit/currency"
	"github.com/vietta-net/agokit/number"
)

func TestNumber(t *testing.T) {
	t.Run("FormatByLocale", func(t *testing.T) {
		var lang = "vi"
		locale, err := currency.GetLocale(lang)
		assert.Nil(t, err)
		assert.NotNil(t, locale)

		actual := number.FormatByLocale(locale, 1234.5)
		assert.Equal(t, "1.234,50", actual)

	})

	t.Run("Format", func(t *testing.T) {
		var lang = "vi"


		actual, err := number.Format(lang, 1234.5)
		assert.Nil(t, err)
		assert.Equal(t, "1.234,50", actual)

	})
}