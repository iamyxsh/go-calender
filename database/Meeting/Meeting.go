package Meeting

import (
	"calendly/schema/entities"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeetingUser struct {
	mgm.DefaultModel `bson:",inline"`
	StartTime        time.Time       `json:"start" bson:"start"`
	EndTime          time.Time       `json:"end" bson:"end"`
	Interval         time.Duration   `json:"interval" bson:"interval"`
	User             []entities.User `json:"user" bson:"user"`
}

func CreateMeeting(userid primitive.ObjectID, start time.Time, end time.Time, interval time.Duration, slots []time.Time) error {
	meeting := &entities.Meeting{
		UserID:    userid,
		StartTime: start,
		EndTime:   end,
		Interval:  interval,
		Slots:     slots,
	}

	return mgm.Coll(meeting).Create(meeting)
}

func FindMeetingById(id primitive.ObjectID) (MeetingUser, error) {
	u := new(entities.User)
	meet := new(entities.Meeting)
	meetArr := new([]MeetingUser)
	username := mgm.Coll(u).Name()

	err := mgm.Coll(meet).SimpleAggregate(meetArr, bson.M{operator.Match: bson.M{"_id": id}}, builder.Lookup(username, "userid", "_id", "user"), bson.M{operator.Project: bson.M{"user.password": 0}})
	m := *meetArr

	return m[0], err
}
