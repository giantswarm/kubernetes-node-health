package service

import (
	"github.com/giantswarm/node-health/flag/service/kubernetes"
)

type Service struct {
	Kubernetes kubernetes.Kubernetes
}
