package Meeting

import (
	"calendly/schema/entities"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMeeting(userid primitive.ObjectID, start time.Time, end time.Time, interval time.Duration) error {
	meeting := &entities.Meeting{
		UserID:    userid,
		StartTime: start,
		EndTime:   end,
		Interval:  interval,
	}

	return mgm.Coll(meeting).Create(meeting)
}
