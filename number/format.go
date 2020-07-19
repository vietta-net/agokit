package number

import (
	"github.com/vietta-net/agokit/currency"
	"github.com/leekchan/accounting"

)

func Format(lang string, number interface{}) ( string, error ) {
	lc,  err := currency.GetLocale(lang)
	if err != nil {
		return "",  err
	}
	return FormatByLocale(lc, number), nil
}


func FormatByLocale(lc *accounting.Locale, number interface{}) (string){
	return accounting.FormatNumber(number, lc.FractionLength, lc.ThouSep, lc.DecSep)
}