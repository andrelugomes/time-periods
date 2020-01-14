# Time Periods

Create periods of Time. Normalize them for custom calculations.

## Install

```bash
go get github.com/andrelugomes/time-periods
```

## Import

```go
import "github.com/andrelugomes/time-periods/periods"
```

## Usage

### Period

```go
period := Period{
		time.Date(2020, 01, 02, 11, 0, 0, 0, time.UTC),
		time.Date(2020, 01, 06, 0, 0, 0, 0, time.UTC),
	}
```

#### Normalize

```go
periods := []Period{p1,p2}

uniques := Normalize(periods)
```

### Comparison

```go
comparison := Comparison{first, second}

comparison.Is...
```




