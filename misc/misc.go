package misc

func SliceContains(s []int, n int) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}

	return false
}
