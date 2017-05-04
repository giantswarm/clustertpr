package iptables

import (
	"github.com/giantswarm/clustertpr/operator/network/iptables/docker"
)

type IPTables struct {
	Docker docker.Docker
}
