package outlook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go-obsidian/note"
)

func Sync() {
	log.Println("loading outlook events...")

	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return
	}

	f, err := os.ReadFile(fmt.Sprintf("%s/Library/Group Containers/UBF8T346G9.Office/Library/Application Support/Calendar Widget/store.json", home))
	if err != nil {
		log.Println("got error", err)
		return
	}

	var result Data
	if err := json.Unmarshal(f, &result); err != nil {
		log.Println("got error", err)
		return
	}

	log.Println("formatting events...")

	outputToday := new(bytes.Buffer)
	outputTomorrow := new(bytes.Buffer)

	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	tomorrow := today.AddDate(0, 0, 1)
	dayAfterTomorrow := tomorrow.AddDate(0, 0, 1)

	for _, appointments := range result.DayToAppointments.Value {
		for _, appointment := range appointments {
			if appointment.StartsAt.Before(now) && !appointment.IsAllDay {
				continue
			}

			if appointment.StartsAt.After(tomorrow) && !appointment.StartsAt.After(dayAfterTomorrow) {
				log.Printf("got event for tomorrow %s @ %s\n", appointment.Subject, appointment.StartsAt)
				writeAppointment(outputTomorrow, appointment)
				continue
			}

			if appointment.StartsAt.After(tomorrow) {
				continue
			}

			log.Printf("got event for today %s @ %s\n", appointment.Subject, appointment.StartsAt)
			writeAppointment(outputToday, appointment)
		}
	}

	note.Save(outputToday, filepath.Join(note.GetNotesPath(note.Outlook), "🌞 Today.md"))
	note.Save(outputTomorrow, filepath.Join(note.GetNotesPath(note.Outlook), "🌞 Tomorrow.md"))
	log.Println("outlook completed.")
}

func writeAppointment(output *bytes.Buffer, appointment Appointment) {
	startsAt := appointment.StartsAt.Format("15:04")
	if appointment.IsAllDay {
		startsAt = "All day"
	}

	footer := []string{startsAt}
	if appointment.Location != "" {
		footer = append(footer, appointment.Location)
	}

	if appointment.WithJoinAction {
		footer = append(footer, "[☎️ Open Teams](https://teams.microsoft.com/l/meetup-join)")
	}

	fmt.Fprintf(output, "> %s \n", appointment.Subject)
	fmt.Fprintf(output, "🕰️ %s\n\n", strings.Join(footer, " | "))
}
