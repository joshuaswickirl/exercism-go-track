package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(f func(int) bool) Ints {
	var kept Ints
	for _, int := range i {
		if f(int) {
			kept = append(kept, int)
		}
	}
	return kept
}

func (i Ints) Discard(f func(int) bool) Ints {
	var discarded Ints
	for _, int := range i {
		if !f(int) {
			discarded = append(discarded, int)
		}
	}
	return discarded
}

func (l Lists) Keep(f func([]int) bool) Lists {
	var kept Lists
	for _, list := range l {
		if f(list) {
			kept = append(kept, list)
		}
	}
	return kept
}

func (s Strings) Keep(f func(string) bool) Strings {
	var kept Strings
	for _, str := range s {
		if f(str) {
			kept = append(kept, str)
		}
	}
	return kept
}
