//go:build !windows

package main

import (
	"bytes"
	"io"
	"time"

	"github.com/creack/pty"
)

func (s *Snapshot) run(finishedQueue chan<- int64, width int, isPty bool) error {
	s.start = time.Now()
	defer func() {
		s.end = time.Now()
	}()

	var b, eb bytes.Buffer

	commands := []string{s.command}
	commands = append(commands, s.args...)

	command := s.prepareCommand(commands)
	command.Stderr = &eb

	if isPty {
		// G115: width is int, but pty.Winsize.Cols is uint16.
		// In this case, width is from terminal size, which is uint16.
		// So, this conversion is safe.
		//nolint:gosec
		pty, err := pty.StartWithSize(command, &pty.Winsize{
			Cols: uint16(width),
		})
		if err != nil {
			return err
		}

		go func() {
			_, _ = io.Copy(&b, pty)
		}()
	} else {
		command.Stdout = &b
		if err := command.Start(); err != nil {
			return err
		}
	}

	go func() {
		if err := command.Wait(); err != nil {
			s.err = err
		}

		s.result = b.Bytes()
		s.errorResult = eb.Bytes()
		s.exitCode = command.ProcessState.ExitCode()
		s.completed = true
		finishedQueue <- s.id
		close(s.finish)
	}()

	return nil
}
