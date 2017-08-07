package instancebinders

import (
	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
)

// Binder holds info about the Binder manager
type Binder struct {
	logger lager.Logger
	config Config
}

// New creates a new binder manager
func New(config Config, logger lager.Logger) *Binder {
	return &Binder{
		logger:  logger,
		backend: mockbackend.New(config),
	}
}

// Unbind unbinds the binding for a particular instance
func (b *Binder) Unbind(instanceID string, bindingID string) error {
	return nil
}

// Exists returns a bool on whether the instance exists
func (b *Binder) Exists(instanceID string) (bool, error) {
	return false, nil
}

// Bind binds a particular binding to instance.
func (b *Binder) Bind(instanceID string, bindingID string) (interface{}, error) {
	return nil, brokerapi.ErrInstanceDoesNotExist
}
