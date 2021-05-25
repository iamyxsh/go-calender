package entities

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
}

type Meeting struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           primitive.ObjectID `json:"userid" bson:"userid"`
	StartTime        time.Time          `json:"start" bson:"start"`
	EndTime          time.Time          `json:"end" bson:"end"`
	Interval         time.Duration      `json:"interval" bson:"interval"`
}
