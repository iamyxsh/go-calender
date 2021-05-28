package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/calendar/v3"
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

func CheckTime(arr []time.Time, slot time.Time) bool {
	for _, i := range arr {
		if i == slot {
			return true
		}
	}
	return false
}

func CreateSlots(start time.Time, slot time.Duration, end time.Time) []time.Time {
	temp := time.Date(0001, 1, 1, 00, 00, 00, 00, time.UTC)
	var slotArr []time.Time

	slotArr = append(slotArr, start)

	for end.Unix() > temp.Unix() {
		if temp.IsZero() {
			temp = start.Add(slot)
		}
		temp = temp.Add(slot)
		if end.Unix() > temp.Unix() || end.Unix() == temp.Unix() {
			slotArr = append(slotArr, temp)
		}
	}
	return slotArr
}

func CreateInvite(c *fiber.Ctx) error {
	ctx := context.Background()

	calSer, err := calendar.NewService(ctx)
	if err != nil {
		fmt.Println(err)
	}

	event := &calendar.Event{
		Summary:     "Noice",
		Location:    "Noice",
		Description: "Noice",
		Start: &calendar.EventDateTime{
			DateTime: time.Date(2021, time.August, 3, 1, 0, 0, 0, time.Local).String(),
			TimeZone: "India/Kolkata",
		},
		End: &calendar.EventDateTime{
			DateTime: time.Date(2021, time.August, 3, 2, 0, 0, 0, time.Local).String(),
			TimeZone: "India/Kolkata",
		},
		Attendees: []*calendar.EventAttendee{
			&calendar.EventAttendee{Email: "anmolsneha2312@gmail.com"},
			&calendar.EventAttendee{Email: "yashsharma170898@gmail.com"},
		},
	}

	event, err = calSer.Events.Insert("n1r4l2k9q6v9gatctn77flbnrk@group.calendar.google.com", event).Do()
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString(event.HtmlLink)
}
