package environment

import (
	"github.com/giantswarm/clustertpr/operator/network/environment/docker"
)

type Environment struct {
	Docker docker.Docker
}
