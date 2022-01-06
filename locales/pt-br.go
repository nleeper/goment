package locales

import (
	"fmt"
	"strings"
)

func getPtBRCalendarPronoun(hours int) string {
	if hours != 1 {
		return "às"
	}
	return "à"
}

// PtBRLocale is the Brazilian Portuguese language locale.
var PtBRLocale = newLocale(
	"pt-br",
	strings.Split("domingo_segunda-feira_terça-feira_quarta-feira_quinta-feira_sexta-feira_sábado", "_"),
	strings.Split("dom_seg_ter_qua_qui_sex_sab", "_"),
	strings.Split("do_se_te_qa_qi_se_sa", "_"),
	strings.Split("janeiro_fevereiro_março_abril_maio_junho_julho_agosto_setembro_outubro_novembro_dezembro", "_"),
	strings.Split("jan_fev_mar_abr_mai_jun_jul_ago_set_out_nov_dez", "_"),
	func(num int, period string) string {
		return fmt.Sprintf("%dº", num)
	},
	nil,
	week{Dow: 1, Doy: 4},
	longDateFormats{
		"LTS":  "H:mm:ss",
		"LT":   "H:mm",
		"L":    "DD/MM/YYYY",
		"LL":   "D [de] MMMM [de] YYYY",
		"LLL":  "D [de] MMMM [de] YYYY H:mm",
		"LLLL": "dddd, D [de] MMMM [de] YYYY H:mm",
	},
	relativeTimeFormats{
		"future": "em %s",
		"past":   "há %s",
		"s":      "alguns segundos",
		"ss":     "%d segundos",
		"m":      "um minuto",
		"mm":     "%d minutos",
		"h":      "uma hora",
		"hh":     "%d horas",
		"d":      "um dia",
		"dd":     "%d dias",
		"M":      "um mês",
		"MM":     "%d meses",
		"y":      "um ano",
		"yy":     "%d anos",
	},
	calendarFunctions{
		"sameDay": func(hours int, day int) string {
			return "[hoje " + getPtBRCalendarPronoun(hours) + "] LT"
		},
		"nextDay": func(hours int, day int) string {
			return "[amanhã " + getPtBRCalendarPronoun(hours) + "] LT"
		},
		"nextWeek": func(hours int, day int) string {
			return "dddd [" + getPtBRCalendarPronoun(hours) + "] LT"
		},
		"lastDay": func(hours int, day int) string {
			return "[ontem " + getPtBRCalendarPronoun(hours) + "] LT"
		},
		"lastWeek": func(hours int, day int) string {
			return "[na] dddd [passada " + getPtBRCalendarPronoun(hours) + "] LT"
		},
		"sameElse": func(hours int, day int) string {
			return "L"
		},
	},
	`(?i)(janeiro|fevereiro|março|abril|maio|junho|julho|agosto|setembro|outubro|novembro|dezembro)`,
	`(?i)(jan\.?|fev\.?|mar\.?|abr\.?|mai\.?|jun\.?|jul\.?|ago\.?|set\.?|out\.?|nov\.?|dez\.?)`,
	`(?i)(domingo|segunda-feira|terça-feira|quarta-feira|quinta-feira|sexta-feira|sábado)`,
	`(?i)(dom\.?|seg\.?|ter\.?|qua\.?|qui\.?|sex\.?|sáb\.?)`,
	`(?i)(do|lu|ma|mi|ju|vi|sá)`,
	`\d{1,2}º`,
)
