package meetingrouter

import (
	meetingcontroller "calendly/controllers/meeting"
	"calendly/middlewares"
	"calendly/middlewares/validations"

	"github.com/gofiber/fiber/v2"
)

func MettingRouter(meeting *fiber.Router) {
	m := *meeting

	m.Post("/link", middlewares.IsAuth, validations.CheckCreateMeetingReq, meetingcontroller.GenerateLink)
	m.Get("/link/:id", meetingcontroller.GetLinkDetails)

	m.Post("/slot", validations.CheckCreateSlotReq, meetingcontroller.CreateSlot)
}
