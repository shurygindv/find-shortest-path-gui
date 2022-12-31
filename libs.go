package main

import "strconv"

func Find[T any](array []T, fn func(T) bool) *T {
	for _, v := range array {
		if fn(v) {
			return &v
		}
	}

	return new(T)
}

func Find2[T any](array []T, fn func(T) bool) (*T, bool) {
	result := Find(array, fn)

	return result, (result == new(T))
}

func Map[T, V any](vs []T, f func(T) V) []V {
	vsm := make([]V, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Filter[T any](s []T, cond func(t T) bool) []T {
	res := []T{}
	for _, v := range s {
		if cond(v) {
			res = append(res, v)
		}
	}
	return res
}

func Contains(array []int, contained int) bool {
	for _, a := range array {
		if a == contained {
			return true
		}
	}
	return false
}

func GenerateEmpty2DMatrix(length int) [][]int {
	matrix := make([][]int, length)

	for i := range matrix {
		matrix[i] = make([]int, length)
	}

	return matrix
}

func GenerateNodeOptionsForSelect(nodesCount int) ([]string, map[string]int) {
	options := make([]string, nodesCount)
	optionsMap := make(map[string]int)

	for i := range options {
		optionLabel := strconv.Itoa(i + 1)

		options[i] = optionLabel
		optionsMap[optionLabel] = i
	}

	return options, optionsMap
}
