package localcache

// Cache interface is defined the behaviour of cache, your implement should include basic get & set method
type Cache interface {

	// Get cache by key
	Get(string) (interface{}, error)

	// Set cache via provides key and value, you can provide any type of value
	Set(string, interface{}) error
}
