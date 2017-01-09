# This is a free time helper library for Go.

### Feel free to use it any way you like it :)
### It should soon have more than 140 function to use

Small examples:

```go
  var (
  	first  = time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
  	second = time.Date(2016, 6, 6, 6, 6, 6, 6, time.UTC)
  	third  = time.Date(2016, 6, 6, 6, 6, 6, 7, time.UTC)
  	fourth = time.Date(2015, 1, 1, 1, 0, 0, 0, time.UTC)

  	futureDate = time.Date(3001, 1, 1, 0, 0, 0, 0, time.UTC)
  	pastDate   = time.Date(1001, 1, 1, 0, 0, 0, 0, time.UTC)
  )

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

  timeSlice := []time.Time{fourth, second.Add(-1 * time.Nanosecond), third.Add(time.Nanosecond), first}
  fmt.Println(ClosestTo(third, timeSlice))
  fmt.Println(ClosestTo(third, nil))
  fmt.Println(ClosestTo(third, []time.Time{}))
  // Output:
  // 2016-06-06 06:06:06.000000008 +0000 UTC <nil>
  // 0001-01-01 00:00:00 +0000 UTC Passed slice of dates was nil
  // 0001-01-01 00:00:00 +0000 UTC Passed slice of dates was of size 0

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
```
