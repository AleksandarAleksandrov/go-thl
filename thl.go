// Package thl is a date and time helper library written for Go.
// Inspired by date-fns library for JavaScript.
package thl

import (
	"errors"
	"math"
	"sort"
	"time"
)

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

// Sorts the given slice of dates depending on the comparator value
func Sort(timeSlice []time.Time, comparator Constant) {
	tm := timeSort(timeSlice)
	if ASC == comparator {
		sort.Sort(tm)
	} else if DESC == comparator {
		sort.Sort(sort.Reverse(tm))
	}
}

// Sorts a slice of dates chronologically
func SortAsc(timeSlice []time.Time) {
	Sort(timeSlice, ASC)
}

// Sorts a slice of date reverse chronologically
func SortDesc(timeSlice []time.Time) {
	Sort(timeSlice, DESC)
}

// Compares two dates and returns:
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

	return -1 * DaysDifference(endRangeStartDate, initialRangeEndDate), nil
}

// Cheks if the passed date is within the range
func IsWithinRange(date, startDate, endDate time.Time) bool {
	return date.After(startDate) && date.Before(endDate)
}

// milisecond helpers

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

// seconds helper

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

// minute helpers

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

// hour helpers
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

//days helper

func LastDayOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 0, 0, 0, 0, t.Location())
}

func FirstDayOfNextYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
}

func DaysDifference(endDate, startDate time.Time) (days int) {
	cur := startDate
	for cur.Year() < endDate.Year() {
		// add 1 to count the last day of the year too.
		days += LastDayOfYear(cur).YearDay() - cur.YearDay() + 1
		cur = FirstDayOfNextYear(cur)
	}
	days += endDate.YearDay() - cur.YearDay()
	return days
}
