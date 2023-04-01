package tools

import (
	"fmt"
	"sync"
	"testing"
)

func TestSnowflake_NextVal(t *testing.T) {
	g := sync.WaitGroup{}
	g.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(Snow.NextVal())
			g.Done()
		}()
	}
	g.Wait()
}
