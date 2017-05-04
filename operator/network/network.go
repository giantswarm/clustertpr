package network

import (
	"github.com/giantswarm/clustertpr/operator/network/bridge"
	"github.com/giantswarm/clustertpr/operator/network/config"
	"github.com/giantswarm/clustertpr/operator/network/environment"
	"github.com/giantswarm/clustertpr/operator/network/iptables"
	"github.com/giantswarm/clustertpr/operator/network/openssl"
)

type Network struct {
	Bridge      bridge.Bridge
	Config      config.Config
	Environment environment.Environment
	IPTables    iptables.IPTables
	OpenSSL     openssl.OpenSSL
}
