package main

import (
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"
)

func main() {
	// Create logger
	log := lager.NewLogger("minio-servicebroker")
	log.RegisterSink(lager.NewWriterSink(os.Stderr, lager.DEBUG))
	log.RegisterSink(lager.NewWriterSink(os.Stderr, lager.INFO))

	mux := bone.New()
	c := Controller{log: log}
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
