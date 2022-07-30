package utils

import "strconv"

func converStringToInt(vars string) int {
	//convert age string to int
	age, err := strconv.Atoi(vars)
	if err != nil {
		panic(err)
	}
	return age
}
