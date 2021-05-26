package Slot

import (
	"calendly/schema/entities"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateSlot(meetingid primitive.ObjectID, name string, email string, start time.Time) error {
	slot := &entities.Slot{
		MeetingID: meetingid,
		Name:      name,
		Email:     email,
		Start:     start,
	}

	return mgm.Coll(slot).Create(slot)
}
