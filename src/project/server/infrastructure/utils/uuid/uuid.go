package uuid

import (
	"sync"
	"time"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
)

var baseTime = int64(0)

func init() {
	t, _ := time.Parse("20060102", "20230102")
	baseTime = t.Unix()
}

type Generate struct {
	workid   int64
	sequence int64
	mu       sync.Mutex
}

func New(config *config.Config) *Generate {
	return &Generate{workid: int64(config.Consul.WorkID)}
}
func (s *Generate) GenID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sequence++
	s.sequence = s.sequence % 10000
	return (time.Now().Unix()-baseTime)*1000000 + s.workid*10000 + s.sequence

}
