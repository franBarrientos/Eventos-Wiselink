package utils

import (
	"strconv"
)

func PageOrDefault(page string, defaultValue int) int {
	if page == "" {
		return defaultValue
	}
	number, err := strconv.Atoi(page)

	if err != nil {
		return defaultValue
	}
	return number
}

func LimitOrDefault(limit string, defaultValue int) int {
	if limit == "" {
		return defaultValue
	}
	number, err := strconv.Atoi(limit)

	if err != nil {
		return defaultValue
	}
	return number
}
