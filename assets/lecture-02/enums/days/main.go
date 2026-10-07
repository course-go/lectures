package main

import "fmt"

// START OMIT

type Day string

const (
	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"
	Sunday    Day = "sunday"
)

func currentDay() Day {
	return Wednesday
}

func main() {
	fmt.Println(currentDay())
	madeupDay := Day("redemption-day") // note that it does not prevent
	fmt.Println(madeupDay)             // you from creating new values
}

// END OMIT
