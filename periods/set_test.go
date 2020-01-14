package periods

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	period1 := Period{
		time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 01, 02, 0, 0, 0, 0, time.UTC),
	}

	period2 := Period{
		time.Date(2020, 01, 03, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 01, 04, 0, 0, 0, 0, time.UTC),
	}

	var set Set

	hashSet := set.New([]Period{period1, period2})

	assert.NotNil(t, hashSet, "Set must be not null")
}
