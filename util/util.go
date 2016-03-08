package util

import (
	"time"
	"net"
	"fmt"
)

// AwaitReachable tries to make a TCP connection to addr regularly.
// It returns an error if it's unable to make a connection before maxWait.
func awaitReachable(addr string, maxWait time.Duration) error {
	done := time.Now().Add(maxWait)
	for time.Now().Before(done) {
		c, err := net.Dial("tcp", addr)
		if err == nil {
			c.Close()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("%v unreachable for %v", addr, maxWait)
}
