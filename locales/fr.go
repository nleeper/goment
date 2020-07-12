package locales

import (
	"strconv"
	"strings"
)

// FrLocale is the French language locale.
var FrLocale = NewLocale(
	"fr",
	strings.Split("dimanche_lundi_mardi_mercredi_jeudi_vendredi_samedi", "_"),
	strings.Split("dim._lun._mar._mer._jeu._ven._sam.", "_"),
	strings.Split("di_lu_ma_me_je_ve_sa", "_"),
	strings.Split("janvier_février_mars_avril_mai_juin_juillet_août_septembre_octobre_novembre_décembre", "_"),
	strings.Split("janv._févr._mars_avr._mai_juin_juil._août_sept._oct._nov._déc.", "_"),
	func(num int, period string) string {
		suffix := ""
		switch period {
		case "D":
			if num == 1 {
				suffix = "er"
			}
			return strconv.Itoa(num) + suffix
		case "M", "Q", "DDD", "d":
			suffix = "e"
			if num == 1 {
				suffix = "er"
			}
		case "w", "W":
			suffix = "e"
			if num == 1 {
				suffix = "re"
			}
		default:
			suffix = "e"
			if num == 1 {
				suffix = "er"
			}
		}
		return strconv.Itoa(num) + suffix
	},
	nil,
	1,
	longDateFormats{
		"LTS":  "HH:mm:ss",
		"LT":   "HH:mm",
		"L":    "DD/MM/YYYY",
		"LL":   "D MMMM YYYY",
		"LLL":  "D MMMM YYYY HH:mm",
		"LLLL": "dddd D MMMM YYYY HH:mm",
	},
)
