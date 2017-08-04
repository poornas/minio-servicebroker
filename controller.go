package main

import (
	"net/http"

	"code.cloudfoundry.org/lager"
)

type Controller struct {
	log lager.Logger
}

func (c *Controller) CatalogHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("Inside cataloghandler")
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
