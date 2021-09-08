package subprocess

import (
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

/*
Exec passes execution control over to the provided "command".

It looks-up the FQ Path to the command, also combines the given os.Environ (copy) and .env KV maps into a
new os.Environ []string and passes these into the syscall.Exec() function.
*/
func (s *Subprocess) Exec() error {
	binPath, err := s.bin()
	if err != nil {
		return err
	}

	// TODO: Debug log?
	log.Infof("Running command %#v (%v)...", s.command, binPath)
	return syscall.Exec(binPath, s.command, s.env)
}

func (s *Subprocess) bin() (string, error) {
	return exec.LookPath(s.command[0])
}
