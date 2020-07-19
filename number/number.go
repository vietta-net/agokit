package number

import (
	"github.com/leekchan/accounting"
	"github.com/pkg/errors"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
)

//Set locate from a language code
func GetLocale(lang string) (*accounting.Locale, error ){
	var  t = language.Make(lang)
	if t.String() != lang {
		return nil, errors.Errorf("Language code `%s` is not support.", lang)
	}
	var val =  1234.5
	p := message.NewPrinter(t)
	s := p.Sprintf("%v", val)

	arr := strings.Split(s, "")

	unit, _ := currency.FromTag(t)

	format := currency.Symbol(unit.Amount(float64(0)))
	formatStr := p.Sprint(format)
	sym := strings.Split(formatStr, "")

	var lc = &accounting.Locale{
		Name: unit.String(),
		FractionLength: 2,
		ThouSep: arr[1],
		DecSep:  arr[5],
		SpaceSep: "",
		ComSymbol: sym[0],
		Pre: true,
	}

	return lc, nil
}