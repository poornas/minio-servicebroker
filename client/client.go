package client

// Client does the actual instance management.
type Client interface {
	CreateInstance(parameters interface{}) (string, error)
	GetInstanceState(instanceId string) (string, error)
	DeleteInstance(instanceId string) error
}
