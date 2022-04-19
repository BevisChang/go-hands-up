package localcache

type Cache interface {
	// Get cache by key
	Get(string) (interface{}, error)

	// Set cache via provides key and value, you can provide any type of value
	Set(string, interface{}) error
}
