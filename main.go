package main

import (
	"fmt"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
)

const (
	// DefaultServiceName is the name of Minio service on the marketplace
	DefaultServiceName = "Minio"

	// DefaultServiceDescription is the description of the default service
	DefaultServiceDescription = "Minio Service Broker"

	// DefaultPlanName is the name of our supported plan
	DefaultPlanName = "default"

	//DefaultPlanDescription describes the default plan offered.
	DefaultPlanDescription = "Secure access to a single instance Minio server"

	// DefaultServiceID is placeholder id for the service broker
	DefaultServiceID = "966fa3f8-c666-461e-acfe-bfae50bb46ad"
)

func main() {
	// Create logger
	log := lager.NewLogger("minio-servicebroker")
	log.RegisterSink(lager.NewWriterSink(os.Stderr, lager.DEBUG))
	log.RegisterSink(lager.NewWriterSink(os.Stderr, lager.INFO))

	// Ensure username and password are present
	username := os.Getenv("USER_NAME")
	if username == "" {
		log.Fatal("missing USER_NAME", nil)
	}
	password := os.Getenv("USER_PASSWORD")
	if password == "" {
		log.Fatal("missing USER_PASSWORD", nil)
	}
	credentials := brokerapi.BrokerCredentials{
		Username: username,
		Password: password,
	}

	// Setup the broker
	broker := &MinioServiceBroker{
		log:                log,
		serviceID:          DefaultServiceID,
		serviceName:        DefaultServiceName,
		serviceDescription: DefaultServiceDescription,
		bindableService:    true,
		planName:           DefaultPlanName,
		planDescription:    DefaultPlanDescription,
		bindablePlan:       true,
	}

	brokerAPI := brokerapi.New(broker, log, credentials)
	http.Handle("/", brokerAPI)
	log.Info("Listening for requests")
	err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
	if err != nil {
		log.Error("Failed to start the server", err)
	}

}
