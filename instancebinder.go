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
	binds  map[string]*InstanceBinder
}

// New creates a new binder manager
func NewBinderMgr(config utils.Config, logger lager.Logger) (b *BinderMgr) {
	return &BinderMgr{
		logger: logger,
		conf:   config,
		binds:  make(map[string]*InstanceBinder, 5),
	}
}

// Unbind unbinds the binding for a particular instance
func (b *BinderMgr) Unbind(instanceID string, bindingID string) error {
	return nil
}

// Exists returns a bool on whether the instance exists
func (b *BinderMgr) Exists(instanceID string, bindingID string) (bool, error) {
	return false, nil
}

// Bind binds a particular binding to instance.
func (b *BinderMgr) Bind(instanceID string, bindingID string) (interface{}, error) {
	return nil, brokerapi.ErrInstanceDoesNotExist
}
