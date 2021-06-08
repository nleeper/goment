package goment

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type compareTest struct {
	Message     string `json:"message"`
	Date        string `json:"date"`
	LaterDate   string `json:"later_date,omitempty"`
	Inclusivity string `json:"inclusivity,omitempty"`
}

type compareTestCase struct {
	Unit       string        `json:"unit"`
	TrueCases  []compareTest `json:"true"`
	FalseCases []compareTest `json:"false"`
}

type compareTestScenario struct {
	Date  string            `json:"date"`
	Cases []compareTestCase `json:"cases"`
}

var compareTestFileMap = map[string]map[string]compareTestScenario{}

func getTestScenario(fileName, scenarioName string) compareTestScenario {
	if _, ok := compareTestFileMap[fileName]; !ok {
		var testMap map[string]compareTestScenario
		raw, _ := ioutil.ReadFile(filepath.Join("testdata", fileName+".json"))
		json.Unmarshal(raw, &testMap)
		compareTestFileMap[fileName] = testMap
	}

	return compareTestFileMap[fileName][scenarioName]
}

func runTestsForScenario(t *testing.T, scenario compareTestScenario, f func(...interface{}) bool) {
	for _, c := range scenario.Cases {
		for _, tt := range c.TrueCases {
			assert.True(t, f(tt.Date, c.Unit), tt.Message)
		}
		for _, ft := range c.FalseCases {
			assert.False(t, f(ft.Date, c.Unit), ft.Message)
		}
	}
}

func runTestsForBetweenScenario(t *testing.T, scenario compareTestScenario, f func(...interface{}) bool) {
	for _, c := range scenario.Cases {
		for _, tt := range c.TrueCases {
			assert.True(t, f(tt.Date, tt.LaterDate, c.Unit, tt.Inclusivity), tt.Message)
		}
		for _, ft := range c.FalseCases {
			assert.False(t, f(ft.Date, ft.LaterDate, c.Unit, ft.Inclusivity), ft.Message)
		}
	}
}

func TestIsBeforeUsesNowIfNoArguments(t *testing.T) {
	testTime := time.Date(2018, 12, 15, 17, 8, 0, 0, chicagoLocation())
	timeNow = func() time.Time {
		return testTime
	}

	lib := simpleString("2010-10-11")
	assert.True(t, lib.IsBefore())

	// Reset timeNow.
	timeNow = time.Now
}

func TestIsBeforeReturnsFalseInvalidFirstArgument(t *testing.T) {
	lib := simpleString("2010-10-11")
	assert.False(t, lib.IsBefore(1))
}

func TestIsBeforeReturnsFalseIfInvalidSecondArgument(t *testing.T) {
	lib := simpleString("2010-10-11")
	assert.False(t, lib.IsBefore("2012-01-01", 1))
}

func TestIsBeforeNoUnits(t *testing.T) {
	scenario := getTestScenario("is_before", "no_units")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsBefore)
	assert.False(t, lib.IsBefore(lib), "goments are not before themselves")
}

func TestIsBeforeYear(t *testing.T) {
	scenario := getTestScenario("is_before", "year")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsBefore)
	assert.False(t, lib.IsBefore(lib, "year"), "same goments are not before the same year")
}

func TestIsBeforeMonth(t *testing.T) {
	scenario := getTestScenario("is_before", "month")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsBefore)
	assert.False(t, lib.IsBefore(lib, "month"), "same goments are not before the same month")
}

func TestIsBeforeDay(t *testing.T) {
	scenario := getTestScenario("is_before", "day")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsBefore)
	assert.False(t, lib.IsBefore(lib, "day"), "same goments are not before the same day")
}

func TestIsBeforeHour(t *testing.T) {
	scenario := getTestScenario("is_before", "hour")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsBefore)
	assert.False(t, lib.IsBefore(lib, "hour"), "same goments are not before the same hour")
}

func TestIsBeforeMinute(t *testing.T) {
	scenario := getTestScenario("is_before", "minute")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsBefore)
	assert.False(t, lib.IsBefore(lib, "minute"), "same goments are not before the same minute")
}

