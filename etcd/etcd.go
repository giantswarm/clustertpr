package etcd

type Etcd struct {
	// AltNames is the alternative names used to generate certificates for Etcd.
	AltNames string `json:"altNames" yaml:"altNames"`
	Port     int    `json:"port" yaml:"port"`
	Prefix   string `json:"prefix" yaml:"prefix"`
}
