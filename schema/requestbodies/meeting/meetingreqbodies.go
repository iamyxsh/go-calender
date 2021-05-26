package meetingreqbodies

type CreateMeetingReq struct {
	Year      int `json:"year" validate:"max=2050,min=2020,required"`
	Month     int `json:"month" validate:"max=11,min=0,required"`
	Day       int `json:"day" validate:"max=31,min=1,required"`
	StartHour int `json:"starthour" validate:"max=24,min=0,required"`
	StartMin  int `json:"startmin" validate:"max=60,min=0,required"`
	EndHour   int `json:"endhour" validate:"max=24,min=0,required"`
	EndMin    int `json:"endmin" validate:"max=60,min=0,required"`
	Slot      int `json:"slot" validate:"max=60,min=0,required"`
}

type CreateSlotReq struct {
	Name      string `json:"name" validate:"max=20,min=5,required"`
	Email     string `json:"email" validate:"max=30,min=5,required,email"`
	Start     string `json:"start" validate:"required"`
	MeetingId string `json:"meetingid" validate:"required"`
}
