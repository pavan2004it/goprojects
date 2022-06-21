package utility

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"time"
)

type CommandOutput struct {
	FirstCommand  string
	SecondCommand string
}

var firstCommandOutput = regexp.MustCompile("(\\d{5})")
var secondCommandOutput = regexp.MustCompile("utility_test.go")

func ParseCommandOutput(data string) (CommandOutput, error) {
	matchFirst := firstCommandOutput.FindStringSubmatch(data)
	matchSecond := secondCommandOutput.FindStringSubmatch(data)

	if len(matchFirst) < 2 {
		return CommandOutput{}, fmt.Errorf("failed to parse first output")
	}
	if len(matchSecond) < 1 {
		return CommandOutput{}, fmt.Errorf("failed to parse second output")
	}
	firstCommand := matchFirst[1]
	secondCommand := matchSecond[0]
	return CommandOutput{FirstCommand: firstCommand, SecondCommand: secondCommand}, nil
}

func WaitMultipleCommand() (string, error) {
	cmd := exec.Command("/Users/pavankumar/printgofiles.sh")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	otherCmd, err := exec.Command("ls", "-l").Output()
	err = cmd.Wait()
	return outb.String() + ", " + string(otherCmd), nil
}

func UserInteraction() {
	binary, lookErr := exec.LookPath("bc")
	if lookErr != nil {
		panic(lookErr)
	}
	cmd := exec.Command(binary, "-q")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	info, _ := os.Stdout.Stat()
	fmt.Println(info.Name())
}

func CommandTimeout() error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	err := exec.CommandContext(ctx, "sleep", "5").Run()
	if err != nil {
		return err
	}
	return nil
}
