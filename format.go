package goment

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nleeper/goment/regexps"
)

type formatReplacement func(*Goment) string

type formatPadding struct {
	Token        string
	TargetLength int
	ForceSign    bool
}

var formatReplacementFuncs = map[string]formatReplacement{}

// Format takes a string of tokens and replaces them with their corresponding values to display the Goment.
func (g *Goment) Format(args ...interface{}) string {
	format := ""

	numArgs := len(args)
	if numArgs < 1 {
		format = "YYYY-MM-DDTHH:mm:ssZ"
	} else {
		format = args[0].(string)
	}

	loadReplacementFunctions()

	return convertFormat(g, format)
}

func loadReplacementFunctions() {
	if len(formatReplacementFuncs) > 0 {
		return
	}

	addReplacementFunction("M", padding("MM", 2), "Mo", func(g *Goment) string {
		return strconv.Itoa(g.Month())
	})
	addReplacementFunction("MMM", emptyPadding(), "", func(g *Goment) string {
		return g.locale.MonthsShort[g.Month()-1]
	})
	addReplacementFunction("MMMM", emptyPadding(), "", func(g *Goment) string {
		return g.locale.Months[g.Month()-1]
	})
	addReplacementFunction("D", padding("DD", 2), "Do", func(g *Goment) string {
		return strconv.Itoa(g.Date())
	})
	addReplacementFunction("Y", emptyPadding(), "", func(g *Goment) string {
		y := g.Year()
		if y <= 9999 {
			return zeroFill(y, 4, false)
		}
		return "+" + strconv.Itoa(y)
	})
	addReplacementFunction("", padding("YY", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Year() % 100)
	})
	addReplacementFunction("", padding("YYYY", 4), "", func(g *Goment) string {
		return strconv.Itoa(g.Year())
	})
	addReplacementFunction("", padding("YYYYY", 5), "", func(g *Goment) string {
		return strconv.Itoa(g.Year())
	})
	addReplacementFunction("", padding("YYYYYY", 6, true), "", func(g *Goment) string {
		return strconv.Itoa(g.Year())
	})
	addReplacementFunction("d", emptyPadding(), "do", func(g *Goment) string {
		return strconv.Itoa(g.Day())
	})
	addReplacementFunction("dd", emptyPadding(), "", func(g *Goment) string {
		return g.locale.WeekdaysMin[g.Day()]
	})
	addReplacementFunction("ddd", emptyPadding(), "", func(g *Goment) string {
		return g.locale.WeekdaysShort[g.Day()]
	})
	addReplacementFunction("dddd", emptyPadding(), "", func(g *Goment) string {
		return g.locale.Weekdays[g.Day()]
	})
	addReplacementFunction("DDD", padding("DDDD", 3), "DDDo", func(g *Goment) string {
		return strconv.Itoa(g.DayOfYear())
	})
	addReplacementFunction("e", emptyPadding(), "", func(g *Goment) string {
		return strconv.Itoa(g.Weekday())
	})
	addReplacementFunction("E", emptyPadding(), "", func(g *Goment) string {
		return strconv.Itoa(g.ISOWeekday())
	})
	addReplacementFunction("w", padding("ww", 2), "wo", func(g *Goment) string {
		return strconv.Itoa(g.Week())
	})
	addReplacementFunction("W", padding("WW", 2), "Wo", func(g *Goment) string {
		return strconv.Itoa(g.ISOWeek())
	})
	addReplacementFunction("Q", emptyPadding(), "Qo", func(g *Goment) string {
		return strconv.Itoa(g.Quarter())
	})
	addReplacementFunction("H", padding("HH", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Hour())
	})
	addReplacementFunction("h", padding("hh", 2), "", func(g *Goment) string {
		val := 0
		mod := g.Hour() % 12
		if mod == 0 {
			val = 12
		} else {
			val = mod
		}
		return strconv.Itoa(val)
	})
	addReplacementFunction("k", padding("kk", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Hour() + 1)
	})
	addReplacementFunction("a", emptyPadding(), "", func(g *Goment) string {
		return g.locale.MeridiemFunc(g.Hour(), g.Minute(), true)
	})
	addReplacementFunction("A", emptyPadding(), "", func(g *Goment) string {
		return g.locale.MeridiemFunc(g.Hour(), g.Minute(), false)
	})
	addReplacementFunction("m", padding("mm", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Minute())
	})
	addReplacementFunction("s", padding("ss", 2), "", func(g *Goment) string {
		return strconv.Itoa(g.Second())
	})
	addReplacementFunction("X", emptyPadding(), "", func(g *Goment) string {
		return fmt.Sprintf("%d", g.ToUnix())
	})
	addReplacementFunction("x", emptyPadding(), "", func(g *Goment) string {
		return fmt.Sprintf("%d", g.ToTime().UnixNano()/int64(time.Millisecond))
	})
	addReplacementFunction("Z", emptyPadding(), "", func(g *Goment) string {
		return offset(g, ":")
	})
	addReplacementFunction("ZZ", emptyPadding(), "", func(g *Goment) string {
		return offset(g, "")
	})
	addReplacementFunction("z", emptyPadding(), "", func(g *Goment) string {
		return timezoneName(g)
	})
	addReplacementFunction("zz", emptyPadding(), "", func(g *Goment) string {
		return timezoneName(g)
	})
}

func padding(token string, length int, forceSign ...bool) formatPadding {
	sign := false
	if len(forceSign) > 0 {
		sign = forceSign[0]
	}

	return formatPadding{
		Token:        token,
		TargetLength: length,
		ForceSign:    sign,
	}
}

func emptyPadding() formatPadding {
	return formatPadding{}
}

func addReplacementFunction(token string, padding formatPadding, ordinal string, cb formatReplacement) {
	if token != "" {
		formatReplacementFuncs[token] = cb
	}
	if padding.Token != "" {
		formatReplacementFuncs[padding.Token] = func(g *Goment) string {
			var val = cb(g)
			i, _ := strconv.Atoi(val)

			return zeroFill(i, padding.TargetLength, padding.ForceSign)
		}
	}
	if ordinal != "" {
		formatReplacementFuncs[ordinal] = func(g *Goment) string {
			var val = cb(g)
			i, _ := strconv.Atoi(val)
			return g.locale.OrdinalFunc(i, token)
		}
	}
}

func convertFormat(g *Goment, layout string) string {
	// Replace any Goment locale specific format tokens (LTS, L, LL, etc).
	layout = replaceFormatTokens(
		layout,
		regexps.LocaleRegex.FindAllStringIndex(layout, -1),
		func(text string) (string, bool) {
			return g.locale.LongDateFormat(text)
		},
	)

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
			match, ok := formatReplacementFuncs[text]
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

func timezoneName(g *Goment) string {
	tz, _ := g.ToTime().Zone()
	return tz
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
