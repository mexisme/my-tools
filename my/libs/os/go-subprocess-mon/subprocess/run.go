package subprocess

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

/*
Run runs the provided "command".

It combines the given os.Environ (copy) and .env KV maps into a
new os.Environ []string and passes these into the exec.Command() function.
*/
func (s *Subprocess) Run() error {
	stderrChan := make(chan string, 1000)
	defer closeChanIfNotClosed(stderrChan)

	bin, args := s.command[0], s.command[1:]

	command := exec.Command(bin, args...)
	command.Env = s.env
	command.Stdin = os.Stdin   // We don't want to capture Stdin
	command.Stdout = os.Stdout // We don't want to capture Stdout

	stderr, err := command.StderrPipe()
	if err != nil {
		return fmt.Errorf("Error creating StderrPipe for Cmd: %v", err)
	}
	defer stderr.Close()

	stderrCapture(stderr, stderrChan)

	log.WithFields(log.Fields{"Command": bin, "Command Args": strings.Join(args, " ")}).Info("Running command...")
	return runThis(command, stderrChan)
}

func closeChanIfNotClosed(channel chan string) {
	select {
	case <-channel:
	default:
		// Close the channel, if not already closed:
		close(channel)
	}
}

func stderrCapture(reader io.Reader, channel chan string) {
	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			channel <- line
			fmt.Fprintf(os.Stderr, "%s\n", line)
		}
	}()
}

func runThis(command *exec.Cmd, stderrChan chan string) error {
	var stderrCapture []string

	if err := command.Start(); err != nil {
		return err
	}

	log.WithFields(log.Fields{"PID": command.Process.Pid}).Info("Process started")

	err := command.Wait()
	if err != nil {
		close(stderrChan)
		for line := range stderrChan {
			stderrCapture = append(stderrCapture, line)
		}
		if len(stderrCapture) < 1 {
			return err
		}

		return fmt.Errorf("%v:\n%s", err, strings.Join(stderrCapture, "\n"))
	}

	return nil
}
