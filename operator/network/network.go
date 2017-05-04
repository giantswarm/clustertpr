package network

import (
	"github.com/giantswarm/clustertpr/operator/network/environment"
	"github.com/giantswarm/clustertpr/operator/network/iptables"
	"github.com/giantswarm/clustertpr/operator/network/openssl"
)

type Network struct {
	Environment environment.Environment
	IPTables    iptables.IPTables
	OpenSSL     openssl.OpenSSL
}
