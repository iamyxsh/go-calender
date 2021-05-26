package utils

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func ReturnMonth(m int) (time.Month, error) {
	switch m {
	case 0:
		return time.January, nil
	case 1:
		return time.February, nil
	case 2:
		return time.March, nil
	case 3:
		return time.April, nil
	case 4:
		return time.May, nil
	case 5:
		return time.June, nil
	case 6:
		return time.July, nil
	case 7:
		return time.August, nil
	case 8:
		return time.September, nil
	case 9:
		return time.October, nil
	case 10:
		return time.November, nil
	case 11:
		return time.December, nil
	default:
		return time.January, errors.New("wrong number")
	}
}

// func CheckTime(start time.Time, slot time.Time, end time.Time) {
// 	startSlot := start.Sub(slot)
// 	endSlot := slot.Sub()
// 	fmt.Println(startSlot)
// 	fmt.Println(math.Signbit(float64(startSlot)))
// }

func CreateSlots(start time.Time, slot time.Duration, end time.Time) {
	temp := time.Date(0001, 1, 1, 00, 00, 00, 00, time.UTC)
	var slotArr []time.Time

	slotArr = append(slotArr, start)

	fmt.Println(math.Signbit(float64(temp.Sub(end))))

	// start.Add(slot).Add(slot).Add(slot).Unix()

	fmt.Println(end.Unix() > start.Add(slot).Add(slot).Add(slot).Unix())

	for end.Unix() > temp.Unix() {
		if temp.IsZero() {
			temp = start.Add(slot)
		}
		temp = temp.Add(slot)
		if end.Unix() > temp.Unix() || end.Unix() == temp.Unix() {
			slotArr = append(slotArr, temp)
		}
	}

	for s, t := range slotArr {
		fmt.Println(t, s)
	}
}
