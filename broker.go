package main

import (
	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
)

// MinioServiceBroker
type MinioServiceBroker struct {
	log lager.Logger
}

// Services Api
func (b *MinioServiceBroker) Services() []brokerapi.Service {
	b.log.Info("Building services catalog...")
	return nil
}

//Provision ...
func (b *MinioServiceBroker) Provision(instanceID string, serviceDetails brokerapi.ProvisionDetails, asyncAllowed bool) (spec brokerapi.ProvisionedServiceSpec, err error) {
	b.log.Info("Provisioning new instance ...")
	return brokerapi.ProvisionedServiceSpec{}, nil
}

// Deprovision Api
func (b *MinioServiceBroker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	b.log.Info("Deprovisioning new instance...")
	return brokerapi.DeprovisionServiceSpec{}, nil
}

// Bind Api
func (b *MinioServiceBroker) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	b.log.Debug("Binding service...", lager.Data{
		"binding-id":  bindingID,
		"instance-id": instanceID,
	})
	return brokerapi.Binding{}, nil
}

// Unbind Api
func (b *MinioServiceBroker) Unbind(instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	b.log.Info("Unbinding service...", lager.Data{
		"binding-id":  bindingID,
		"instance-id": instanceID,
	})
	return nil
}

// LastOperation ...
func (b *MinioServiceBroker) LastOperation(instanceID, operationData string) (brokerapi.LastOperation, error) {
	b.log.Info("Last operation", lager.Data{
		"instance-id": instanceID,
	})
	return brokerapi.LastOperation{}, nil
}

// Update implements brokerapi.ServiceBroker
func (b *MinioServiceBroker) Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	b.log.Debug("Updating service...", lager.Data{
		"instance-id": instanceID,
	})
	return brokerapi.UpdateServiceSpec{}, nil
}
