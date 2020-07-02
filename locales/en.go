package locales

import (
	"strconv"
	"strings"
)

// EnLocale is the US English language locale.
var EnLocale = NewLocale(
	"en",
	strings.Split("Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday", "_"),
	strings.Split("Sun_Mon_Tue_Wed_Thu_Fri_Sat", "_"),
	strings.Split("Su_Mo_Tu_We_Th_Fr_Sa", "_"),
	strings.Split("January_February_March_April_May_June_July_August_September_October_November_December", "_"),
	strings.Split("Jan_Feb_Mar_Apr_May_Jun_Jul_Aug_Sep_Oct_Nov_Dec", "_"),
	func(num int, period string) string {
		suffix := "th"
		switch num % 10 {
		case 1:
			if num%100 != 11 {
				suffix = "st"
			}
		case 2:
			if num%100 != 12 {
				suffix = "nd"
			}
		case 3:
			if num%100 != 13 {
				suffix = "rd"
			}
		}
		return strconv.Itoa(num) + suffix
	},
	nil,
	0,
	longDateFormats{
		"LTS":  "h:mm:ss A",
		"LT":   "h:mm A",
		"L":    "MM/DD/YYYY",
		"LL":   "MMMM D, YYYY",
		"LLL":  "MMMM D, YYYY h:mm A",
		"LLLL": "dddd, MMMM D, YYYY h:mm A",
	},
)
