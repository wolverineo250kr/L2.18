// cmd/server/main.go
package main

import (
	"log"
	"net/http"
	"os"

	"calendar/internal/calendar"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	svc := calendar.NewService()

	mux := http.NewServeMux()

	mux.Handle("/create_event", logMiddleware(createEvent(svc)))
	mux.Handle("/update_event", logMiddleware(updateEvent(svc)))
	mux.Handle("/delete_event", logMiddleware(deleteEvent(svc)))
	mux.Handle("/events_for_day", logMiddleware(eventsForDay(svc)))
	mux.Handle("/events_for_week", logMiddleware(eventsForWeek(svc)))
	mux.Handle("/events_for_month", logMiddleware(eventsForMonth(svc)))

	log.Println("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
