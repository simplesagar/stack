package internal

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/egymgmbh/go-prefix-writer/prefixer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ory/dockertest/v3"
)

var (
	benthosResource *dockertest.Resource
)

func startBenthosServer() {
	entrypoint := []string{
		"/benthos",
		"-c", "/config/config.yml",
		"-t", "/config/templates/*.yaml",
		"-r", "/config/resources/*.yaml",
	}
	if testing.Verbose() {
		entrypoint = append(entrypoint, "--log.level", "trace")
	}
	entrypoint = append(entrypoint, "streams", "/config/streams/*.yaml")
	wd, err := os.Getwd()

	host := os.Getenv("DOCKER_HOSTNAME")
	if host == "" {
		host = "host.docker.internal"
	}

	Expect(err).To(BeNil())
	benthosResource = runDockerResource(&dockertest.RunOptions{
		Repository: "jeffail/benthos",
		Tag:        "4.11",
		Mounts: []string{
			fmt.Sprintf("%s/../../../components/search/benthos:/config", wd),
		},
		Tty:        true,
		Entrypoint: entrypoint,
		Env: []string{
			fmt.Sprintf("OPENSEARCH_URL=http://%s:9200", host), // TODO: Make configurable
			"BASIC_AUTH_ENABLED=true",
			"BASIC_AUTH_USERNAME=admin",
			"BASIC_AUTH_PASSWORD=admin",
			fmt.Sprintf("OPENSEARCH_INDEX=%s", actualTestID),
			fmt.Sprintf("NATS_URL=nats://%s:%s", host, natsPort()),
			fmt.Sprintf("TOPIC_PREFIX=%s-", actualTestID),
		},
	})

	go func() {
		defer GinkgoRecover()
		reader, err := dockerClient.ContainerLogs(TestContext(), benthosResource.Container.ID, types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     true,
			Details:    false,
		})
		Expect(err).To(BeNil())

		io.Copy(prefixer.New(GinkgoWriter, func() string {
			return "benthos | "
		}), reader)
	}()
}

func stopBenthosServer() {
	Expect(benthosResource.Close()).Should(BeNil())
}
