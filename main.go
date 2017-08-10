package main

import (
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"

	"github.com/pivotal-cf/brokerapi"
	"github.com/poornas/minio-servicebroker/utils"
)

const (
	// DefaultServiceName is the name of Minio service on the marketplace
	DefaultServiceName = "Minio"

	// DefaultServiceDescription is the description of the default service
	DefaultServiceDescription = "Minio Service Broker"

	// DefaultPlanName is the name of our supported plan
	DefaultPlanName = "default"
	// DefaultPlanID is the ID of our supported plan
	DefaultPlanID = "1234"
	//DefaultPlanDescription describes the default plan offered.
	DefaultPlanDescription = "Secure access to a single instance Minio server"

	// DefaultServiceID is placeholder id for the service broker
	DefaultServiceID = "minio-broker-id"
)

// this is just a stub - #TODO load any config from file
func getConfig() (conf utils.Config) {
	conf = utils.Config{
		Endpoint:  "play.minio.io:9000",
		AccessKey: "minio",
		SecretKey: "minio123",
		Secure:    true,
	}
	return
}

func main() {
	// Create logger
	log := lager.NewLogger("minio-servicebroker")
	log.RegisterSink(lager.NewWriterSink(os.Stderr, lager.DEBUG))
	log.RegisterSink(lager.NewWriterSink(os.Stderr, lager.INFO))

	// Ensure username and password are present
	username := os.Getenv("SECURITY_USER_NAME")
	if username == "" {
		username = "miniobroker"
	}
	password := os.Getenv("SECURITY_USER_PASSWORD")
	if password == "" {
		password = "miniobroker123"
	}
	credentials := brokerapi.BrokerCredentials{
		Username: username,
		Password: password,
	}
	// Load endpoint config
	conf := getConfig()

	// Setup the broker
	broker := &MinioServiceBroker{
		log:                log,
		serviceID:          DefaultServiceID,
		serviceName:        DefaultServiceName,
		serviceDescription: DefaultServiceDescription,
		bindableService:    true,
		planName:           DefaultPlanName,
		planID:             DefaultPlanID,
		planDescription:    DefaultPlanDescription,
		bindablePlan:       true,
		binderMgr:          NewBinderMgr(conf, log),
		instanceMgr:        NewInstanceMgr(conf, log),
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	brokerAPI := brokerapi.New(broker, log, credentials)
	http.Handle("/", brokerAPI)
	log.Info("Listening for requests")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Error("Failed to start the server", err)
	}

}
