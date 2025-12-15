// internal/calendar/service_test.go
package calendar

import "testing"

func TestCreateAndGet(t *testing.T) {
	svc := NewService()

	ev, err := svc.Create(1, "2025-01-01", "нг")
	if err != nil {
		t.Fatal(err)
	}

	res := svc.EventsForPeriod(1, ev.Date, ev.Date)
	if len(res) != 1 {
		t.Fatal("ожидается 1 событие")
	}
}

func TestDelete(t *testing.T) {
	svc := NewService()
	ev, _ := svc.Create(1, "2025-01-01", "test")

	if err := svc.Delete(ev.ID); err != nil {
		t.Fatal(err)
	}

	if err := svc.Delete(ev.ID); err == nil {
		t.Fatal("ожидается ошибка")
	}
}
