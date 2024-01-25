package outlook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func Sync(home string, path string) {
	log.Println("loading outlook events...")

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

	save(outputToday, fmt.Sprintf("%s/%s/🌞 Today.md", home, path))
	save(outputTomorrow, fmt.Sprintf("%s/%s/🌞 Tomorrow.md", home, path))

	log.Println("outlook completed.")
}

// save the buffer to a file (including a timestamp)
func save(output *bytes.Buffer, path string) {
	fmt.Fprintf(output, "_last updated @ %s_ \n\n", time.Now().Format("2006-01-02 15:04"))

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Println("got error", err)
		return
	}

	if _, err := io.Copy(file, output); err != nil {
		log.Println("got error", err)
		return
	}
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
