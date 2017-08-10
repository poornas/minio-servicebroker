package client

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"code.cloudfoundry.org/lager"

	"github.com/minio/minio-go"
	"github.com/poornas/minio-servicebroker/utils"
)

type ApiClient struct {
	log    lager.Logger
	client *minio.Client
	conf   utils.Config
}

func New(config utils.Config, logger lager.Logger) (*ApiClient, error) {
	b := ApiClient{conf: config, log: logger}
	fmt.Println("creating minio client......")
	// defaultEndpoint := os.Getenv("SERVER_ENDPOINT")
	// defaultAccessKey := os.Getenv("ACCESS_KEY")
	// defaultSecretKey := os.Getenv("SECRET_KEY")
	// defaultSecure := mustParseBool(os.Getenv("ENABLE_HTTPS"))

	// minioClient, err := minio.New(defaultEndpoint, defaultAccessKey, defaultSecretKey, defaultSecure)
	minioClient, err := minio.New("play.minio.io:9000", "Q3AM3UQ867SPQQA43P2F", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", true)
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
	fmt.Println("Should have created me ===%s", parameters["instanceID"])
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
	var buf []byte
	for i := 0; i < 4; i++ {
		buf = append(buf, bytes.Repeat([]byte(string('a'+i)), 1000)...)
	}

	_, err := c.client.PutObject(parameters["instanceID"], parameters["bindingID"], bytes.NewReader(buf), "application/octet-stream")

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
