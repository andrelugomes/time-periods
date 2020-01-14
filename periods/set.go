package periods

//Set struct like using map as a HashSet
type Set struct {
	hashSet map[int64]Period
}

//New HashSet of Periods
func (s *Set) New(p []Period) Set {
	set := Set{make(map[int64]Period)}

	for _, period := range p {
		set.hashSet[period.Hash()] = period
	}
	return set
}

func toSet(periods []Period) Set {
	var set Set
	return set.New(periods)
}

func toArray(set Set) []Period {
	var periods []Period
	for _, period := range set.hashSet {
		periods = append(periods, period)
	}
	return periods
}

//Deduplicate is useful to convert a simple array to an array of unique values
func Deduplicate(periods []Period) []Period {
	return toArray(toSet(periods))
}
