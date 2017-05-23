package ingress

type IngressController struct {
	// InsecurePort is the HTTP node port of the Ingress Controller.
	InsecurePort int `json:"insecurePort" yaml:"insecurePort"`
	// SecurePort is the HTTPS node port of the Ingress Controller.
	SecurePort int `json:"securePort" yaml:"securePort"`
}
