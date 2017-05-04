package endpointupdater

import (
	"github.com/giantswarm/clustertpr/operator/endpointupdater/docker"
)

type EndpointUpdater struct {
	Docker docker.Docker
}
