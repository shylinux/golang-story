package config

import "runtime"

const (
	SOURCE  = "source"
	LINUX   = "linux"
	DARWIN  = "darwin"
	WINDOWS = "windows"
)

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
		return target
	}
	if target, ok := s.Source[name]; ok {
		target.Type = SOURCE
		return target
	}
	return Target{}
}
