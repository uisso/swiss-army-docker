package library

import (
	"testing"
	"time"
	"github.com/uisso/swiss-army-docker/util"
)

const (
	mongoImage = "library/mongo"
)

// SetupMongoContainer sets up a real MongoDB instance for testing purposes,
// using a Docker container. It returns the container ID and its IP address,
// or makes the test fail on error.
// Currently using https://hub.docker.com/_/mongo/
func SetupMongoContainer(t *testing.T) (c util.ContainerID, ip string) {
	start := func() (string, error) {
		return util.Run("-d", "-p", "27017:27017", "--name", "army_leaf", mongoImage)
	}
	return util.SetupContainer(t, mongoImage, 27017, 10 * time.Second, start)
}