package config

import "runtime"

const (
	SOURCE  = "source"
	BINARY  = "binary"
	LINUX   = "linux"
	DARWIN  = "darwin"
	WINDOWS = "windows"
)

type Replace struct {
	From string
	To   string
}
type Install struct {
	Source  map[string]Target
	Binary  map[string]Target
	Linux   map[string]Target
	Darwin  map[string]Target
	Windows map[string]Target
}
type Target struct {
	Type    string
	Name    string
	Address string
	Unpack  string
	Install string
	Plugin  []string
	Build   []string
	Start   string
	Daemon  bool
	Export  bool
}

func (s Install) GetTarget(name string) Target {
	switch runtime.GOOS {
	case LINUX:
		if target, ok := s.Linux[name]; ok {
			return target
		}
	case DARWIN:
		if target, ok := s.Darwin[name]; ok {
			return target
		}
	case WINDOWS:
		if target, ok := s.Windows[name]; ok {
			return target
		}
	}
	if target, ok := s.Binary[name]; ok {
		target.Type = BINARY
		return target
	}
	if target, ok := s.Source[name]; ok {
		target.Type = SOURCE
		return target
	}
	return Target{}
}
func (s Install) ForEach(cb func(string, Target)) {
	each := func(list map[string]Target) {
		for k, v := range list {
			cb(k, v)
		}
	}
	switch runtime.GOOS {
	case LINUX:
		each(s.Linux)
	case DARWIN:
		each(s.Darwin)
	case WINDOWS:
		each(s.Windows)
	}
	each(s.Binary)
	each(s.Source)
}
