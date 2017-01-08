package thl

import (
	"fmt"
	"time"
)

var (
	first  = time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	second = time.Date(2016, 6, 6, 6, 6, 6, 6, time.UTC)
	third  = time.Date(2016, 6, 6, 6, 6, 6, 7, time.UTC)
	fourth = time.Date(2015, 1, 1, 1, 0, 0, 0, time.UTC)

	futureDate = time.Date(3001, 1, 1, 0, 0, 0, 0, time.UTC)
	pastDate   = time.Date(1001, 1, 1, 0, 0, 0, 0, time.UTC)
)

func ExampleSort() {
	timeSlice := []time.Time{fourth, third, second, first}
	Sort(timeSlice, ASC)
	for _, val := range timeSlice {
		fmt.Println(val)
	}
	fmt.Println("------------------")
	Sort(timeSlice, DESC)
	for _, val := range timeSlice {
		fmt.Println(val)
	}

	// Output:
	// 2015-01-01 01:00:00 +0000 UTC
	// 2016-06-06 06:06:06.000000006 +0000 UTC
	// 2016-06-06 06:06:06.000000007 +0000 UTC
	// 2017-01-01 00:00:00 +0000 UTC
	// ------------------
	// 2017-01-01 00:00:00 +0000 UTC
	// 2016-06-06 06:06:06.000000007 +0000 UTC
	// 2016-06-06 06:06:06.000000006 +0000 UTC
	// 2015-01-01 01:00:00 +0000 UTC
}

func ExampleSortAsc() {
	timeSlice := []time.Time{fourth, third, second, first}
	SortAsc(timeSlice)
	for _, val := range timeSlice {
		fmt.Println(val)
	}

	// Output:
	// 2015-01-01 01:00:00 +0000 UTC
	// 2016-06-06 06:06:06.000000006 +0000 UTC
	// 2016-06-06 06:06:06.000000007 +0000 UTC
	// 2017-01-01 00:00:00 +0000 UTC
}

func ExampleSortDesc() {
	timeSlice := []time.Time{fourth, third, second, first}
	SortDesc(timeSlice)
	for _, val := range timeSlice {
		fmt.Println(val)
	}

	// Output:
	// 2017-01-01 00:00:00 +0000 UTC
	// 2016-06-06 06:06:06.000000007 +0000 UTC
	// 2016-06-06 06:06:06.000000006 +0000 UTC
	// 2015-01-01 01:00:00 +0000 UTC
}

func ExampleCompare() {
	fmt.Println(Compare(first, second))
	fmt.Println(Compare(fourth, third))
	fmt.Println(Compare(second, second))
	// Output:
	// 1
	// -1
	// 0
}

func ExampleClosestIndexTo() {
	timeSlice := []time.Time{fourth, second.Add(-1 * time.Nanosecond), third.Add(time.Nanosecond), first}
	fmt.Println(ClosestIndexTo(third, timeSlice))
	fmt.Println(ClosestIndexTo(third, nil))
	fmt.Println(ClosestIndexTo(third, []time.Time{}))
	//Output:
	// 2 <nil>
	// 0 Passed slice of dates was nil
	// 0 Passed slice of dates was of size 0
}

func ExampleClosestTo() {
	timeSlice := []time.Time{fourth, second.Add(-1 * time.Nanosecond), third.Add(time.Nanosecond), first}
	fmt.Println(ClosestTo(third, timeSlice))
	fmt.Println(ClosestTo(third, nil))
	fmt.Println(ClosestTo(third, []time.Time{}))
	// Output:
	// 2016-06-06 06:06:06.000000008 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC Passed slice of dates was nil
	// 0001-01-01 00:00:00 +0000 UTC Passed slice of dates was of size 0
}

func ExampleIsFuture() {
	fmt.Println(IsFuture(pastDate))
	fmt.Println(IsFuture(futureDate))
	// Output:
	// false
	// true
}

func ExampleIsPast() {
	fmt.Println(IsPast(pastDate))
	fmt.Println(IsPast(futureDate))
	// Output:
	// true
	// false
}

func ExampleMax() {
	timeSlice := []time.Time{fourth, third, second, first}
	fmt.Println(Max(timeSlice))
	fmt.Println(Max(nil))
	fmt.Println(Max([]time.Time{}))
	// Output:
	// 2017-01-01 00:00:00 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC Passed slice of dates was nil
	// 0001-01-01 00:00:00 +0000 UTC Passed slice of dates was of size 0
}