func TestIsBeforeSecond(t *testing.T) {
	scenario := getTestScenario("is_before", "second")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsBefore)
	assert.False(t, lib.IsBefore(lib, "second"), "same goments are not before the same second")
}

func TestIsSameReturnsFalseIfNoArguments(t *testing.T) {
	testTime := time.Date(2018, 12, 15, 17, 8, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	lib := simpleTime(testTime)
	assert.False(t, lib.IsSame())

	// Reset timeNow.
	timeNow = time.Now
}

func TestIsSameReturnsFalseInvalidFirstArgument(t *testing.T) {
	lib := simpleString("2010-10-11")
	assert.False(t, lib.IsSame(1))
}

func TestIsSameReturnsFalseIfInvalidSecondArgument(t *testing.T) {
	lib := simpleString("2010-10-11")
	assert.False(t, lib.IsSame("2012-01-01", 1))
}

func TestIsSameNoUnits(t *testing.T) {
	scenario := getTestScenario("is_same", "no_units")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSame)
	assert.True(t, lib.IsSame(lib), "goments are the same as themselves")
}

func TestIsSameDeref(t *testing.T) {
	date := "2020-05-10"

	lib := simpleString(date)
	lib2 := *lib
	lib3 := *lib

	assert.True(t, lib2.IsSame(lib3), "deref goments are the same as themselves")
}

func TestIsSameYear(t *testing.T) {
	scenario := getTestScenario("is_same", "year")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSame)
	assert.True(t, lib.IsSame(lib, "year"), "same goments are in the same year")
}

func TestIsSameMonth(t *testing.T) {
	scenario := getTestScenario("is_same", "month")

	lib := simpleTime(time.Date(2011, 3, 3, 4, 5, 6, 7, time.UTC))

	runTestsForScenario(t, scenario, lib.IsSame)
	assert.True(t, lib.IsSame(lib, "month"), "same goments are in the same month")
}

func TestIsSameDay(t *testing.T) {
	scenario := getTestScenario("is_same", "day")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSame)
	assert.True(t, lib.IsSame(lib, "day"), "same goments are in the same day")
}

func TestIsSameHour(t *testing.T) {
	scenario := getTestScenario("is_same", "hour")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSame)
	assert.True(t, lib.IsSame(lib, "hour"), "same goments are in the same hour")
}

func TestIsSameMinute(t *testing.T) {
	scenario := getTestScenario("is_same", "minute")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSame)
	assert.True(t, lib.IsSame(lib, "minute"), "same goments are in the same minute")
}

func TestIsSameSecond(t *testing.T) {
	scenario := getTestScenario("is_same", "second")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSame)
	assert.True(t, lib.IsSame(lib, "second"), "same goments are in the same second")
}

func TestIsAfterUsesNowIfNoArguments(t *testing.T) {
	testTime := time.Date(2018, 12, 15, 17, 8, 0, 0, chicagoLocation())
	timeNow = func() time.Time {
		return testTime
	}

	lib := simpleString("2020-10-11")
	assert.True(t, lib.IsAfter())

	// Reset timeNow.
	timeNow = time.Now
}

func TestIsAfterReturnsFalseInvalidFirstArgument(t *testing.T) {
	lib := simpleString("2010-10-11")
	assert.False(t, lib.IsAfter(1))
}

func TestIsAfterReturnsFalseIfInvalidSecondArgument(t *testing.T) {
	lib := simpleString("2010-10-11")
	assert.False(t, lib.IsAfter("2012-01-01", 1))
}

func TestIsAfterNoUnits(t *testing.T) {
	scenario := getTestScenario("is_after", "no_units")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsAfter)
	assert.False(t, lib.IsAfter(lib), "goments are not after themselves")
}

func TestIsAfterYear(t *testing.T) {
	scenario := getTestScenario("is_after", "year")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsAfter)
	assert.False(t, lib.IsAfter(lib, "year"), "same goments are not after the same year")
}

func TestIsAfterMonth(t *testing.T) {
	scenario := getTestScenario("is_after", "month")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsAfter)
	assert.False(t, lib.IsAfter(lib, "month"), "same goments are not after the same month")
}

