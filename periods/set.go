package periods

type Set struct {
	hashSet map[int64]Period
}

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

func Deduplicate(periods []Period) []Period {
	return toArray(toSet(periods))
}
