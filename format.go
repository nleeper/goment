package goment

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nleeper/goment/locales"
	"github.com/nleeper/goment/regexps"
	"github.com/tkuchiki/go-timezone"
)

type formatReplacementFunc func(*Goment) string

type formatPadding struct {
	token        string
	targetLength int
	forceSign    bool
}

var formatReplacements = map[string]formatReplacementFunc{}

// Format takes a string of tokens and replaces them with their corresponding values to display the Goment.
func (g *Goment) Format(args ...interface{}) string {
	format := ""

	numArgs := len(args)
	if numArgs < 1 {
		format = "YYYY-MM-DDTHH:mm:ssZ"
	} else {
		format = args[0].(string)
	}

	return convertFormat(g, format)
}

func loadFormatReplacements() {
	if len(formatReplacements) > 0 {
		return
	}

	addFormatReplacement("M", padding("MM", 2), "Mo", func(g *Goment) string {
		return strconv.Itoa(g.Month())
	})
	addFormatReplacement("MMM", emptyPadding(), "", func(g *Goment) string {
		return g.locale.MonthsShort[g.Month()-1]
	})
	addFormatReplacement("MMMM", emptyPadding(), "", func(g *Goment) string {
		return g.locale.Months[g.Month()-1]
	})

	addFormatReplacement("D", padding("DD", 2), "Do", func(g *Goment) string {
		return strconv.Itoa(g.Date())
	})
	addFormatReplacement("DDD", padding("DDDD", 3), "DDDo", func(g *Goment) string {
		return strconv.Itoa(g.DayOfYear())
	})

	addFormatReplacement("Y", emptyPadding(), "", func(g *Goment) string {
		y := g.Year()
		if y <= 9999 {
			return zeroFill(y, 4, false)
		}
		return "+" + strconv.Itoa(y)
	})
	addFormatReplacement("", padding("YY", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Year() % 100)
	})
	addFormatReplacement("", padding("YYYY", 4), "", func(g *Goment) string {
		return strconv.Itoa(g.Year())
	})
	addFormatReplacement("", padding("YYYYY", 5), "", func(g *Goment) string {
		return strconv.Itoa(g.Year())
	})
	addFormatReplacement("", padding("YYYYYY", 6, true), "", func(g *Goment) string {
		return strconv.Itoa(g.Year())
	})

	addFormatReplacement("d", emptyPadding(), "do", func(g *Goment) string {
		return strconv.Itoa(g.Day())
	})
	addFormatReplacement("dd", emptyPadding(), "", func(g *Goment) string {
		return g.locale.WeekdaysMin[g.Day()]
	})
	addFormatReplacement("ddd", emptyPadding(), "", func(g *Goment) string {
		return g.locale.WeekdaysShort[g.Day()]
	})
	addFormatReplacement("dddd", emptyPadding(), "", func(g *Goment) string {
		return g.locale.Weekdays[g.Day()]
	})

	addFormatReplacement("e", emptyPadding(), "", func(g *Goment) string {
		return strconv.Itoa(g.Weekday())
	})
	addFormatReplacement("E", emptyPadding(), "", func(g *Goment) string {
		return strconv.Itoa(g.ISOWeekday())
	})

	addFormatReplacement("w", padding("ww", 2), "wo", func(g *Goment) string {
		return strconv.Itoa(g.Week())
	})
	addFormatReplacement("W", padding("WW", 2), "Wo", func(g *Goment) string {
		return strconv.Itoa(g.ISOWeek())
	})

	addFormatReplacement("", padding("gg", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.WeekYear() % 100)
	})
	addFormatReplacement("", padding("gggg", 4), "", func(g *Goment) string {
		return strconv.Itoa(g.WeekYear())
	})
	addFormatReplacement("", padding("ggggg", 5), "", func(g *Goment) string {
		return strconv.Itoa(g.WeekYear())
	})
	addFormatReplacement("", padding("GG", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.ISOWeekYear() % 100)
	})
	addFormatReplacement("", padding("GGGG", 4), "", func(g *Goment) string {
		return strconv.Itoa(g.ISOWeekYear())
	})
	addFormatReplacement("", padding("GGGGG", 5), "", func(g *Goment) string {
		return strconv.Itoa(g.ISOWeekYear())
	})

	addFormatReplacement("Q", emptyPadding(), "Qo", func(g *Goment) string {
		return strconv.Itoa(g.Quarter())
	})

	addFormatReplacement("H", padding("HH", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Hour())
	})

	addFormatReplacement("h", padding("hh", 2), "", func(g *Goment) string {
		val := 0
		mod := g.Hour() % 12
		if mod == 0 {
			val = 12
		} else {
			val = mod
		}
		return strconv.Itoa(val)
	})

	addFormatReplacement("k", padding("kk", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Hour() + 1)
	})

	addFormatReplacement("a", emptyPadding(), "", func(g *Goment) string {
		return g.locale.MeridiemFunc(g.Hour(), g.Minute(), true)
	})
	addFormatReplacement("A", emptyPadding(), "", func(g *Goment) string {
		return g.locale.MeridiemFunc(g.Hour(), g.Minute(), false)
	})

	addFormatReplacement("m", padding("mm", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Minute())
	})

	addFormatReplacement("s", padding("ss", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Second())
	})

	addFormatReplacement("X", emptyPadding(), "", func(g *Goment) string {
		return fmt.Sprintf("%d", g.ToUnix())
	})
	addFormatReplacement("x", emptyPadding(), "", func(g *Goment) string {
		return fmt.Sprintf("%d", g.ToTime().UnixNano()/int64(time.Millisecond))
	})

	addFormatReplacement("Z", emptyPadding(), "", func(g *Goment) string {
		return offset(g, ":")
	})
	addFormatReplacement("ZZ", emptyPadding(), "", func(g *Goment) string {
		return offset(g, "")
	})
	addFormatReplacement("z", emptyPadding(), "", func(g *Goment) string {
		return timezoneAbbr(g)
	})
	addFormatReplacement("zz", emptyPadding(), "", func(g *Goment) string {
		return timezoneAbbr(g)
	})
	addFormatReplacement("zzzz", emptyPadding(), "", func(g *Goment) string {
		return timezoneFullName(g)
	})
}

