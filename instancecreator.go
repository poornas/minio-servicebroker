package main

import (
	"code.cloudfoundry.org/lager"
	"github.com/minio/minio-servicebroker/client"
	"github.com/minio/minio-servicebroker/utils"
	"github.com/pivotal-cf/brokerapi"
)

// InstanceMgr holds instances info
type InstanceMgr struct {
	logger    lager.Logger
	conf      utils.Config
	instances map[string]*InstanceInfo
	client    *client.ApiClient
}

// InstanceInfo holds instance state
type InstanceInfo struct {
	instanceID string
	// other state info
}

// NewInstanceMgr manages running instances
func NewInstanceMgr(config utils.Config, logger lager.Logger) *InstanceMgr {
	c, err := client.New(config, logger)
	if err != nil {
		return nil
	}
	return &InstanceMgr{
		logger:    logger,
		conf:      config,
		instances: make(map[string]*InstanceInfo, 10),
		client:    c,
	}
}

// Returns instance if it exists
func (mgr *InstanceMgr) getInstanceByID(instanceID string) *InstanceInfo {
	//check if isntance is in the map and return it.
	if instance, found := mgr.instances[instanceID]; found {
		return instance
	}
	return nil
}

// Create creates an instance
func (mgr *InstanceMgr) Create(instanceID string) error {
	//TODO create instance here
	settings := map[string]string{
		"instanceID": instanceID,
	}
	instance, err := mgr.client.CreateInstance(settings)
	if err != nil {
		return err
	}
	mgr.instances[instanceID] = &InstanceInfo{instanceID: instance} //hold reference to provisioned instance state

	return nil
}

// Destroy destroys the instance
func (mgr *InstanceMgr) Destroy(instanceID string) error {
	found, _ := mgr.Exists(instanceID)
	if found {
		delete(mgr.instances, instanceID)
		err := mgr.client.DeleteInstance(instanceID)
		if err != nil {
			return err
		}
	}
	return brokerapi.ErrInstanceDoesNotExist
}

//Exists returns true if instance exists
func (mgr *InstanceMgr) Exists(instanceID string) (bool, error) {
	for _, instance := range mgr.instances {
		if instance.instanceID == instanceID {
			return true, nil
		}
	}
	return false, nil
}
