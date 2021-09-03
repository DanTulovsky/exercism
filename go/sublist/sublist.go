package sublist

// Relation is one of: equal, sublist, superlist, unequal
type Relation string

func Sublist(one, two []int) Relation {

	if listsEqual(one, two) {
		return "equal"
	}

	if sublist(one, two) {
		return "sublist"
	}

	if sublist(two, one) {
		return "superlist"
	}

	return "unequal"
}

func listsEqual(one, two []int) bool {
	if len(one) != len(two) {
		return false
	}

	if len(one) == 0 {
		return true
	}

	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}

	return true
}

// returns true if one is a sublist of two
func sublist(one, two []int) bool {
	if len(one) > len(two) {
		return false
	}

	if len(one) == 0 {
		return true
	}

	for i := range two {
		if one[i] != two[i] {
			if len(two) == 1 {
				return false
			}
			return sublist(one, two[1:])
		}

		if i == len(one)-1 {
			return true
		}
	}

	return true
}
