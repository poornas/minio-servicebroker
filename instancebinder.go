package main

import (
	"code.cloudfoundry.org/lager"
	"github.com/minio/minio-servicebroker/utils"
	"github.com/pivotal-cf/brokerapi"
)

// BinderMgr holds info about the InstanceBinders
type BinderMgr struct {
	logger lager.Logger
	conf   utils.Config
	binds  map[string]*BindingInfo
}

// BindingInfo holds binding state
type BindingInfo struct {
	instanceID string
	bindingID  string
	creds      Credentials
	// other state info
}

// New creates a new binder manager
func NewBinderMgr(config utils.Config, logger lager.Logger) (b *BinderMgr) {
	return &BinderMgr{
		logger: logger,
		conf:   config,
		binds:  make(map[string]*BindingInfo, 5),
	}
}

// Returns bindinginfo if it exists
func (mgr *BinderMgr) getBindingByID(bindingID string) *BindingInfo {
	//check if binding is in the map and return state info.
	// Assuming bindingId is unique across instances.
	if binding, found := mgr.binds[bindingID]; found {
		return binding
	}
	return nil
}

// Unbind unbinds the binding for a particular instance
func (mgr *BinderMgr) Unbind(instanceID string, bindingID string) error {
	if _, found := mgr.binds[bindingID]; found {
		delete(mgr.binds, bindingID)
		return nil
	}
	return brokerapi.ErrBindingDoesNotExist
}

// Exists returns a bool on whether the instance exists
func (mgr *BinderMgr) Exists(instanceID string, bindingID string) (bool, error) {
	for _, binding := range mgr.binds {
		if binding.instanceID == instanceID && binding.bindingID == bindingID {
			return true, nil
		}
	}
	return false, nil
}

// Bind binds a particular binding to instance.
func (mgr *BinderMgr) Bind(instanceID string, bindingID string) (interface{}, error) {
	credentials := Credentials{
		instanceID: instanceID,
		bindingID:  bindingID,
		// 	"EndpointURL": instanceCredentials.EndpointURL
		// 	"AccessKey":   instanceCredentials.AccessKey,
		// 	"SecretKey":   instanceCredentials.SecretKey,
	}
	bind := &BindingInfo{instanceID: instanceID,
		bindingID: bindingID,
		creds:     credentials}
	mgr.binds[bindingID] = bind

	return bind, nil
}
