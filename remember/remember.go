package remember

import (
	"bufio"
	"fmt"
	"os"
)

type reminder struct {
	reminders []string
}

type option func(*reminder) error

func WithReminders(reminders []string) option {
	return func(r *reminder) error {
		r.reminders = reminders
		return nil
	}
}

func NewReminder(opts ...option) (reminder, error) {
	r := reminder{reminders: os.Args[1:]}
	for _, opt := range opts {
		err := opt(&r)
		if err != nil {
			return reminder{}, err
		}
	}
	return r, nil
}

func (r reminder) StoreReminder() (int, error) {
	bytesWritten := 0
	f, err := os.OpenFile("testdata/reminders.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return 0, err
	}
	for _, data := range r.reminders {
		bytesWritten, err = f.WriteString(data + "\n")
	}
	defer f.Close()
	f.Sync()
	return bytesWritten, nil
}

func (r reminder) PrintReminders() string {
	f, err := os.Open("testdata/reminders.txt")
	if err != nil {
		panic(err)
	}
	data := ""
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		data += scanner.Text() + "\n"
	}
	return data
}

func RunReminders() {
	r, err := NewReminder(WithReminders(os.Args[1:]))
	if err != nil {
		panic(err)
	}
	if len(os.Args) <= 1 {
		fmt.Println(r.PrintReminders())
	} else {
		bytesWritten, err := r.StoreReminder()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d bytes written\n", bytesWritten)
	}

}
