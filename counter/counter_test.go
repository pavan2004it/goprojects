package counter_test

import (
	"bytes"
	"counter"
	"testing"
)

func TestCounterInitial(t *testing.T) {
	t.Parallel()
	c := counter.Counter{}
	want := 0
	got := c.Next()
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestCounterMultiple(t *testing.T) {
	t.Parallel()
	c := counter.Counter{}
	want := 1
	c.Next()
	got := c.Next()
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestSetCounterValue(t *testing.T) {
	c := counter.Counter{}
	want := 10
	got := c.SetValue(want)
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	p := &counter.Counter{Output: fakeTerminal}
	p.RunWait()
	want := "1\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}

//func TestWaitRun(t *testing.T) {
//	t.Parallel()
//	fakeTerminal := &bytes.Buffer{}
//	p := &counter.Counter{Output: fakeTerminal}
//	p.RunWait()
//	want := "1\n" +
//		"2\n" +
//		"3\n"
//	got := fakeTerminal.String()
//	if want != got {
//		t.Errorf("want %q, got %q", want, got)
//	}
//
//}
