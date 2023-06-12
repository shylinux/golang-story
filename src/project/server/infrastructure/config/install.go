package config

import "runtime"

type Install struct {
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
	case "linux":
		return s.Linux[name]
	case "darwin":
		return s.Darwin[name]
	case "windows":
		return s.Windows[name]
	default:
		return Target{}
	}
}
