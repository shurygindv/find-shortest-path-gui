package main

func find[T any](array []T, fn func(T) bool) T {
	for _, v := range array {
		if fn(v) {
			return v
		}
	}

	return *new(T)
}

func filter[T any](s []T, cond func(t T) bool) []T {
	res := []T{}
	for _, v := range s {
		if cond(v) {
			res = append(res, v)
		}
	}
	return res
}
