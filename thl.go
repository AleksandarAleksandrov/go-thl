// Package thl is a date and time helper library written for Go.
// Inspired by date-fns library for JavaScript.
package thl

import (
	"errors"
	"math"
	"sort"
	"time"
)

/***********************
 *** General Helpers ***
 ***********************/

// Internal structure used for sorting slices of dates
type timeSort []time.Time

func (ts timeSort) Len() int {
	return len(ts)
}

func (ts timeSort) Less(i, j int) bool {
	return ts[i].Before(ts[j])
}

func (ts timeSort) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// Constant is the type of constant used in the package
type Constant int

// ASC is used for sorting dates chronologically
const ASC Constant = 1

// DESC is used to sort dates reverse chronologically
const DESC Constant = -1

// Sort the given slice of dates depending on the comparator value
func Sort(timeSlice []time.Time, comparator Constant) {
	tm := timeSort(timeSlice)
	if ASC == comparator {
		sort.Sort(tm)
	} else if DESC == comparator {
		sort.Sort(sort.Reverse(tm))
	}
}

// SortAsc sorts a slice of dates chronologically
func SortAsc(timeSlice []time.Time) {
	Sort(timeSlice, ASC)
}

// SortDesc sorts a slice of date reverse chronologically
func SortDesc(timeSlice []time.Time) {
	Sort(timeSlice, DESC)
}

// Compare two dates and return:
// -1 if the first date is before the second date
// 0 if they are the same date
// 1 is the second date is before the first date
func Compare(first time.Time, second time.Time) int {
	if first.Before(second) {
		return -1
	} else if first.After(second) {
		return 1
	} else {
		return 0
	}
}

// Finds index of the date from the slice that is closest to the date passed
func ClosestIndexTo(dateToCompare time.Time, datesSlice []time.Time) (int, error) {

	if datesSlice == nil {
		return 0, errors.New("Passed slice of dates was nil")
	}

	if len(datesSlice) == 0 {
		return 0, errors.New("Passed slice of dates was of size 0")
	}

	var closestIndex int
	var currentMinMili int64 = math.MaxInt64
	var currentMinNano int64 = math.MaxInt64

	dateMiliUnix := dateToCompare.Unix()
	dateNanoUnix := dateToCompare.UnixNano()
	for index, date := range datesSlice {
		unixDiffMili := dateMiliUnix - date.Unix()
		unixDiffNano := dateNanoUnix - date.UnixNano()

		// get the positive values
		if unixDiffMili < 0 {
			unixDiffMili = -1 * unixDiffMili
		}

		if unixDiffNano < 0 {
			unixDiffNano = -1 * unixDiffNano
		}

		if unixDiffMili < currentMinMili {
			currentMinMili = unixDiffMili
			currentMinNano = unixDiffNano
			closestIndex = index
		} else if unixDiffMili == currentMinMili {
			if unixDiffNano < currentMinNano {
				currentMinMili = unixDiffMili
				currentMinNano = unixDiffNano
				closestIndex = index
			}
		}
	}
	return closestIndex, nil
}

// Finds the date from the slice that is closest to the date passed
func ClosestTo(dateToCompare time.Time, datesSlice []time.Time) (time.Time, error) {
	index, err := ClosestIndexTo(dateToCompare, datesSlice)

	if err != nil {
		return time.Time{}, err
	}

	return datesSlice[index], nil
}

// Checks if the date is in the future
func IsFuture(dateToTest time.Time) bool {
	return dateToTest.After(time.Now())
}

// Checks if the date is in the past
func IsPast(dateToTest time.Time) bool {
	return dateToTest.Before(time.Now())
}

// Finds the latest date chronologically
func Max(datesSlice []time.Time) (time.Time, error) {
	if datesSlice == nil {
		return time.Time{}, errors.New("Passed slice of dates was nil")
	}

	if len(datesSlice) == 0 {
		return time.Time{}, errors.New("Passed slice of dates was of size 0")
	}

	dateToReturn := datesSlice[0]

	for _, testDate := range datesSlice {
		if testDate.After(dateToReturn) {
			dateToReturn = testDate
		}
	}

	return dateToReturn, nil
}

