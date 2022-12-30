package main

func Find[T any](array []T, fn func(T) bool) T {
	for _, v := range array {
		if fn(v) {
			return v
		}
	}

	return *new(T)
}

func Map[T, V any](vs []T, f func(T) V) []V {
	vsm := make([]V, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func GenerateEmpty2DMatrix(length int) [][]int {
	matrix := make([][]int, length)

	for i := range matrix {
		matrix[i] = make([]int, length)
	}

	return matrix
}