func TestIsAfterDay(t *testing.T) {
	scenario := getTestScenario("is_after", "day")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsAfter)
	assert.False(t, lib.IsAfter(lib, "day"), false, "same goments are not after the same day")
}

func TestIsAfterHour(t *testing.T) {
	scenario := getTestScenario("is_after", "hour")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsAfter)
	assert.False(t, lib.IsAfter(lib, "hour"), "same goments are not after the same hour")
}

func TestIsAfterMinute(t *testing.T) {
	scenario := getTestScenario("is_after", "minute")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsAfter)
	assert.False(t, lib.IsAfter(lib, "minute"), "same goments are not after the same minute")
}

func TestIsAfterSecond(t *testing.T) {
	scenario := getTestScenario("is_after", "second")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsAfter)
	assert.False(t, lib.IsAfter(lib, "second"), "same goments are not after the same second")
}

func TestIsSameOrBeforeNoUnits(t *testing.T) {
	scenario := getTestScenario("is_same_or_before", "no_units")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrBefore)
	assert.True(t, lib.IsSameOrBefore(lib), "goments are the same as themselves")
}

func TestIsSameOrBeforeDeref(t *testing.T) {
	date := "2020-05-10"

	lib := simpleString(date)
	lib2 := *lib
	lib3 := *lib

	assert.True(t, lib2.IsSameOrBefore(lib3), "deref goments are the same as themselves")
}

func TestIsSameOrBeforeYear(t *testing.T) {
	scenario := getTestScenario("is_same_or_before", "year")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrBefore)
	assert.True(t, lib.IsSameOrBefore(lib, "year"), "same goments are in the same year")
}

func TestIsSameOrBeforeMonth(t *testing.T) {
	scenario := getTestScenario("is_same_or_before", "month")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrBefore)
	assert.True(t, lib.IsSameOrBefore(lib, "month"), "same goments are in the same month")
}

func TestIsSameOrBeforeDay(t *testing.T) {
	scenario := getTestScenario("is_same_or_before", "day")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrBefore)
	assert.True(t, lib.IsSameOrBefore(lib, "day"), "same goments are in the same day")
}

func TestIsSameOrBeforeHour(t *testing.T) {
	scenario := getTestScenario("is_same_or_before", "hour")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrBefore)
	assert.True(t, lib.IsSameOrBefore(lib, "hour"), "same goments are in the same hour")
}

func TestIsSameOrBeforeMinute(t *testing.T) {
	scenario := getTestScenario("is_same_or_before", "minute")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrBefore)
	assert.True(t, lib.IsSameOrBefore(lib, "minute"), "same goments are in the same minute")
}

func TestIsSameOrBeforeSecond(t *testing.T) {
	scenario := getTestScenario("is_same_or_before", "second")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrBefore)
	assert.True(t, lib.IsSameOrBefore(lib, "second"), "same goments are in the same second")
}

func TestIsSameOrAfterNoUnits(t *testing.T) {
	scenario := getTestScenario("is_same_or_after", "no_units")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrAfter)
	assert.True(t, lib.IsSameOrAfter(lib), "goments are the same as themselves")
}

func TestIsSameOrAfterDeref(t *testing.T) {
	date := "2020-05-10"

	lib := simpleString(date)
	lib2 := *lib
	lib3 := *lib

	assert.True(t, lib2.IsSameOrAfter(lib3), "deref goments are the same as themselves")
}

func TestIsSameOrAfterYear(t *testing.T) {
	scenario := getTestScenario("is_same_or_after", "year")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrAfter)
	assert.True(t, lib.IsSameOrAfter(lib, "year"), "same goments are in the same year")
}

func TestIsSameOrAfterMonth(t *testing.T) {
	scenario := getTestScenario("is_same_or_after", "month")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrAfter)
	assert.True(t, lib.IsSameOrAfter(lib, "month"), "same goments are in the same month")
}

func TestIsSameOrAfterDay(t *testing.T) {
	scenario := getTestScenario("is_same_or_after", "day")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrAfter)
	assert.True(t, lib.IsSameOrAfter(lib, "day"), "same goments are in the same day")
}

