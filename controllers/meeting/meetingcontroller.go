package meetingcontroller

import (
	"calendly/database/Meeting"
	"calendly/database/Slot"
	meetingreqbodies "calendly/schema/requestbodies/meeting"
	"calendly/schema/resbodies"
	"calendly/utils"
	errorhandler "calendly/utils/error"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateLink(c *fiber.Ctx) error {
	body := c.Locals("body").(*meetingreqbodies.CreateMeetingReq)

	userid, err := primitive.ObjectIDFromHex(c.Locals("userid").(string))
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	month, err := utils.ReturnMonth(body.Month)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	start := time.Date(body.Year, month, body.Day, body.StartHour, body.StartMin, 0, 0, time.UTC)
	end := time.Date(body.Year, month, body.Day, body.EndHour, body.EndMin, 0, 0, time.UTC)

	interval, err := time.ParseDuration(fmt.Sprintf("%vm", body.Slot))
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	slotArr := utils.CreateSlots(start, interval, end)

	err = Meeting.CreateMeeting(userid, start, end, interval, slotArr)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	return c.Status(fiber.StatusCreated).JSON(resbodies.SuccessRes("msg", "Meeting created successfully"))
}

func GetLinkDetails(c *fiber.Ctx) error {
	linkID := c.Params("id")
	id, err := primitive.ObjectIDFromHex(linkID)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}
	meeting, err := Meeting.FindMeetingById(id)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(meeting)
}

func CreateSlot(c *fiber.Ctx) error {
	body := c.Locals("body").(*meetingreqbodies.CreateSlotReq)

	meetingid, err := primitive.ObjectIDFromHex(body.MeetingId)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	slot, err := time.Parse(time.RFC3339, body.Start)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	err = Slot.CreateSlot(meetingid, body.Name, body.Email, slot)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	meeting, err := Meeting.FindMeetingById(meetingid)
	if err != nil {
		return errorhandler.SendErr(err, c)
	}

	ok := utils.CheckTime(meeting.Slots, slot)

	if !ok {
		e := errors.New("slot is not valid")
		return errorhandler.SendErr(e, c)
	}

	return c.Status(fiber.StatusCreated).JSON(resbodies.SuccessRes("msg", "Slot created successfully."))
}
