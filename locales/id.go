package locales

import (
	"fmt"
	"strings"
)

// IdLocale is the Indonesian language locale.
var IdLocale = newLocale(
	"id",
	strings.Split("Minggu_Senin_Selasa_Rabu_Kamis_Jumat_Sabtu", "_"),
	strings.Split("Min_Sen_Sel_Rab_Kam_Jum_Sab", "_"),
	strings.Split("Mg_Sn_Sl_Rb_Km_Jm_Sb", "_"),
	strings.Split("Januari_Februari_Maret_April_Mei_Juni_Juli_Agustus_September_Oktober_November_Desember", "_"),
	strings.Split("Jan_Feb_Mar_Apr_Mei_Jun_Jul_Agt_Sep_Okt_Nov_Des", "_"),
	func(num int, period string) string {
		return fmt.Sprintf("%d", num)
	},
	func(hours int, minutes int, isLower bool) string {
		lowerOrTitle := func(word string) string {
			if isLower {
				return strings.ToLower(word)
			}

			return strings.Title(word)
		}

		switch {
		case hours < 5:
			return lowerOrTitle("dini hari")
		case hours < 11:
			return lowerOrTitle("pagi")
		case hours < 15:
			return lowerOrTitle("siang")
		case hours < 19:
			return lowerOrTitle("sore")
		default:
			return lowerOrTitle("malam")
		}
	},
	week{Dow: 0, Doy: 6},
	longDateFormats{
		"LTS":  "HH:mm:ss",
		"LT":   "HH:mm",
		"L":    "DD/MM/YYYY",
		"LL":   "D MMMM YYYY",
		"LLL":  "D MMMM YYYY, HH:mm",
		"LLLL": "dddd, D MMMM YYYY, HH:mm",
	},
	relativeTimeFormats{
		"future": "dalam %s",
		"past":   "%s yang lalu",
		"s":      "beberapa detik",
		"ss":     "%d detik",
		"m":      "semenit",
		"mm":     "%d menit",
		"h":      "sejam",
		"hh":     "%d jam",
		"d":      "sehari",
		"dd":     "%d hari",
		"M":      "sebulan",
		"MM":     "%d bulan",
		"y":      "setahun",
		"yy":     "%d tahun",
	},
	calendarFunctions{
		"sameDay": func(hours int, day int) string {
			return "[hari ini pukul] LT"
		},
		"nextDay": func(hours int, day int) string {
			return "[besok pukul] LT"
		},
		"nextWeek": func(hours int, day int) string {
			return "dddd [pukul] LT"
		},
		"lastDay": func(hours int, day int) string {
			return "[kemarin pukul] LT"
		},
		"lastWeek": func(hours int, day int) string {
			return "dddd [lalu pukul] LT"
		},
		"sameElse": func(hours int, day int) string {
			return "L"
		},
	},
	`(?i)(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)`,
	`(?i)(Jan|Feb|Mar|Apr|Mei|Jun|Jul|Agt|Sep|Okt|Nov|Des)`,
	`(?i)(Minggu|Senin|Selasa|Rabu|Kamis|Jumat|Sabtu)`,
	`(?i)(Min|Sen|Sel|Rab|Kam|Jum|Sab)`,
	`(?i)(Mg|Sn|Sl|Rb|Km|Jm|Sb)`,
	`\d{1,2}`,
)