func ExampleMin() {
	timeSlice := []time.Time{fourth, third, second, first}
	fmt.Println(Min(timeSlice))
	fmt.Println(Min(nil))
	fmt.Println(Min([]time.Time{}))
	// Output:
	// 2015-01-01 01:00:00 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC Passed slice of dates was nil
	// 0001-01-01 00:00:00 +0000 UTC Passed slice of dates was of size 0
}

func ExampleAreRangesOverlapping() {
	fmt.Println(
		AreRangesOverlapping(time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 1, 18, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 1, 16, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 2, 1, 0, 0, 0, 0, time.UTC),
		))
	fmt.Println(
		AreRangesOverlapping(time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 1, 15, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 1, 16, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 2, 1, 0, 0, 0, 0, time.UTC),
		))
	// Output:
	// true
	// false
}

func ExampleGetOverlappingDaysInRanges() {
	fmt.Println(
		GetOverlappingDaysInRanges(time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 1, 18, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 1, 16, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 2, 1, 0, 0, 0, 0, time.UTC),
		))
	fmt.Println(
		GetOverlappingDaysInRanges(time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 1, 15, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 1, 16, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 2, 1, 0, 0, 0, 0, time.UTC),
		))
	// Output:
	// 2 <nil>
	// 0 Ranges do not overlap
}

func ExampleIsWithinRange() {
	fmt.Println(IsWithinRange(
		time.Date(2017, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 1, 3, 0, 0, 0, 0, time.UTC),
	))
	fmt.Println(IsWithinRange(
		time.Date(2017, 1, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 1, 3, 0, 0, 0, 0, time.UTC),
	))
	// Output:
	// true
	// false
}

func ExampleAddMilliseconds() {
	fmt.Println(futureDate)
	fmt.Println(AddMilliseconds(futureDate, 300))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:00:00.3 +0000 UTC
}

func ExampleDifferenceInMilliseconds() {
	fmt.Println(DifferenceInMilliseconds(futureDate, AddMilliseconds(futureDate, 300)))
	// Output:
	// -300
}

func ExampleGetMilliseconds() {
	fmt.Println(GetMilliseconds(AddMilliseconds(futureDate, 300)))
	// Output:
	// 300
}

func ExampleSetMillisecond() {
	fmt.Println(futureDate)
	fmt.Println(SetMillisecond(futureDate, 999))
	fmt.Println(SetMillisecond(futureDate, 1000))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:00:00.999 +0000 UTC <nil>
	// 3001-01-01 00:00:00 +0000 UTC Passed amount was less than 0 or more than 999. Date left unchanged.
}

func ExampleAddSeconds() {
	fmt.Println(futureDate)
	fmt.Println(AddSeconds(futureDate, 30))
	//Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:00:30 +0000 UTC
}

func ExampleDifferenceInSeconds() {
	fmt.Println(futureDate)
	fmt.Println(AddSeconds(futureDate, 30))
	fmt.Println(DifferenceInSeconds(AddSeconds(futureDate, 30), futureDate))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:00:30 +0000 UTC
	// 30
}

func ExampleEndOfSecond() {
	fmt.Println(futureDate)
	fmt.Println(EndOfSecond(futureDate))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:00:00.999999999 +0000 UTC
}

func ExampleIsSameSecond() {
	dateOne := time.Date(2017, 1, 1, 1, 1, 1, 0, time.UTC)
	dateTwo := time.Date(2017, 1, 1, 1, 1, 1, 999999999, time.UTC)
	fmt.Println(IsSameSecond(dateOne, dateTwo))
	fmt.Println(IsSameSecond(dateOne, AddSeconds(dateTwo, 1)))
	// Output:
	// true
	// false
}

func ExampleIsThisSecond() {
	fmt.Println(IsThisSecond(time.Now()))
	fmt.Println(IsThisSecond(AddSeconds(time.Now(), 1)))
	// Output:
	// true
	// false
}

func ExampleSetSeconds() {
	fmt.Println(SetSeconds(futureDate, 33))
	fmt.Println(SetSeconds(futureDate, 60))
	// Output:
	// 3001-01-01 00:00:33 +0000 UTC <nil>
	// 3001-01-01 00:00:00 +0000 UTC Passed amount was less than 0 or more than 59. Date left unchanged.
}

func ExampleStartOfSecond() {
	fmt.Println(third)
	fmt.Println(StartOfSecond(third))
	// Output:
	// 2016-06-06 06:06:06.000000007 +0000 UTC
	// 2016-06-06 06:06:06 +0000 UTC
}

