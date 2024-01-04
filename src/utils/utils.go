package utils

func Map[X any, Y any](list []X, mapFunc func(X) Y) []Y {
	mappedList := make([]Y, len(list))
	for i, elem := range list {
		mappedList[i] = mapFunc(elem)
	}
	return mappedList
}

func Filter[X any](list []X, filterFunc func(X) bool) []X {
	filteredList := make([]X, 0)
	for _, elem := range list {
		if filterFunc(elem) {
			filteredList = append(filteredList, elem)
		}
	}
	return filteredList
}

func IsStringEmpty(s string) bool {
	return len(s) <= 0
}

func IsStringNotEmpty(s string) bool {
	return len(s) > 0
}
