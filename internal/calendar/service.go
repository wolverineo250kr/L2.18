// internal/calendar/service.go
package calendar

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrNotFound   = errors.New("событие не найдено")
	ErrBadDate    = errors.New("ложная дата")
)

type Service struct {
	mu     sync.RWMutex
	events map[int]Event
	nextID int
}

func NewService() *Service {
	return &Service{
		events: make(map[int]Event),
		nextID: 1,
	}
}

func parseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func (s *Service) Create(userID int, date, text string) (Event, error) {
	d, err := parseDate(date)
	if err != nil {
		return Event{}, ErrBadDate
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	ev := Event{
		ID:     s.nextID,
		UserID: userID,
		Date:   d,
		Text:   text,
	}

	s.events[s.nextID] = ev
	s.nextID++

	return ev, nil
}

func (s *Service) Update(id int, text string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	ev, ok := s.events[id]
	if !ok {
		return ErrNotFound
	}

	ev.Text = text
	s.events[id] = ev
	return nil
}

func (s *Service) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[id]; !ok {
		return ErrNotFound
	}

	delete(s.events, id)
	return nil
}

func (s *Service) EventsForPeriod(userID int, from, to time.Time) []Event {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res []Event
	for _, ev := range s.events {
		if ev.UserID == userID &&
			!ev.Date.Before(from) &&
			!ev.Date.After(to) {
			res = append(res, ev)
		}
	}
	return res
}
