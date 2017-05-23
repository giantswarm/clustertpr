package kubernetes

import (
	"github.com/giantswarm/clustertpr/kubernetes/api"
	"github.com/giantswarm/clustertpr/kubernetes/dns"
	"github.com/giantswarm/clustertpr/kubernetes/hyperkube"
	"github.com/giantswarm/clustertpr/kubernetes/ingress"
	"github.com/giantswarm/clustertpr/kubernetes/kubelet"
	"github.com/giantswarm/clustertpr/kubernetes/networksetup"
)

type Kubernetes struct {
	API               api.API                   `json:"api" yaml:"api"`
	DNS               dns.DNS                   `json:"dns" yaml:"dns"`
	Hyperkube         hyperkube.Hyperkube       `json:"hyperkube" yaml:"hyperkube"`
	IngressController ingress.IngressController `json:"ingressController" yaml:"ingressController"`
	Kubelet           kubelet.Kubelet           `json:"kubelet" yaml:"kubelet"`
	NetworkSetup      networksetup.NetworkSetup `json:"networkSetup" yaml:"networkSetup"`
}
