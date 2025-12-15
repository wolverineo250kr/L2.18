package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"calendar/internal/calendar"
)

func createEvent(s *calendar.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID int    `json:"user_id"`
			Date   string `json:"date"`
			Text   string `json:"event"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"ошибка":"зачем мне шлеш плохой json?"}`, http.StatusBadRequest)
			return
		}

		ev, err := s.Create(req.UserID, req.Date, req.Text)
		if err != nil {
			http.Error(w, `{"ошибка":"`+err.Error()+`"}`, http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(map[string]any{"result": ev})
	})
}

func updateEvent(s *calendar.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID   int    `json:"id"`
			Text string `json:"event"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"ошибка":"bad request"}`, http.StatusBadRequest)
			return
		}

		if err := s.Update(req.ID, req.Text); err != nil {
			http.Error(w, `{"ошибка":"`+err.Error()+`"}`, http.StatusServiceUnavailable)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"result": "updated"})
	})
}

func deleteEvent(s *calendar.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID int `json:"id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"bad request"}`, http.StatusBadRequest)
			return
		}

		if err := s.Delete(req.ID); err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusServiceUnavailable)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"result": "deleted"})
	})
}

func eventsForDay(s *calendar.Service) http.Handler {
	return periodHandler(s, 1)
}

func eventsForWeek(s *calendar.Service) http.Handler {
	return periodHandler(s, 7)
}

func eventsForMonth(s *calendar.Service) http.Handler {
	return periodHandler(s, 30)
}

func periodHandler(s *calendar.Service, days int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		dateStr := r.URL.Query().Get("date")

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, `{"ошибка":"неверная дата"}`, http.StatusBadRequest)
			return
		}

		events := s.EventsForPeriod(
			userID,
			date,
			date.AddDate(0, 0, days),
		)

		json.NewEncoder(w).Encode(map[string]any{"result": events})
	})
}
