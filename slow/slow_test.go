package slow_test

import (
	"fmt"
	"slow"
	"testing"
	"time"
)

//func TestPrintSlow(t *testing.T) {
//	t.Parallel()
//	data := bytes.NewBufferString("1")
//	input := []io.Reader{data}
//	s, err := slow.NewSlowData(slow.WithInput(input))
//	if err != nil {
//		t.Fatal(err)
//	}
//	want := 3
//	start := time.Now()
//	err = s.PrintSlow()
//	got := int(time.Since(start).Seconds())
//	if want != got {
//		t.Errorf("Mismatch in time want %d got %d", want, got)
//	}
//}

func TestPrintSlowWithArgs(t *testing.T) {
	t.Parallel()
	args := []string{"hello"}
	s, err := slow.NewSlowData(slow.WithArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	want := "h"
	timeoutChan := make(chan bool, 1)
	got := ""
	go func() {
		err, got = s.PrintSlowWithArgs()
		fmt.Println(got)
		timeoutChan <- false
	}()
	go func() {
		time.Sleep(3 * time.Second)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		t.Log("Test Succeeded")
	}

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
