package client

import (
	"log"

	"github.com/minio/minio-go"
)

type backend struct {
	client *client
}

func New(config *Config) *backend {
	backend.client, err = minio.New("play.minio.io:9000", "YOUR-ACCESS", "YOUR-SECRET", true)
	if err != nil {
		log.Fatalln(err)
	}
	return backend.client
}
func (b *backend) CreateInstance(parameters interface{}) (string, error) {

}
func (b *backend) GetInstanceState(instanceId string) (string, error) {

}
func (b *backend) DeleteInstance(instanceId string) error {

}
