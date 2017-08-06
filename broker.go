package main

import (
	"context"
	"errors"
	"fmt"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
	uuid "github.com/satori/go.uuid"
)

type Credentials struct {
	EndpointURL string
	AccessKey   string
	SecretKey   string
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
	Exists(instanceID string) (bool, error)
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
	spec = brokerapi.ProvisionedServiceSpec{IsAsync: false}
	if b.instanceExists(instanceID) {
		return spec, errors.New("Instance already exists")
	}
	if serviceDetails.ServiceID != b.serviceID {
		return spec, fmt.Errorf("Service %s does not exist", serviceDetails.ServiceID)
	}

	if serviceDetails.PlanID == "" {
		return spec, errors.New("planId required")
	}
	// Only default plan for now
	if serviceDetails.PlanID != b.planID {
		return spec, errors.New("plan id not recognized")
	}

	instanceCreator, ok := b.InstanceCreators[b.planID]
	if !ok {
		return spec, errors.New("instance creator not found for plan")
	}
	err = instanceCreator.Create(instanceID)
	return brokerapi.ProvisionedServiceSpec{}, err
}

// Deprovision Api
func (b *MinioServiceBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	b.log.Info("Deprovisioning new instance...")
	spec := brokerapi.DeprovisionServiceSpec{}

	for _, instanceCreator := range b.InstanceCreators {
		exists, _ := instanceCreator.Exists(instanceID)
		if exists {
			return spec, instanceCreator.Destroy(instanceID)
		}
	}
	return spec, brokerapi.ErrInstanceDoesNotExist

}

// Bind Api
func (b *MinioServiceBroker) Bind(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {

	b.log.Debug("Binding service...", lager.Data{
		"binding-id":  bindingID,
		"instance-id": instanceID,
	})
	binding := brokerapi.Binding{}

	for _, bindingCreator := range b.InstanceBinders {
		exists, _ := bindingCreator.Exists(instanceID)
		if exists {
			instanceCredentials, err := bindingCreator.Bind(instanceID, bindingID)
			if err != nil {
				return binding, err
			}
			credentialsMap := map[string]interface{}{
				"EndpointURL": instanceCredentials.EndpointURL,
				"AccessKey":   instanceCredentials.AccessKey,
				"SecretKey":   instanceCredentials.SecretKey,
			}

			binding.Credentials = credentialsMap
			return binding, nil
		}
	}
	return brokerapi.Binding{}, brokerapi.ErrInstanceDoesNotExist
}

// Unbind Api
func (b *MinioServiceBroker) Unbind(ctx context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	b.log.Info("Unbinding service...", lager.Data{
		"binding-id":  bindingID,
		"instance-id": instanceID,
	})
	for _, repo := range b.InstanceBinders {
		instanceExists, _ := repo.Exists(instanceID)
		if instanceExists {
			err := repo.Unbind(instanceID, bindingID)
			if err != nil {
				return brokerapi.ErrBindingDoesNotExist
			}
			return nil
		}
	}
	return brokerapi.ErrInstanceDoesNotExist
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

func (b *MinioServiceBroker) instanceExists(instanceID string) bool {
	for _, instanceCreator := range b.InstanceCreators {
		exists, _ := instanceCreator.Exists(instanceID)
		if exists {
			return true
		}
	}
	return false
}
