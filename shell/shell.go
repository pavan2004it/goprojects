package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Session struct {
	Stdin          io.Reader
	Stdout, Stderr io.Writer
	DryRun         bool
}

func NewSession(stdin io.Reader, stdout, stderr io.Writer, dryrun bool) *Session {
	return &Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
		DryRun: dryrun,
	}
}

func (s *Session) Run() {
	input := bufio.NewReader(s.Stdin)
	for {
		fmt.Fprintf(s.Stdout, "> ")
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprintln(s.Stdout, "\nBe seeing you!")
			break
		}
		cmd, err := CmdFromString(line)

		if err != nil {
			continue
		}

		if s.DryRun {
			fmt.Fprintf(s.Stdout, "%s", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "error:", err)
		}
		fmt.Fprintf(s.Stdout, "%s", output)
	}
}

func CmdFromString(input string) (*exec.Cmd, error) {
	args := strings.Fields(input)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}

	return exec.Command(args[0], args[1:]...), nil
}

func (s *Session) RunWithPrompt(f func() string) {
	inputPrompt := bufio.NewReader(s.Stdin)
	for {

		fmt.Fprintf(s.Stdout, f())
		line, err := inputPrompt.ReadString('\n')
		if err != nil {
			fmt.Fprintln(s.Stdout, "\nBe seeing you!")
			break
		}
		cmd, err := CmdFromString(line)
		if err != nil {
			continue
		}

		if s.DryRun {
			fmt.Fprintf(s.Stdout, "%s", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "error:", err)
		}
		fmt.Fprintf(s.Stdout, "%s", output)
	}
}

func GeneratePrompt() string {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(out), "\n") + " "
}

func setCommandInternals(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}

func (s *Session) RunInteractively(f func() string) {
	input := bufio.NewReader(s.Stdin)
	for {
		fmt.Fprintf(s.Stdout, f())
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprintln(s.Stdout, "\nBe seeing you!")
			break
		}
		cmd, err := CmdFromString(line)
		setCommandInternals(cmd)

		if err != nil {
			continue
		}

		if s.DryRun {
			_, _ = fmt.Fprintf(s.Stdout, "%s", line)
			continue
		}
		err = cmd.Run()
		if err != nil {
			_, _ = fmt.Fprintln(s.Stderr, "error:", err)
		}

	}
}

func expansion(input string) (error, string) {
	args := strings.Fields(input)
	if len(args) < 1 {
		return errors.New("empty input"), ""
	}
	if len(args) == 1 {
		return nil, args[0] + "\n"
	}
	if args[1] == "*" {
		files, _ := filepath.Glob("." + "/" + args[1])
		for _, file := range files {
			args[0] += " " + file
		}
	} else {
		for _, arg := range args[1:] {
			args[0] += " " + arg
		}
	}
	return nil, args[0] + "\n"
}

func (s *Session) RunWithExpansion() {
	input := bufio.NewReader(s.Stdin)
	for {
		fmt.Fprintf(s.Stdout, "> ")
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprintln(s.Stdout, "\nBe seeing you!")
			break
		}
		cmd, err := CmdFromString(line)

		if err != nil {
			continue
		}

		if s.DryRun {
			_, output := expansion(line)
			fmt.Fprintf(s.Stdout, "%s", output)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "error:", err)
		}
		fmt.Fprintf(s.Stdout, "%s", output)
	}
}

func RunCLI() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr, false)
	session.Run()
}

func RunCliWithPrompt(f func() string) {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr, false)
	session.RunWithPrompt(f)
}

func RunCliWithInteractive(f func() string) {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr, false)
	session.RunInteractively(f)
}

func RunCliWithExpansion() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr, false)
	session.RunWithExpansion()
}
