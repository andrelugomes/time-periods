package periods

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsBigger(t *testing.T) {

	d1 := time.Date(2009, 12, 12, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(2009, 12, 22, 0, 0, 0, 0, time.UTC)

	p1 := Period{d1, d2}

	d3 := time.Date(2009, 12, 15, 0, 0, 0, 0, time.UTC)
	d4 := time.Date(2009, 12, 16, 0, 0, 0, 0, time.UTC)

	p2 := Period{d3, d4}

	period1 := Comparison{p1, p2}

	assert.True(t, period1.IsBigger(), "P1 is bigger than P2")

	period2 := Comparison{p2, p1}

	assert.False(t, period2.IsBigger(), "P2 is shorter than P1")

}

func TestIsLeftCross(t *testing.T) {

	dates1 := Period{time.Date(2009, 12, 01, 0, 0, 0, 0, time.UTC), time.Date(2009, 12, 03, 0, 0, 0, 0, time.UTC)}
	dates2 := Period{time.Date(2009, 11, 30, 0, 0, 0, 0, time.UTC), time.Date(2009, 12, 02, 0, 0, 0, 0, time.UTC)}

	period := Comparison{dates1, dates2}

	assert.True(t, period.IsLeftCross(), "no left cross")
}

func TestIsLeftCrossWhenEndsInSameTime(t *testing.T) {
	//2019-09-20 10:57:16 +0000  - 2019-09-23 15:22:12 +0000
	dates1 := Period{
		Start: time.Date(2019, 9, 20, 10, 57, 16, 0, time.UTC),
		End:   time.Date(2019, 9, 23, 15, 22, 12, 0, time.UTC),
	}
	//2019-09-17 16:47:56 +0000 - 2019-09-23 15:22:12 +0000
	dates2 := Period{
		Start: time.Date(2019, 9, 17, 16, 47, 56, 0, time.UTC),
		End: time.Date(2019, 9, 23, 15, 22, 12, 0, time.UTC),
	}

	period := Comparison{dates1, dates2}

	assert.True(t, period.IsLeftCross(), "no left cross")
}

func TestIsRightCrossWhenStartsInSameTime(t *testing.T) {
	//2019-09-20 10:57:16 +0000  - 2019-09-23 15:22:12 +0000
	dates1 := Period{
		Start: time.Date(2019, 9, 20, 10, 57, 16, 0, time.UTC),
		End:   time.Date(2019, 9, 23, 15, 22, 12, 0, time.UTC),
	}
	//2019-09-17 16:47:56 +0000 - 2019-09-23 15:22:12 +0000
	dates2 := Period{
		Start: time.Date(2019, 9, 20, 10, 57, 16, 0, time.UTC),
		End: time.Date(2019, 9, 25, 11, 11, 11, 0, time.UTC),
	}

	period := Comparison{dates1, dates2}

	assert.True(t, period.IsRightCross(), "no right cross")
}

func TestIsBiggerBySecondsLater(t *testing.T) {

	dates1 := Period{time.Date(2009, 12, 01, 0, 0, 0, 0, time.UTC), time.Date(2009, 12, 02, 0, 0, 0, 0, time.UTC)}
	dates2 := Period{time.Date(2009, 12, 01, 0, 0, 0, 0, time.UTC), time.Date(2009, 12, 01, 23, 59, 30, 0, time.UTC)}

	period := Comparison{dates1, dates2}

	assert.True(t, period.IsBigger(), "P1 isn't bigger than P2 by 30 seconds")
}
