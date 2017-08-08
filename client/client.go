package client

// Client does the actual instance management.
type Client interface {
	CreateInstance(parameters map[string]interface{}) (string, error)
	GetInstanceState(instanceId string) (string, error)
	DeleteInstance(instanceId string) error
	CreateBinding(parameters map[string]interface{}) (string, error)
	DeleteBinding(instanceId string, bindingID string) error
}
