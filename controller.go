package main

import (
	"net/http"

	"github.com/minio/minio-servicebroker/utils"

	"code.cloudfoundry.org/lager"
)

type Controller struct {
	log    lager.Logger
	broker *MinioServiceBroker
}

func (c *Controller) CatalogHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("Inside cataloghandler")
	utils.WriteResponse(w, http.StatusOK, c.broker.Catalog())
}

func (c *Controller) ProvisionHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("Inside ProvisionHandler")

}

func (c *Controller) DeprovisionHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("Inside DeprovisionHandler")

}

func (c *Controller) BindHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("Inside BindHandler")

}

func (c *Controller) UnBindHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("Inside UnBindHandler")

}
