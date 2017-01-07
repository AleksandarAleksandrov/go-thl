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
	// 0 Ranges do no overlap
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
	fmt.Println(SetMillisecond(futureDate, 1001))
	// Output:
	// 3001-01-01 00:00:00 +0000 UTC
	// 3001-01-01 00:00:00.999 +0000 UTC <nil>
	// 3001-01-01 00:00:00 +0000 UTC Passed amount was less than 0 or more than 1000. Date left unchanged.
}

func someShitExample() {
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
}
