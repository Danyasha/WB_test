package main

import (
	"regexp"
	"strconv"
	"strings"
)

func isLong(month int) bool {
	longMonths := [7]int{1, 3, 5, 7, 8, 10, 11}
	for _, i := range longMonths {
		if i == month {
			return true
		}
	}
	return false
}

func isLeap(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

func isValidTime(date string) bool {
	splited := strings.Split(date, "-")
	year, err := strconv.Atoi(splited[0])
	if err != nil {
		return false
	}
	month, err := strconv.Atoi(splited[1])
	if err != nil {
		return false
	}
	day, err := strconv.Atoi(splited[2])
	if err != nil {
		return false
	}
	if day < 28 || (month != 2 && day < 30) {
		return true
	}
	if day == 31 && isLong(month) {
		return true
	}
	if day == 29 && month == 2 && isLeap(year) {
		return true
	}
	return false
}

func checkDate(date string) (bool, error) {
	pattern := `^20[1-9][0-9]$`
	matched, err := regexp.Match(pattern, []byte(date))
	if err != nil {
		return false, err
	}
	if matched {
		return true, err
	}
	pattern = `^20[1-9][0-9]-([0][\d]|[1][012])-[\d]{2}$`
	matched, err = regexp.Match(pattern, []byte(date))
	if err != nil || !matched {
		return false, err
	}
	isTime := isValidTime(date)
	if isTime {
		return true, nil
	}
	return false, nil
}
