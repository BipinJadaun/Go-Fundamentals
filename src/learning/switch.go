package learning

import (
	"fmt"
	"time"
)

func SwitchExp() {

	today := time.Now().Weekday()

	switch today {
	case time.Monday:
		fmt.Println("Monday: one day")
	case time.Tuesday:
		fmt.Println("Tuesday: two day")
	case time.Wednesday:
		fmt.Println("Wednesdaty: when, what, which day")
	case time.Thursday:
		fmt.Println("Thursday: the third day")
	case time.Friday:
		fmt.Println("Thank god its friday")
	default:
		fmt.Println("yeehaw, its weekend!")

	}
}
