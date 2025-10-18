package main

import "fmt"

// START OMIT

type Day string

const (
	Monday    Day = "monday"
	Tuesday       = "tuesday"
	Wednesday     = "wednesday"
	Thursday      = "thursday"
	Friday        = "friday"
	Saturday      = "saturday"
	Sunday        = "sunday"
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
