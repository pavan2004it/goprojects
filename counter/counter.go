package counter

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Counter struct {
	Count  int
	Output io.Writer
}

func (c *Counter) Next() int {
	c.Count++
	if c.Count == 1 {
		return 0
	}
	c.Count--
	return c.Count

}

func (c *Counter) SetValue(newCounter int) int {
	c.Count = newCounter
	return c.Count
}

func (c *Counter) Run() {
	for {
		c.RunPrint()
	}
}

func (c *Counter) RunPrint() {
	fmt.Fprintln(c.Output, c.Next())
}

func (c *Counter) RunWait() {
	x := 0
	for x < 3 {
		x++
		time.Sleep(60 * time.Second)
		fmt.Fprintln(c.Output, c.Next())

	}
}

func NewRun() *Counter {
	return &Counter{
		Output: os.Stdout,
	}
}

func Run() {
	NewRun().RunWait()
}
