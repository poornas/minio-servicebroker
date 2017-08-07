package main

import (
	"code.cloudfoundry.org/lager"
	"github.com/minio/minio-servicebroker/utils"
)

// InstanceMgr holds instances info
type InstanceMgr struct {
	logger    lager.Logger
	conf      utils.Config
	instances map[string]*InstanceCreator
}

// NewInstanceMgr manages running instances
func NewInstanceMgr(config utils.Config, logger lager.Logger) (m *InstanceMgr) {
	return &InstanceMgr{
		logger:    logger,
		conf:      config,
		instances: make(map[string]*InstanceCreator, 10),
	}
}

// Returns instance if it exists
func (mgr *InstanceMgr) getInstanceByID(instanceID string) *InstanceCreator {
	//check if isntance is in the map and return it.
	if instance, found := mgr.instances[instanceID]; found {
		return instance
	}
	return nil
}

// Create creates an instance
func (mgr *InstanceMgr) Create(instanceID string) error {
	//TODO create instance here
	mgr.instances[instanceID] = InstanceCreator{} //hold reference to provisioned instance state
	return nil
}

// Destroy destroys the instance
func (mgr *InstanceMgr) Destroy(instanceID string) error {
	return nil
}

//Exists returns true if instance exists
func (mgr *InstanceMgr) Exists(instanceID string) (bool, error) {
	return false, nil
}
