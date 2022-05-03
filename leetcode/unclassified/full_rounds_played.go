package main

// 1904. The Number of Full Rounds You Have Played

import (
	"fmt"
	"strconv"
)

func numberOfRounds(loginTime string, logoutTime string) int {
	login := calMinute(loginTime)
	logout := calMinute(logoutTime)

	loginMinute := loginTominute(loginTime)
	logoutMinute := logoutTominute(logoutTime)
	if logout >= login && logoutMinute >= loginMinute {
		return (logoutMinute - loginMinute) / 15
	} else if logout >= login && logoutMinute < loginMinute {
		return 0
	} else {
		return (logoutMinute + 24*60 - loginMinute) / 15
	}
}

func calMinute(timeStamp string) int {
	h, _ := strconv.Atoi(timeStamp[:2])
	m, _ := strconv.Atoi(timeStamp[3:])
	return h*60 + m
}

func loginTominute(loginTime string) int {
	// h_and_m := strings.Split(loginTime, ":")
	h, _ := strconv.Atoi(loginTime[:2])
	m, _ := strconv.Atoi(loginTime[3:])
	if m <= 0 {
		m = 0
	} else if m <= 15 {
		m = 15
	} else if m <= 30 {
		m = 30
	} else if m <= 45 {
		m = 45
	} else {
		m = 0
		h++
	}

	return h*60 + m
}
func logoutTominute(logoutTime string) int {
	// h_and_m := strings.Split(loginTime, ":")
	h, _ := strconv.Atoi(logoutTime[:2])
	m, _ := strconv.Atoi(logoutTime[3:])
	if m >= 45 {
		m = 45
	} else if m >= 30 {
		m = 30
	} else if m >= 15 {
		m = 15
	} else {
		m = 0
	}

	return h*60 + m
}

func main() {
	res := numberOfRounds("21:47", "21:57")
	fmt.Println(res)
}
