package uuid

import (
	"sync"
	"time"
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

func New(workid int) *Generate {
	return &Generate{workid: int64(workid)}
}
func (s *Generate) GenID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sequence++
	s.sequence = s.sequence % 10000
	return (time.Now().Unix()-baseTime)*1000000 + s.workid*10000 + s.sequence

}
