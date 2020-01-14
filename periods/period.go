package periods

import (
	"time"
)

//Period struct composed by a start Time and a end time
type Period struct {
	Start time.Time
	End   time.Time
}

//Hash this period with an unique int64
func (p *Period) Hash() int64{
	return p.Start.UnixNano() + p.End.Unix()
}

//Normalize two or more periods removing conflicts
func Normalize(periods []Period) []Period {
	Deduplicate(periods)
	if len(periods) > 1 {

		for i, date := range periods {
			for j, d := range periods {
				comparison := Comparison{First: date, Second: d}

				if comparison.IsBigger() {
					arr := remove(j, periods)
					return Normalize(arr)
				} else if comparison.IsSame() || comparison.IsDifferent() {
					continue
				} else if comparison.IsInside() {
					arr := remove(i, periods)
					return Normalize(arr)
				} else if comparison.IsLeftCross() {
					arr := delete(date, periods)
					arr = delete(d, arr)
					arr = append(arr, Period{Start: comparison.Second.Start, End: comparison.First.End})
					return Normalize(arr)
				} else if comparison.IsRightCross() {
					arr := delete(date, periods)
					arr = delete(d, arr)
					arr = append(arr, Period{Start: comparison.First.Start, End: comparison.Second.End})
					return Normalize(arr)
				} else {
					comparison.Print()
					panic("Unmapped comparision.")
				}
			}
		}
	}
	return periods
}

//WorkingDays get how many working days has
func WorkingDays(period Period) int{
	days := 0
	start := dateWithoutTime(period)
	for day := start; !day.After(period.End); day = day.AddDate(0, 0, 1) {
		if day.Weekday() != 6 && day.Weekday() != 0 {
			days++
		}
	}
	return days
}

//Weekends get how many weekends days has
func Weekends(period Period) int{
	days := 0
	start := dateWithoutTime(period)
	for day := start; !day.After(period.End); day = day.AddDate(0, 0, 1) {
		if day.Weekday() == 6 || day.Weekday() == 0 {
			days++
		}
	}
	return days
}

func dateWithoutTime(period Period) time.Time {
	start := time.Date(period.Start.Year(), period.Start.Month(), period.Start.Day(), 0, 0, 0, 0, time.UTC)
	return start
}

func remove(index int, periods []Period) []Period {
	periods[index] = periods[len(periods)-1] // Copy last element to index index.
	periods = periods[:len(periods)-1]       // Truncate slice.
	return periods
}

func delete(period Period, periods []Period) []Period {
	for i, p := range periods {
		comparison := Comparison{period, p}
		if comparison.IsSame() {
			periods = remove(i, periods)
		}
	}
	return periods
}

