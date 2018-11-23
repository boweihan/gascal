package main

import (
	"strconv"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func getNumber(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
	}
	return val
}
