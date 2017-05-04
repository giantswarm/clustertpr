package openssl

import (
	"github.com/giantswarm/clustertpr/operator/network/openssl/docker"
)

type OpenSSL struct {
	Docker docker.Docker
}
