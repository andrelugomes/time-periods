package periods

import "fmt"

//Comparison struct using by Periods
type Comparison struct {
	First  Period
	Second Period
}

//First date IsBigger than Second one?
func (c *Comparison) IsBigger() bool {
	return (c.First.Start.Before(c.Second.Start) && c.First.End.After(c.Second.End)) || (c.First.Start.Equal(c.Second.Start) && c.First.End.After(c.Second.End)) || (c.First.Start.Before(c.Second.Start) && c.First.End.Equal(c.Second.End))
}

//First date IsSame the Second one?
func (c *Comparison) IsSame() bool {
	return c.First.Start.Equal(c.Second.Start) && c.First.End.Equal(c.Second.End)
}

//First date IsDifferent of Second one?
func (c *Comparison) IsDifferent() bool {
	return (c.First.Start.Before(c.Second.Start) && c.First.End.Before(c.Second.Start)) || (c.First.Start.After(c.Second.End) && c.First.End.After(c.Second.End))
}

//First date IsInside of Second one?
func (c *Comparison) IsInside() bool {
	return (c.First.Start.After(c.Second.Start) && c.First.Start.Before(c.Second.End)) && (c.First.End.After(c.Second.Start) && c.First.End.Before(c.Second.End))
}

//First date IsCross by Second one?
func (c *Comparison) IsCross() bool {
	return c.IsLeftCross() || c.IsRightCross()
}

//First date IsLeftCross by Second one?
func (c *Comparison) IsLeftCross() bool {
	return c.First.Start.After(c.Second.Start) && c.First.End.After(c.Second.Start) && (c.First.End.After(c.Second.End) || c.First.End.Equal(c.Second.End))
}

//First date IsRightCross by Second one?
func (c *Comparison) IsRightCross() bool {
	return c.First.End.After(c.Second.Start) && c.First.Start.Before(c.Second.End) && (c.First.Start.Before(c.Second.Start) || c.First.Start.Equal(c.Second.Start))
}

//Print like a to_string
func (c *Comparison) Print() {
	layout := "02/01/2006 15:04:05"
	fmt.Println("c.First.Start", c.First.Start.Format(layout))
	fmt.Println("c.First.End", c.First.End.Format(layout))
	fmt.Println("c.Second.Start", c.Second.Start.Format(layout))
	fmt.Println("c.Second.End", c.Second.End.Format(layout))
}