// Finds the latest date reverse chronologically
func Min(datesSlice []time.Time) (time.Time, error) {
	if datesSlice == nil {
		return time.Time{}, errors.New("Passed slice of dates was nil")
	}

	if len(datesSlice) == 0 {
		return time.Time{}, errors.New("Passed slice of dates was of size 0")
	}

	dateToReturn := datesSlice[0]

	for _, testDate := range datesSlice {
		if testDate.Before(dateToReturn) {
			dateToReturn = testDate
		}
	}

	return dateToReturn, nil
}

// Cheks if the ranges overlap
func AreRangesOverlapping(
	initialRangeStartDate,
	initialRangeEndDate,
	endRangeStartDate,
	endRangeEndDate time.Time) bool {
	return initialRangeStartDate.Before(initialRangeEndDate) &&
		endRangeStartDate.Before(initialRangeEndDate) &&
		endRangeStartDate.Before(endRangeEndDate)
}

// Gets the number of days that the ranges overlap
func GetOverlappingDaysInRanges(
	initialRangeStartDate,
	initialRangeEndDate,
	endRangeStartDate,
	endRangeEndDate time.Time) (int, error) {
	areOverlapping := AreRangesOverlapping(initialRangeStartDate, initialRangeEndDate, endRangeStartDate, endRangeEndDate)

	if !areOverlapping {
		return 0, errors.New("Ranges do not overlap")
	}

	return -1 * DifferenceInDays(endRangeStartDate, initialRangeEndDate), nil
}

// Cheks if the passed date is within the range
func IsWithinRange(date, startDate, endDate time.Time) bool {
	return date.After(startDate) && date.Before(endDate)
}

/****************************
 *** Millisecond Helpers ***
 ****************************/

func AddMilliseconds(date time.Time, amount int) time.Time {
	return date.Add(time.Millisecond * time.Duration(amount))
}

func DifferenceInMilliseconds(dateLeft, dateRight time.Time) int64 {
	leftInMill := dateLeft.UnixNano() / int64(time.Millisecond)
	rightInMil := dateRight.UnixNano() / int64(time.Millisecond)
	return leftInMill - rightInMil
}

func GetMilliseconds(date time.Time) int {
	return date.Nanosecond() / int(time.Millisecond)
}

func SetMillisecond(date time.Time, amount int) (time.Time, error) {
	if amount < 0 || amount > 999 {
		return date, errors.New("Passed amount was less than 0 or more than 999. Date left unchanged.")
	}
	return time.Date(date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second(),
		int(time.Millisecond)*amount,
		date.Location()), nil
}

/**********************
 *** Second Helpers ***
 **********************/

func AddSeconds(date time.Time, amount int) time.Time {
	return date.Add(time.Second * time.Duration(amount))
}

func DifferenceInSeconds(dateLeft, dateRight time.Time) float64 {
	return dateLeft.Sub(dateRight).Seconds()
}

func EndOfSecond(date time.Time) time.Time {
	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second(),
		999999999,
		date.Location())
}

func IsSameSecond(dateLeft, dateRight time.Time) bool {
	return dateLeft.Year() == dateRight.Year() &&
		dateLeft.Month() == dateRight.Month() &&
		dateLeft.Day() == dateRight.Day() &&
		dateLeft.Hour() == dateRight.Hour() &&
		dateLeft.Minute() == dateRight.Minute() &&
		dateLeft.Second() == dateRight.Second()
}

func IsThisSecond(date time.Time) bool {
	now := time.Now()
	return IsSameSecond(date, now)
}

func SetSeconds(date time.Time, seconds int) (time.Time, error) {
	if seconds < 0 || seconds > 59 {
		return date, errors.New("Passed amount was less than 0 or more than 59. Date left unchanged.")
	}
	return time.Date(date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		seconds,
		date.Nanosecond(),
		date.Location()), nil
}

