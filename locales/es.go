package locales

import (
	"fmt"
	"strings"
)

// EsLocale is the Spanish language locale.
var EsLocale = NewLocale(
	"es",
	strings.Split("domingo_lunes_martes_miércoles_jueves_viernes_sábado", "_"),
	strings.Split("dom._lun._mar._mié._jue._vie._sáb.", "_"),
	strings.Split("do_lu_ma_mi_ju_vi_sá", "_"),
	strings.Split("enero_febrero_marzo_abril_mayo_junio_julio_agosto_septiembre_octubre_noviembre_diciembre", "_"),
	strings.Split("ene_feb_mar_abr_may_jun_jul_ago_sep_oct_nov_dic", "_"),
	func(num int, period string) string {
		return fmt.Sprintf("%dº", num)
	},
	nil,
	1,
	longDateFormats{
		"LTS":  "H:mm:ss",
		"LT":   "H:mm",
		"L":    "DD/MM/YYYY",
		"LL":   "D [de] MMMM [de] YYYY",
		"LLL":  "D [de] MMMM [de] YYYY H:mm",
		"LLLL": "dddd, D [de] MMMM [de] YYYY H:mm",
	},
)