func ExampleAddMinutes() {
	fmt.Println(futureDate)
	fmt.Println(AddMinutes(futureDate, 3))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:03:00 +0000 UTC
}

func ExampleDifferenceInMinutes() {
	fmt.Println(futureDate)
	fmt.Println(AddMinutes(futureDate, 59))
	fmt.Println(DifferenceInMinutes(AddMinutes(futureDate, 59), futureDate))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:59:00 +0000 UTC
	// 59
}

func ExampleEndOfMinute() {
	fmt.Println(third)
	fmt.Println(EndOfMinute(third))
	// Output:
	// 2016-06-06 06:06:06.000000007 +0000 UTC
	// 2016-06-06 06:06:59.999999999 +0000 UTC
}

func ExampleIsSameMinute() {
	val, _ := SetSeconds(futureDate, 59)
	fmt.Println(IsSameMinute(futureDate, val))
	fmt.Println(IsSameMinute(futureDate, AddMinutes(futureDate, 1)))
	// Output:
	// true
	// false
}

func ExampleIsThisMinute() {
	fmt.Println(IsThisMinute(time.Now()))
	fmt.Println(IsThisMinute(AddMinutes(time.Now(), 1)))
	// Output:
	// true
	// false
}

func ExampleSetMinutes() {
	fmt.Println(futureDate)
	fmt.Println(SetMinutes(futureDate, 59))
	fmt.Println(SetMinutes(futureDate, 60))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:59:00 +0000 UTC <nil>
	// 3001-01-01 00:00:00 +0000 UTC Passed amount was less than 0 or more than 59. Date left unchanged.
}

func ExampleStartOfMinute() {
	fmt.Println(third)
	fmt.Println(StartOfMinute(third))
	// Output:
	// 2016-06-06 06:06:06.000000007 +0000 UTC
	// 2016-06-06 06:06:00 +0000 UTC
}

func ExampleAddHours() {
	fmt.Println(futureDate)
	fmt.Print(AddHours(futureDate, 23))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 23:00:00 +0000 UTC
}

func ExampleDifferenceInHours() {
	fmt.Println(futureDate)
	fmt.Println(AddHours(futureDate, 23))
	fmt.Println(DifferenceInHours(AddHours(futureDate, 23), futureDate))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 23:00:00 +0000 UTC
	// 23
}

func ExampleEndOfHour() {
	fmt.Println(futureDate)
	fmt.Println(EndOfHour(futureDate))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:59:59.999999999 +0000 UTC
}

func ExampleIsSameHour() {
	fmt.Println(IsSameHour(futureDate, EndOfHour(futureDate)))
	fmt.Println(IsSameHour(futureDate, AddHours(futureDate, 1)))
	// Output:
	// true
	// false
}

func ExampleIsThisHour() {
	fmt.Println(IsThisHour(time.Now()))
	fmt.Println(IsThisHour(AddHours(time.Now(), 1)))
	// Output:
	// true
	// false
}

func ExampleSetHours() {
	fmt.Println(futureDate)
	fmt.Println(SetHours(futureDate, 23))
	fmt.Println(SetHours(futureDate, 24))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 23:00:00 +0000 UTC <nil>
	// 3001-01-01 00:00:00 +0000 UTC Passed amount was less than 0 or more than 23. Date left unchanged.
}

func ExampleStartOfHour() {
	fmt.Println(third)
	fmt.Println(StartOfHour(third))
	// Output:
	// 2016-06-06 06:06:06.000000007 +0000 UTC
	// 2016-06-06 06:00:00 +0000 UTC
}

func ExampleLastDayOfYear() {
	fmt.Println(first)
	fmt.Println(LastDayOfYear(first))
	// Output:
	// 2017-01-01 00:00:00 +0000 UTC
	// 2017-12-31 00:00:00 +0000 UTC
}

func ExampleFirstDayOfNextYear() {
	fmt.Println(first)
	fmt.Println(FirstDayOfNextYear(first))
	// Output:
	// 2017-01-01 00:00:00 +0000 UTC
	// 2018-01-01 00:00:00 +0000 UTC
}

func ExampleDaysDifference() {
	fmt.Println(third)
	fmt.Println(LastDayOfYear(third))
	fmt.Println(DaysDifference(LastDayOfYear(third), third))
	fmt.Println(DaysDifference(third, LastDayOfYear(third)))
	// Output:
	// 2016-06-06 06:06:06.000000007 +0000 UTC
	// 2016-12-31 00:00:00 +0000 UTC
	// 208
	// -208
}
