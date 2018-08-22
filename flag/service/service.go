package service

import (
	"github.com/giantswarm/kubernetes-node-health/flag/service/kubernetes"
)

type Service struct {
	Kubernetes kubernetes.Kubernetes
}
