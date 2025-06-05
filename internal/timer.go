package timer

import (
	"fmt"
	"time"
)

func StartTimer() (stop func()) {
	start := time.Now()
	done := make(chan struct{})

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				elapsed := time.Since(start).Seconds()
				fmt.Printf("\rWaiting for response… %.1fs", elapsed)
			case <-done:
				return
			}
		}
	}()

	return func() {
		close(done)
		fmt.Println()
	}
}
