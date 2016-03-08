package library

import (
	"testing"
	"github.com/uisso/swiss-army-docker/util"
	"time"
)

const (
	rabbitImage = "registry-aws.neoway.com.br/prevention/rabbitmq:3.5.6"
)

// SetupRabbitContainer sets up a real Rabbit instance for testing purposes,
// using a Docker container. It returns the container ID and its IP address,
// or makes the test fail on error.
func SetupRabbitContainer(t *testing.T) (c util.ContainerID, ip string) {
	start := func() (string, error) {
		return util.Run("-d", "-p", "5672:5672", "--name", "army_bugsy", rabbitImage)
	}
	return util.SetupContainer(t, rabbitImage, 5672, 10 * time.Second, start)
}
