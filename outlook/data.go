package outlook

import (
	"encoding/json"
	"log"
	"time"
)

// Mon Jan 01 2001 00:00:00 GMT+0000
const startPrefix int64 = 1702564200 - 724257000

type Data struct {
	DayToAppointments DayToAppointments `json:"dayToAppointments"`
	DateInterval      DateInterval      `json:"dateInterval"`
}

type DayToAppointments struct {
	Value map[float64][]Appointment
}

func (d *DayToAppointments) UnmarshalJSON(data []byte) error {
	var raw []interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	d.Value = map[float64][]Appointment{}

	var day float64
	for _, item := range raw {
		switch item.(type) {
		case float64:
			day = item.(float64)
			d.Value[day] = []Appointment{}
		case []interface{}:
			for _, appointmentInterface := range item.([]interface{}) {
				appointmentJson, err := json.Marshal(appointmentInterface)
				if err != nil {
					log.Printf("got error: %s\n", err)
					continue
				}

				var appointment Appointment
				if err := json.Unmarshal(appointmentJson, &appointment); err != nil {
					log.Printf("got error: %s\n", err)
					continue
				}

				appointment.StartsAt = time.Unix(startPrefix+appointment.StartTime, 0)
				d.Value[day] = append(d.Value[day], appointment)
			}
		}
	}

	return nil
}

type DateInterval struct {
	Duration int `json:"duration"`
	Start    int `json:"start"`
}

type Appointment struct {
	FreeBusyStatus int    `json:"freeBusyStatus"`
	ResponseStatus int    `json:"responseStatus"`
	Subject        string `json:"subject"`
	IsAllDay       bool   `json:"isAllDay"`
	EndTime        int64  `json:"endTime"`
	WithJoinAction bool   `json:"withJoinAction"`
	StartTime      int64  `json:"startTime"`
	StartsAt       time.Time
	Calendar       struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Color struct {
			Alpha   int `json:"alpha"`
			HexCode int `json:"hexCode"`
		} `json:"color"`
	} `json:"calendar"`
	Category struct {
		Color struct {
			Alpha   int `json:"alpha"`
			HexCode int `json:"hexCode"`
		} `json:"color"`
	} `json:"category"`
	IsCancelled         bool   `json:"isCancelled"`
	StableAppointmentID string `json:"stableAppointmentID"`
	Location            string `json:"location"`
	AccountID           int64  `json:"accountID"`
}
