package client

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"code.cloudfoundry.org/lager"

	"github.com/minio/minio-go"
	"github.com/minio/minio-servicebroker/utils"
)

type ApiClient struct {
	log    lager.Logger
	client *minio.Client
	conf   utils.Config
}

func New(config utils.Config, logger lager.Logger) (*ApiClient, error) {
	b := ApiClient{conf: config, log: logger}
	fmt.Println("creating minio client......")
	defaultEndpoint := os.Getenv("SERVER_ENDPOINT")
	defaultAccessKey := os.Getenv("ACCESS_KEY")
	defaultSecretKey := os.Getenv("SECRET_KEY")
	defaultSecure := mustParseBool(os.Getenv("ENABLE_HTTPS"))
	minioClient, err := minio.New(defaultEndpoint, defaultAccessKey, defaultSecretKey, defaultSecure)
	if err != nil {
		return nil, errors.New("Apiclient not created")
	}
	b.client = minioClient
	return &b, nil
}
func (c *ApiClient) CreateInstance(parameters map[string]string) (string, error) {
	// Make bucket

	err := c.client.MakeBucket(parameters["instanceID"], "us-east-1")
	fmt.Println("error-====", err, c.client)
	if err != nil {
		return "", errors.New("Instance could not be provisioned")
	}
	return parameters["instanceID"], nil
}
func (client *ApiClient) GetInstanceState(instanceID string) (string, error) {
	//TODO
	return "", nil
}
func (c *ApiClient) DeleteInstance(instanceID string) error {
	err := c.client.RemoveBucket(instanceID)
	if err != nil {
		return errors.New("Instance could not be deprovisioned")
	}
	return nil
}
func (c *ApiClient) CreateBinding(parameters map[string]string) (string, error) {
	object, err := os.Open("/home/kris/code/src/github.com/minio/minio-servicebroker/data/my-testfile")
	if err != nil {
		errors.New("error loading mock binding info...")
	}
	fmt.Println("object----", object)
	defer object.Close()
	fmt.Println("object2----", object)

	_, err = c.client.PutObject(parameters["instanceID"], parameters["bindingID"], object, "application/octet-stream")
	fmt.Println("err------>", err, parameters["instanceID"], parameters["bindingID"], object)
	fmt.Println("err2------>", parameters["instanceID"], parameters["bindingID"], object)

	if err != nil {
		return "", errors.New("Instance could not be provisioned")
	}
	return parameters["bindingID"], nil
}
func (c *ApiClient) DeleteBinding(instanceId string, bindingID string) error {
	err := c.client.RemoveObject(instanceId, bindingID)
	if err != nil {
		return errors.New("Instance could not be deprovisioned")
	}
	return nil
}

// Convert string to bool and always return false if any error
func mustParseBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return b
}
