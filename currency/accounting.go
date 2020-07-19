package currency

import (
	"fmt"
	"github.com/leekchan/accounting"
)

//Set Account with locate
func GetAccounting(lc *accounting.Locale) *accounting.Accounting {
	var ac = &accounting.Accounting{
		Symbol:         lc.ComSymbol,
		Precision:      lc.FractionLength,
		Thousand:       lc.ThouSep,
		Decimal:        lc.DecSep,
		Format:         "%s%v",
		FormatNegative: "%s-%v",
		FormatZero:     "%s0." + fmt.Sprintf("%0*d", lc.FractionLength, 0),
	}
	if ! lc.Pre{
		ac.Format = "%v%s"
		ac.FormatNegative = "-%v%s"
		ac.FormatZero = fmt.Sprintf("0.%0*d", lc.FractionLength, 0) + "%s"
	}

	return ac
}

//Format currency based on Locale
func FormatByLocale(lc *accounting.Locale, number interface{}) string {
	ac := GetAccounting(lc)
	return FormatByAccounting(ac, number)
}

//Convert currency to string based on Locale
func Format(lang string, number interface{}) ( string, error ) {
	locale, err := GetLocale(lang)
	if err != nil {
		return "", err
	}
	return FormatByLocale(locale, number), nil
}

//Convert to  accounting format
func FormatByAccounting(ac *accounting.Accounting, number interface{}) string {
	return ac.FormatMoney(number)
}

//Convert accounting String like đ45.000,50 to 45000.50
func ParseByLocale(lc *accounting.Locale, number string) string {
	return accounting.UnformatNumber(number, lc.FractionLength, lc.Name)
}

func Parse(lang string, number string) (string , error){
	locale, err := GetLocale(lang)
	if err != nil {
		return "", err
	}
	return ParseByLocale(locale, number), nil
}
