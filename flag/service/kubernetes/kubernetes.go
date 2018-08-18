package kubernetes

import (
	"github.com/giantswarm/node-health/flag/service/kubernetes/tls"
)

type Kubernetes struct {
	Address   string
	InCluster string
	TLS       tls.TLS
}
