package subprocess

type watchLog struct {
	cpu    bool
	memory bool
}

// Subprocess is simply the struct method-wrapper for the "exec" package
type Subprocess struct {
	env      []string
	command  []string
	watchLog watchLog
}

// New creates a new exec.Subprocess struct
func New() *Subprocess {
	new := &Subprocess{}
	new.watchLog = watchLog{cpu: false, memory: false}

	return new
}

// SetEnviron creates a new subprocess.Subprocess struct with the os.Environ object copied-in
func (s *Subprocess) SetEnviron(env []string) *Subprocess {
	s.env = env
	return s
}

// SetCommand creates a new subprocess.Subprocess struct with the command []string copied-in
func (s *Subprocess) SetCommand(command []string) *Subprocess {
	s.command = command
	return s
}

// SetWatchLog creates a new subprocess.Subprocess struct with CPU or Memory watch loggers enabled
func (s *Subprocess) SetWatchLog(cpu, memory bool) *Subprocess {
	s.watchLog.cpu = cpu
	s.watchLog.memory = memory

	return s
}
