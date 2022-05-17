package learning

import (
	"fmt"
	"time"
)

func ConditionalExp() {

	today := time.Now()

	if today.Weekday() == time.Saturday || today.Weekday() == time.Sunday { // if any condition is true
		fmt.Println("happy weekend")
	} else if today.Month() == time.January && today.Day() == 1 { // both condition must be true
		fmt.Println("happy new year")
	} else {
		fmt.Println("Happy", today.Weekday())
	}
}
