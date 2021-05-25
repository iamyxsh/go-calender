package meetingcontroller

import (
	"calendly/database/Meeting"
	meetingreqbodies "calendly/schema/requestbodies/meeting"
	"calendly/schema/resbodies"
	"calendly/utils"
	errorhandler "calendly/utils/error"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateLink(c *fiber.Ctx) error {
	body := c.Locals("body").(*meetingreqbodies.CreateMeetingReq)
	// fmt.Print(c.Locals("userid"))

	userid, err := primitive.ObjectIDFromHex(c.Locals("userid").(string))
	errorhandler.CheckErr(err, c)

	month, err := utils.ReturnMonth(body.Month)
	errorhandler.CheckErr(err, c)

	start := time.Date(body.Year, month, body.Day, body.StartHour, body.StartMin, 0, 0, time.UTC)
	end := time.Date(body.Year, month, body.Day, body.EndHour, body.EndMin, 0, 0, time.UTC)

	slot, err := time.ParseDuration(fmt.Sprintf("%vm", body.Slot))

	errorhandler.CheckErr(err, c)

	err = Meeting.CreateMeeting(userid, start, end, slot)
	errorhandler.CheckErr(err, c)

	return c.Status(fiber.StatusCreated).JSON(resbodies.SuccessRes("msg", "Meeting created successfully"))
}