func TestIsSameOrAfterHour(t *testing.T) {
	scenario := getTestScenario("is_same_or_after", "hour")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrAfter)
	assert.True(t, lib.IsSameOrAfter(lib, "hour"), "same goments are in the same hour")
}

func TestIsSameOrAfterMinute(t *testing.T) {
	scenario := getTestScenario("is_same_or_after", "minute")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrAfter)
	assert.True(t, lib.IsSameOrAfter(lib, "minute"), "same goments are in the same minute")
}

func TestIsSameOrAfterSecond(t *testing.T) {
	scenario := getTestScenario("is_same_or_after", "second")

	lib := simpleString(scenario.Date)

	runTestsForScenario(t, scenario, lib.IsSameOrAfter)
	assert.True(t, lib.IsSameOrAfter(lib, "second"), "same goments are in the same second")
}

func TestIsBetweenReturnsFalseNotEnoughArguments(t *testing.T) {
	lib := simpleString("2010-10-11")
	assert.False(t, lib.IsBetween(1))
}

func TestIsBetweenReturnsFalseInvalidFirstArgument(t *testing.T) {
	lib := simpleString("2010-10-11")
	assert.False(t, lib.IsBetween(1, "2010-11-11"))
}

func TestIsBetweenReturnsFalseInvalidSecondArgument(t *testing.T) {
	lib := simpleString("2010-10-11")
	assert.False(t, lib.IsBetween("2009-10-11", 2))
}

func TestIsBetweenNoUnits(t *testing.T) {
	scenario := getTestScenario("is_between", "no_units")

	lib := simpleString(scenario.Date)

	runTestsForBetweenScenario(t, scenario, lib.IsBetween)
	assert.False(t, lib.IsBetween(lib, lib), "goments are not between themselves")
}

func TestIsBetweenYear(t *testing.T) {
	scenario := getTestScenario("is_between", "year")

	lib := simpleString(scenario.Date)

	runTestsForBetweenScenario(t, scenario, lib.IsBetween)
	assert.False(t, lib.IsBetween(lib, lib, "year"), "same goments are not between the same year")
}

func TestIsBetweenMonth(t *testing.T) {
	scenario := getTestScenario("is_between", "month")

	lib := simpleString(scenario.Date)

	runTestsForBetweenScenario(t, scenario, lib.IsBetween)
	assert.False(t, lib.IsBetween(lib, lib, "month"), "same goments are not between the same month")
}

func TestIsBetweenDay(t *testing.T) {
	scenario := getTestScenario("is_between", "day")

	lib := simpleString(scenario.Date)

	runTestsForBetweenScenario(t, scenario, lib.IsBetween)
	assert.False(t, lib.IsBetween(lib, lib, "day"), "same goments are not between the same day")
}

func TestIsBetweenHour(t *testing.T) {
	scenario := getTestScenario("is_between", "hour")

	lib := simpleString(scenario.Date)

	runTestsForBetweenScenario(t, scenario, lib.IsBetween)
	assert.False(t, lib.IsBetween(lib, lib, "hour"), "same goments are not between the same hour")
}

func TestIsBetweenMinute(t *testing.T) {
	scenario := getTestScenario("is_between", "minute")

	lib := simpleString(scenario.Date)

	runTestsForBetweenScenario(t, scenario, lib.IsBetween)
	assert.False(t, lib.IsBetween(lib, lib, "minute"), "same goments are not between the same minute")
}

func TestIsBetweenSecond(t *testing.T) {
	scenario := getTestScenario("is_between", "second")

	lib := simpleString(scenario.Date)

	runTestsForBetweenScenario(t, scenario, lib.IsBetween)
	assert.False(t, lib.IsBetween(lib, lib, "second"), "same goments are not between the same second")
}

func TestIsBetweenNoUnitsInclusivity(t *testing.T) {
	scenario := getTestScenario("is_between", "no_units_inclusivity")

	lib := simpleString(scenario.Date)

	runTestsForBetweenScenario(t, scenario, lib.IsBetween)
}

func TestIsBetweenSecondInclusivity(t *testing.T) {
	scenario := getTestScenario("is_between", "second_inclusivity")

	lib := simpleString(scenario.Date)

	runTestsForBetweenScenario(t, scenario, lib.IsBetween)
}
