package remember_test

import (
	"remember"
	"testing"
)

func TestStoreReminder(t *testing.T) {
	t.Parallel()
	args := []string{"buy milk", "buy chicken"}
	r, err := remember.NewReminder(remember.WithReminders(args))
	if err != nil {
		t.Fatal(err)
	}
	want := 12
	got, err := r.StoreReminder()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestPrintReminders(t *testing.T) {
	t.Parallel()
	args := []string{}
	r, err := remember.NewReminder(remember.WithReminders(args))
	if err != nil {
		t.Fatal(err)
	}
	want := "buy milk\nbuy chicken\n"
	got := r.PrintReminders()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