func addFormatReplacement(token string, padding formatPadding, ordinal string, f formatReplacementFunc) {
	if token != "" {
		formatReplacements[token] = f
	}

	if padding.token != "" {
		formatReplacements[padding.token] = func(g *Goment) string {
			var val = f(g)
			i, _ := strconv.Atoi(val)

			return zeroFill(i, padding.targetLength, padding.forceSign)
		}
	}

	if ordinal != "" {
		formatReplacements[ordinal] = func(g *Goment) string {
			var val = f(g)
			i, _ := strconv.Atoi(val)
			return g.locale.OrdinalFunc(i, token)
		}
	}
}

func padding(token string, length int, forceSign ...bool) formatPadding {
	sign := false
	if len(forceSign) > 0 {
		sign = forceSign[0]
	}

	return formatPadding{
		token:        token,
		targetLength: length,
		forceSign:    sign,
	}
}

func emptyPadding() formatPadding {
	return formatPadding{}
}

func expandLocaleFormats(layout string, locale locales.LocaleDetails) string {
	return replaceFormatTokens(
		layout,
		regexps.LocaleRegex.FindAllStringIndex(layout, -1),
		func(text string) (string, bool) {
			return locale.LongDateFormat(text)
		},
	)
}

func convertFormat(g *Goment, layout string) string {
	// Replace any Goment locale specific format tokens (LTS, L, LL, etc).
	layout = expandLocaleFormats(layout, g.locale)

	// Replace any bracketed text in layout.
	bracketMatch := regexps.BracketRegex.FindAllString(layout, -1)
	bracketsFound := len(bracketMatch) > 0

	// Replace bracketed text with token like $1.
	if bracketsFound {
		for i := range bracketMatch {
			layout = strings.Replace(layout, bracketMatch[i], makeBracketToken(i), 1)
		}
	}

	// Replace any Goment format tokens that are not standard to Go formatting (DDD, Mo, etc).
	layout = replaceFormatTokens(
		layout,
		regexps.TokenRegex.FindAllStringIndex(layout, -1),
		func(text string) (string, bool) {
			match, ok := formatReplacements[text]
			if ok {
				return match(g), ok
			}
			return "", false
		},
	)

	// Replace back any bracketed text.
	if bracketsFound {
		for i := range bracketMatch {
			layout = strings.Replace(layout, makeBracketToken(i), bracketMatch[i][1:len(bracketMatch[i])-1], 1)
		}
	}

	return layout
}

func replaceFormatTokens(layout string, matches [][]int, replacementFunc func(string) (string, bool)) string {
	for i := range matches {
		start, end := matches[i][0], matches[i][1]
		matchText := layout[start:end]

		replaceText, ok := replacementFunc(matchText)
		if !ok {
			replaceText = matchText
		}

		diff := len(replaceText) - len(matchText)
		layout = layout[0:start] + replaceText + layout[end:len(layout)]

		// If the replacement text is longer/shorter than the match, shift the remaining indexes.
		if diff != 0 {
			for j := i + 1; j < len(matches); j++ {
				matches[j][0] += diff
				matches[j][1] += diff
			}
		}
	}

	return layout
}

func offset(g *Goment, sep string) string {
	os := g.UTCOffset()
	sign := "+"

	if os < 0 {
		os = os * -1
		sign = "-"
	}

	return sign + zeroFill(os/60, 2, false) + sep + zeroFill(os%60, 2, false)
}

func timezoneAbbr(g *Goment) string {
	tz, _ := g.ToTime().Zone()
	return tz
}

func timezoneFullName(g *Goment) string {
	tz := timezone.New()
	tzAbbrInfos, _ := tz.GetTzAbbreviationInfo(timezoneAbbr(g))
	return tzAbbrInfos[0].Name()
}

func zeroFill(val int, length int, forceSign bool) string {
	absNumber := abs(val)

	sign := val >= 0
	signValue := ""
	if sign {
		if forceSign {
			signValue = "+"
		}
	} else {
		signValue = "-"
	}

	return signValue + fmt.Sprintf("%0"+strconv.Itoa(length)+"d", absNumber)
}

func makeBracketToken(num int) string {
	return fmt.Sprintf("$%v", num+1)
}
