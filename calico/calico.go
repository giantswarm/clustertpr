package calico

type Calico struct {
	CIDR   int    `json:"cidr" yaml:"cidr"`
	MTU    int    `json:"mtu" yaml:"mtu"`
	Subnet string `json:"subnet" yaml:"subnet"`
}
