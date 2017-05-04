package operator

import (
	"github.com/giantswarm/clustertpr/operator/certctl"
	"github.com/giantswarm/clustertpr/operator/endpointupdater"
	"github.com/giantswarm/clustertpr/operator/k8svm"
	"github.com/giantswarm/clustertpr/operator/kubectl"
	"github.com/giantswarm/clustertpr/operator/network"
)

type Operator struct {
	Certctl         certctl.Certctl                 `json:"certctl" yaml:"certctl"`
	EndpointUpdater endpointupdater.EndpointUpdater `json:"endpointUpdater" yaml:"endpointUpdater"`
	K8sVM           k8svm.K8sVM                     `json:"k8sVM" yaml:"k8sVM"`
	Kubectl         kubectl.Kubectl                 `json:"kubectl" yaml:"kubectl"`
	Network         network.Network                 `json:"network" yaml:"network"`
}