func StartOfSecond(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), 0, date.Location())
}

/**********************
 *** Minute Helpers ***
 **********************/

func AddMinutes(date time.Time, amount int) time.Time {
	return date.Add(time.Minute * time.Duration(amount))
}

func DifferenceInMinutes(dateLeft, dateRight time.Time) float64 {
	return dateLeft.Sub(dateRight).Minutes()
}

func EndOfMinute(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 59, 999999999, date.Location())
}

func IsSameMinute(dateLeft, dateRight time.Time) bool {
	return dateLeft.Year() == dateRight.Year() &&
		dateLeft.Month() == dateRight.Month() &&
		dateLeft.Day() == dateRight.Day() &&
		dateLeft.Hour() == dateRight.Hour() &&
		dateLeft.Minute() == dateRight.Minute()
}

func IsThisMinute(date time.Time) bool {
	now := time.Now()
	return IsSameMinute(date, now)
}

func SetMinutes(date time.Time, minutes int) (time.Time, error) {

	if minutes < 0 || minutes > 59 {
		return date, errors.New("Passed amount was less than 0 or more than 59. Date left unchanged.")
	}

	return time.Date(date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		minutes,
		date.Second(),
		date.Nanosecond(),
		date.Location()), nil
}

func StartOfMinute(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0, date.Location())
}

/********************
 *** Hour Helpers ***
 ********************/

func AddHours(date time.Time, amount int) time.Time {
	return date.Add(time.Hour * time.Duration(amount))
}

func DifferenceInHours(dateLeft, dateRight time.Time) float64 {
	return dateLeft.Sub(dateRight).Hours()
}

func EndOfHour(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), 59, 59, 999999999, date.Location())
}

func IsSameHour(dateLeft, dateRight time.Time) bool {
	return dateLeft.Year() == dateRight.Year() &&
		dateLeft.Month() == dateRight.Month() &&
		dateLeft.Day() == dateRight.Day() &&
		dateLeft.Hour() == dateRight.Hour()
}

func IsThisHour(date time.Time) bool {
	return IsSameHour(time.Now(), date)
}

func SetHours(date time.Time, hours int) (time.Time, error) {
	if hours < 0 || hours > 23 {
		return date, errors.New("Passed amount was less than 0 or more than 23. Date left unchanged.")
	}

	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		hours,
		date.Minute(),
		date.Second(),
		date.Nanosecond(),
		date.Location()), nil
}

func StartOfHour(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), 0, 0, 0, date.Location())
}

/*******************
 *** Day Helpers ***
 *******************/

func LastDayOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 0, 0, 0, 0, t.Location())
}

func FirstDayOfNextYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
}

func FirstDayOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

func DifferenceInDays(endDate, startDate time.Time) (days int) {
	cur := startDate
	for cur.Year() < endDate.Year() {
		// add 1 to count the last day of the year too.
		days += LastDayOfYear(cur).YearDay() - cur.YearDay() + 1
		cur = FirstDayOfNextYear(cur)
	}
	days += endDate.YearDay() - cur.YearDay()
	return days
}

func AddDays(date time.Time, amount int) time.Time {
	return date.Add(time.Hour * 24 * time.Duration(amount))
}

func EndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())
}

func StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func EachDay(startDate, endDate time.Time) ([]time.Time, error) {
	var datesRange []time.Time

	if endDate.Before(startDate) {
		return datesRange, errors.New("End date can not be before start date. Returned empty slice.")
	}

	counterDate := AddDays(startDate, 1)

	for counterDate.Before(endDate) {
		datesRange = append(datesRange, counterDate)
		counterDate = StartOfDay(AddDays(counterDate, 1))
	}

	return datesRange, nil
}

func IsSameDay(firstDay, secondDay time.Time) bool {
	return firstDay.Year() == secondDay.Year() &&
		firstDay.Month() == secondDay.Month() &&
		firstDay.Day() == secondDay.Day()
}

func EndOfToday() time.Time {
	return EndOfDay(time.Now())
}

