package cluster

type Cluster struct {
	ID string `json:"id" yaml:"id"`
	// Domain is the base domain for records in the cluster, e.g.
	// <cluster-id>.g8s.fra-1.giantswarm.io.
	Domain string `json:"domain" yaml:"domain"`
}
