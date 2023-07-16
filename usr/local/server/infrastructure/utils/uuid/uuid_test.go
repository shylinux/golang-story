package uuid

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	gen := New(1)
	for i := 0; i < 10; i++ {
		fmt.Printf("%x\n", gen.GenID())
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", gen.GenID())
	}
}
func BenchmarkUUID(b *testing.B) {
	gen := New(1)
	for i := 0; i < b.N; i++ {
		gen.GenID()
	}
}