func EndOfTomorrow() time.Time {
	return EndOfDay(AddDays(time.Now(), 1))
}

func EndOfYesterday() time.Time {
	return EndOfDay(AddDays(time.Now(), -1))
}

func StartOfToday() time.Time {
	return StartOfDay(time.Now())
}

func StartOfTomorrow() time.Time {
	return StartOfDay(AddDays(time.Now(), 1))
}

func StartOfYesterday() time.Time {
	return StartOfDay(AddDays(time.Now(), -1))
}

func IsToday(date time.Time) bool {
	return IsSameDay(date, time.Now())
}

func IsTomorrow(date time.Time) bool {
	return IsSameDay(date, AddDays(time.Now(), 1))
}

func IsYesterday(date time.Time) bool {
	return IsSameDay(date, AddDays(time.Now(), -1))
}

func SetDayOfYear(date time.Time, dayNumber int) (time.Time, error) {
	daysInYear := 365
	if IsLeapYear(date.Year()) {
		daysInYear++
	}

	if dayNumber < 0 || dayNumber > daysInYear {
		return date, errors.New("Given day number if out of range. Returned unchanged date.")
	}

	return AddDays(FirstDayOfYear(date), dayNumber), nil
}

func SetDayOfMonth(date time.Time, dayMonthNumber int) (time.Time, error) {
	daysInMonth := GetDaysInMonth(date)
	if dayMonthNumber > daysInMonth {
		return time.Time{}, errors.New("Passed days count is bigger than the days count in the month of the passed date")
	}
	return time.Date(date.Year(),
		date.Month(),
		dayMonthNumber,
		date.Hour(),
		date.Minute(),
		date.Second(),
		date.Nanosecond(),
		date.Location()), nil
}

/***********************
 *** Weekday Helpers ***
 ***********************/

func IsWeekend(date time.Time) bool {
	return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday
}

func IsMonToFri(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Monday ||
		weekday == time.Tuesday ||
		weekday == time.Wednesday ||
		weekday == time.Thursday ||
		weekday == time.Friday
}

/********************
 *** Week Helpers ***
 ********************/

func EndOfWeek(date time.Time) time.Time {
	weekday := date.Weekday()
	// 0 == Sunday
	if weekday == 0 {
		return EndOfDay(date)
	}
	weekDiff := 7 - weekday
	return EndOfDay(AddDays(date, int(weekDiff)))
}

func StartOfWeek(date time.Time) time.Time {
	return StartOfDay(AddDays(EndOfWeek(date), -7))
}

func IsSameWeek(dateOne, dateTwo time.Time) bool {
	weekOne := EndOfWeek(dateOne)
	weekTwo := EndOfWeek(dateTwo)
	return IsSameDay(weekOne, weekTwo)
}

func IsThisWeek(date time.Time) bool {
	return IsSameWeek(date, time.Now())
}

func AddWeeks(date time.Time, amount int) time.Time {
	return AddDays(date, 7*amount)
}

func DifferenceInWeeks(endDate, startDate time.Time) int {
	return DifferenceInDays(endDate, startDate) / 7
}

/*********************
 *** Month Helpers ***
 *********************/

func IsSameMonth(dateOne, dateTwo time.Time) bool {
	return dateOne.Year() == dateTwo.Year() && dateOne.Month() == dateTwo.Month()
}

func AddMonths(date time.Time, amount int) time.Time {
	years := amount / 12
	leftOverMonths := amount % 12
	dateToReturn := time.Date(
		date.Year()+years,
		date.Month()+time.Month(leftOverMonths),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second(),
		date.Nanosecond(),
		date.Location())
	return dateToReturn
}

func GetDaysInMonth(date time.Time) int {
	if IsLeapYear(date.Year()) && time.February == date.Month() {
		return 29
	}

	if time.February == date.Month() {
		return 28
	}

	month := date.Month()
	if time.January == month ||
		time.March == month ||
		time.May == month ||
		time.July == month ||
		time.August == month ||
		time.October == month ||
		time.December == month {
		return 31
	}

	return 30
}

