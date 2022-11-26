package maps

func Copy(a map[int]struct{}) map[int]struct{} {
	b := make(map[int]struct{})
	for k, v := range a {
		b[k] = v
	}
	return b
}
