package number_test

import (
	"github.com/leekchan/accounting"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vietta-net/agokit/number"
)

func TestNumber(t *testing.T) {

	t.Run("Positive", func(t *testing.T) {
		var lang = "vi"
		var  locale *accounting.Locale
		locale, err := number.GetLocale(lang)
		assert.Nil(t, err)
		assert.NotNil(t, locale)
		assert.Equal(t, ".", locale.ThouSep)
		assert.Equal(t, ",", locale.DecSep)
		assert.Equal(t, "VND", locale.Name)
		assert.Equal(t, "₫", locale.ComSymbol)
		assert.Equal(t, true, locale.Pre)
		assert.Equal(t, 2, locale.FractionLength)
		assert.Equal(t, "", locale.SpaceSep)
	})

	t.Run("Nagetive", func(t *testing.T) {
		var lang = "afafd"
		locale, err := number.GetLocale(lang)
		assert.Error(t, err)
		assert.Nil(t, locale)
	})

	t.Run("Parse-Minnus", func(t *testing.T) {
		var lang = "vi"
		locale, err := number.GetLocale(lang)
		assert.Nil(t, err)
		assert.NotNil(t, locale)

		var str = number.ParseByLocale (locale, "-(45.000,50)VND")
		assert.Equal(t, "-45000.50", str)
	})

	t.Run("Parse-Positive", func(t *testing.T) {
		var lang = "vi"
		locale, err := number.GetLocale(lang)
		assert.Nil(t, err)
		assert.NotNil(t, locale)

		var str = number.ParseByLocale (locale, "đ45.000,50")
		assert.Equal(t, "45000.50", str)
	})

	t.Run("ParseFromLangue", func(t *testing.T) {
		var lang = "vi"
		var str =  "đ45.000,50"
		val, err := number.Parse(lang, str)
		assert.Nil(t, err)
		assert.Equal(t, "45000.50", val)
	})

	t.Run("ParseAndRevert", func(t *testing.T) {
		var lang = "vi"
		var str =  "₫ 45.000,50"
		val, err := number.Parse(lang, str)

		assert.Nil(t, err)
		assert.Equal(t, "45000.50", val)
	})


	t.Run("FormatNumber", func(t *testing.T) {
		var lang = "vi"
		var str =  45000.50
		val, err := number.Format(lang, str)

		assert.Nil(t, err)
		assert.Equal(t, "₫45.000,50", val)
	})
}