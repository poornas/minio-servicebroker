package main

import (
	"context"
	"fmt"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
	"github.com/satori/go.uuid"
)

type Credentials struct {
	Endpoint  string
	Port      string
	AccessKey string
	SecretKey string
	Secure    bool
}

type BrokerConfig struct {
	Password string
	Username string
}
type InstanceCreator interface {
	Create(instanceID string) error
	Destroy(instanceID string) error
	Exists(instanceID string) (bool, error)
}

type InstanceBinder interface {
	Bind(instanceID string, bindingID string) (Credentials, error)
	Unbind(instanceID string, bindingID string) error
}

// Broker
type MinioServiceBroker struct {
	log lager.Logger
	// Serviceplan Info
	serviceName        string
	serviceID          string
	serviceDescription string
	serviceTags        []string
	bindableService    bool

	// plan-specific customization
	planName         string
	planDescription  string
	planID           string
	bindablePlan     bool
	InstanceCreators map[string]InstanceCreator
	InstanceBinders  map[string]InstanceBinder

	// Broker Config
	Config BrokerConfig
}

// Services Api
func (b *MinioServiceBroker) Services(ctx context.Context) []brokerapi.Service {
	b.log.Info("Building services catalog...")
	brokerID := uuid.NewV4().String()

	return []brokerapi.Service{
		brokerapi.Service{
			ID:            brokerID,
			Name:          b.serviceName,
			Description:   b.serviceDescription,
			Tags:          []string{},
			Bindable:      b.bindableService,
			PlanUpdatable: b.bindablePlan,
			Plans: []brokerapi.ServicePlan{
				brokerapi.ServicePlan{
					ID:          fmt.Sprintf("%s.%s", brokerID, b.planName),
					Name:        b.planName,
					Description: b.planDescription,
					Free:        brokerapi.FreeValue(true),
				},
			},
		},
	}
}

//Provision ...
func (b *MinioServiceBroker) Provision(ctx context.Context, instanceID string, serviceDetails brokerapi.ProvisionDetails, asyncAllowed bool) (spec brokerapi.ProvisionedServiceSpec, err error) {
	b.log.Info("Provisioning new instance ...")
	return brokerapi.ProvisionedServiceSpec{}, nil
}

// Deprovision Api
func (b *MinioServiceBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	b.log.Info("Deprovisioning new instance...")
	return brokerapi.DeprovisionServiceSpec{}, nil
}

// Bind Api
func (b *MinioServiceBroker) Bind(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {

	b.log.Debug("Binding service...", lager.Data{
		"binding-id":  bindingID,
		"instance-id": instanceID,
	})
	return brokerapi.Binding{}, nil
}

// Unbind Api
func (b *MinioServiceBroker) Unbind(ctx context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	b.log.Info("Unbinding service...", lager.Data{
		"binding-id":  bindingID,
		"instance-id": instanceID,
	})
	return nil
}

// LastOperation ...
func (b *MinioServiceBroker) LastOperation(ctx context.Context, instanceID, operationData string) (brokerapi.LastOperation, error) {
	b.log.Info("Last operation", lager.Data{
		"instance-id": instanceID,
	})
	return brokerapi.LastOperation{}, nil
}

// Update implements brokerapi.ServiceBroker
func (b *MinioServiceBroker) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	b.log.Debug("Updating service...", lager.Data{
		"instance-id": instanceID,
	})
	return brokerapi.UpdateServiceSpec{}, nil
}
