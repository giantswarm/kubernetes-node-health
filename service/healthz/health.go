package healthz

import (
	"context"

	"github.com/giantswarm/microendpoint/service/healthz"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

const (
	// Description describes which functionality this health check implements.
	Description = "Ensure node availability."
	// Name is the identifier of the health check. This can be used for emitting
	// metrics.
	Name = "node"
	// SuccessMessage is the message returned in case the health check did not
	// fail.
	SuccessMessage = "all good"
)

// Config represents the configuration used to create a healthz service.
type Config struct {
	Logger micrologger.Logger
}

// Healthz implements the health service.
type Healthz struct {
	logger micrologger.Logger
}

// New creates a new configured healthz service.
func New(config Config) (*Healthz, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	h := &Healthz{
		logger: config.Logger,
	}

	return h, nil
}

func (h *Healthz) GetHealthz(ctx context.Context) (healthz.Response, error) {
	r := healthz.Response{
		Description: Description,
		Failed:      false,
		Message:     SuccessMessage,
		Name:        Name,
	}

	return r, nil
}
