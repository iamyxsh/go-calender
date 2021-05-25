package utils

import (
	"errors"
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
