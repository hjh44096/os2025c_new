package main

import (
	"fmt"
	"week13/pkg/calendar"
)

func main() {
	today := calendar.Date{}
	//today.Year = 2025
	today.SetYear(2025)
	today.SetMonth(11)
	today.SetDay(24)
	fmt.Println(today.Year(), "년", today.Month(), "월", today.Day(), "일")
}
