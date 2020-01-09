package periods

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestINormalizeWhenIsBigger(t *testing.T) {

	p1 := Period{time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 06, 0, 0, 0, 0, time.UTC)}
	p2 := Period{time.Date(2020, 01, 02, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 05, 0, 0, 0, 0, time.UTC)}

	dates := []Period{p1,p2}

	unique := Normalize(dates)

	assert.Equal(t, unique[0].Start , p1.Start, "Result needs to be equal p1.Start")
	assert.Equal(t, unique[0].End , p1.End, "Result needs to be equal p2.End")
}

func TestNormalizeWhenIsInside(t *testing.T) {

	p1 := Period{time.Date(2020, 01, 03, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 04, 0, 0, 0, 0, time.UTC)}
	p2 := Period{time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 06, 0, 0, 0, 0, time.UTC)}

	dates := []Period{p1,p2}

	unique := Normalize(dates)

	assert.Equal(t, unique[0] , p2, "Result needs to be equal to p2")
}

func TestNormalizeWhenIsBiggerThanIsInside(t *testing.T) {

	p1 := Period{time.Date(2020, 01, 02, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 06, 0, 0, 0, 0, time.UTC)}
	p2 := Period{time.Date(2020, 01, 03, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 04, 0, 0, 0, 0, time.UTC)}
	p3 := Period{time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 10, 0, 0, 0, 0, time.UTC)}

	dates := []Period{p1,p2,p3}

	unique := Normalize(dates)

	assert.Equal(t, unique[0] , p3, "Result needs to be equal to p3")

}

func TestNormalizeWhenIsRightCross(t *testing.T) {

	p1 := Period{time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 03, 0, 0, 0, 0, time.UTC)}
	p2 := Period{time.Date(2020, 01, 02, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 04, 0, 0, 0, 0, time.UTC)}

	dates := []Period{p1,p2}

	unique := Normalize(dates)

	assert.Equal(t, unique[0].Start , p1.Start, "Result needs to be equal to p1.Start")
	assert.Equal(t, unique[0].End , p2.End, "Result needs to be equal to p2.End")
}

func TestNormalizeWhenIsLeftCross(t *testing.T) {

	p1 := Period{time.Date(2020, 01, 03, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 06, 0, 0, 0, 0, time.UTC)}
	p2 := Period{time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 04, 0, 0, 0, 0, time.UTC)}

	dates := []Period{p1,p2}

	unique := Normalize(dates)

	assert.Equal(t, unique[0].Start , p2.Start, "Result needs to be equal to p1.Start")
	assert.Equal(t, unique[0].End , p1.End, "Result needs to be equal to p1.Start")
}

func TestNormalizeWhenIsSame(t *testing.T) {

	p1 := Period{time.Date(2020, 01, 03, 0, 0, 0, 0, time.UTC), time.Date(2020, 01, 06, 0, 0, 0, 0, time.UTC)}

	dates := []Period{p1,p1}

	unique := Normalize(dates)

	assert.Equal(t, unique[0], p1, "Result needs to be equal to p1")
}

func TestRemoveWeekend(t *testing.T) {

	period := Period{
		time.Date(2020, 01, 02, 11, 10, 9, 0, time.UTC),
		time.Date(2020, 01, 06, 0, 0, 0, 0, time.UTC),
	}

	days := WorkingDays(period)

	assert.Equal(t, 3, days, "Three days")
}

func TestRemoveWeekendEndingOnSunday(t *testing.T) {
	period1 := Period{
		time.Date(2020, 01, 01, 11, 10, 9, 0, time.UTC),
		time.Date(2020, 01, 05, 22, 0, 0, 0, time.UTC),
	}

	days1 := WorkingDays(period1)

	assert.Equal(t, 3, days1, "Three days")
}

func TestWeekends(t *testing.T) {
	period1 := Period{
		time.Date(2020, 01, 03, 11, 10, 9, 0, time.UTC),
		time.Date(2020, 01, 06, 22, 0, 0, 0, time.UTC),
	}

	w := Weekends(period1)

	assert.Equal(t, 2, w, "Weekend with 2 days")
}
