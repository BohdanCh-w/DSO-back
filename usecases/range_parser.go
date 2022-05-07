package usecases

import (
	"math"
	"regexp"
	"strconv"
)

func ParsePI(input string) (float64, error) {
	match, err := regexp.MatchString("(?i)pi", input)
	if err != nil {
		return 0, err
	}

	if !match {
		return strconv.ParseFloat(input, 64)
	}

	match, err = regexp.MatchString("^(?i)pi$", input)
	if err != nil {
		return 0, err
	}

	if match {
		return math.Pi, nil
	}

	val, err := strconv.ParseFloat(input[:len(input)-2], 64)
	return val * math.Pi, err
}
