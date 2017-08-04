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
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
