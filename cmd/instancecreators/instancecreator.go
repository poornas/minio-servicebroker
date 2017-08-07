package instancecreators

import (
	"code.cloudfoundry.org/lager"
)

// Creator creates Instances
type Creator struct {
	logger lager.Logger
	config Config
}

// New constructs a new InstanceCreator
func New(conf Config, logger lager.Logger) *Creator {
	return &Creator{
		logger:  logger,
		backend: mockbackend.New(config),
	}
}

// Create creates an instance
func (c *Creator) Create(instanceID string) error {
	return nil
}

// Destroy destroys the instance
func (c *Creator) Destroy(instanceID string) error {
	return nil
}

//Exists returns true if instance exists
func (c *Creator) Exists(instanceID string) (bool, error) {
	return false, nil
}
