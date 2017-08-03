package main

import (
	"fmt"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"
)

func main() {

	log := lager.NewLogger("minio-servicebroker")
	log.RegisterSink(lager.NewWriterSink(os.Stderr, lager.DEBUG))
	log.RegisterSink(lager.NewWriterSink(os.Stderr, lager.INFO))

	mux := bone.New()

	// Handle take http.Handler
	mux.Handle("/", http.HandlerFunc(Handler))

	// GetFunc, PostFunc etc ... takes http.HandlerFunc
	mux.GetFunc("/test", Handler)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}

// Handler - default handler to serve requests
func Handler(rw http.ResponseWriter, req *http.Request) {
	// Get the value of the "id" parameters.
	val := bone.GetValue(req, "id")
	fmt.Println("running service broker")
	rw.Write([]byte(val))
}