func EndOfMonth(date time.Time) time.Time {
	monthDays := GetDaysInMonth(date)
	dateToReturn, _ := SetDayOfMonth(date, monthDays)
	return time.Date(dateToReturn.Year(),
		dateToReturn.Month(),
		dateToReturn.Day(),
		23, 59, 59, 999999999,
		dateToReturn.Location())
}

func IsFirstDayOfMonth(date time.Time) bool {
	return date.Day() == 1
}

func IsLastDayOfMonth(date time.Time) bool {
	daysIsMonth := GetDaysInMonth(date)
	return daysIsMonth == date.Day()
}

func IsThisMonth(date time.Time) bool {
	return IsSameMonth(date, time.Now())
}

func StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

/***********************
 *** Quarter Helpers ***
 ***********************/

func AddQuarters(date time.Time, amount int) time.Time {
	return AddMonths(date, amount*3)
}

func IsFirstQuarter(date time.Time) bool {
	return time.January == date.Month() || time.February == date.Month() || time.March == date.Month()
}

func IsSecondQuarter(date time.Time) bool {
	return time.April == date.Month() || time.May == date.Month() || time.June == date.Month()
}

func IsThirdQuarter(date time.Time) bool {
	return time.July == date.Month() || time.August == date.Month() || time.September == date.Month()
}

func IsFourthQuarter(date time.Time) bool {
	return time.October == date.Month() || time.November == date.Month() || time.December == date.Month()
}

func EndOfQuarter(date time.Time) time.Time {
	if IsFirstQuarter(date) {
		return time.Date(date.Year(), time.March, 31, 23, 59, 59, 999999999, date.Location())
	}

	if IsSecondQuarter(date) {
		return time.Date(date.Year(), time.June, 31, 23, 59, 59, 999999999, date.Location())
	}

	if IsFirstQuarter(date) {
		return time.Date(date.Year(), time.September, 31, 23, 59, 59, 999999999, date.Location())
	}

	return time.Date(date.Year(), time.December, 31, 23, 59, 59, 999999999, date.Location())
}

func StartOfQuarter(date time.Time) time.Time {
	if IsFirstQuarter(date) {
		return time.Date(date.Year(), time.January, 1, 0, 0, 0, 0, date.Location())
	}

	if IsSecondQuarter(date) {
		return time.Date(date.Year(), time.April, 1, 0, 0, 0, 0, date.Location())
	}

	if IsFirstQuarter(date) {
		return time.Date(date.Year(), time.July, 1, 0, 0, 0, 0, date.Location())
	}

	return time.Date(date.Year(), time.October, 1, 0, 0, 0, 0, date.Location())
}

func GetQuarter(date time.Time) int {
	if IsFirstQuarter(date) {
		return 1
	}
	if IsSecondQuarter(date) {
		return 2
	}
	if IsFirstQuarter(date) {
		return 3
	}
	return 4
}

func IsSameQuarter(dateOne, dateTwo time.Time) bool {
	if dateOne.Year() != dateTwo.Year() {
		return false
	}
	return EndOfQuarter(dateOne) == EndOfQuarter(dateTwo)
}

func IsThisQuarter(date time.Time) bool {
	return IsSameQuarter(date, time.Now())
}

/********************
 *** Year Helpers ***
 ********************/

func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else {
		return true
	}
}

func AddYears(date time.Time, amount int) time.Time {
	return time.Date(date.Year()+amount,
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second(),
		date.Nanosecond(),
		date.Location())
}

func SetYear(date time.Time, year int) time.Time {
	return time.Date(year,
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second(),
		date.Nanosecond(),
		date.Location())
}

func EndOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), time.December,
		31, 23, 59, 59, 999999999, date.Location())
}

func StartOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), time.January, 1, 0, 0, 0, 0, date.Location())
}

func IsSameYear(dateOne, dateTwo time.Time) bool {
	return dateOne.Year() == dateTwo.Year()
}

func IsThisYear(date time.Time) bool {
	return IsSameYear(date, time.Now())
}
