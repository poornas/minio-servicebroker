package main

import (
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"
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
	c := Controller{
		log:    log,
		broker: broker,
	}
	// Setup routers
	mux := bone.New()

	// Handle take http.Handler
	mux.Get("/v2/catalog", http.HandlerFunc(c.CatalogHandler))
	mux.Put("/v2/service_instances/{service_instance_guid}", http.HandlerFunc(c.ProvisionHandler))
	mux.Delete("/v2/service_instances/{service_instance_guid}", http.HandlerFunc(c.DeprovisionHandler))
	mux.Put("/v2/service_instances/{service_instance_guid}/service_bindings/{service_binding_guid}", http.HandlerFunc(c.BindHandler))
	mux.Delete("/v2/service_instances/{service_instance_guid}/service_bindings/{service_binding_guid}", http.HandlerFunc(c.UnBindHandler))

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
