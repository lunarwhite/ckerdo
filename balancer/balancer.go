package balancer

import "errors"

var (
	ErrNoHost           = errors.New("no host")
	ErrAlgoNotSupported = errors.New("algorithm not supported")
)

// Balancer interface is the load balancer for the reverse proxy
type Balancer interface {
	Add(string)
	Remove(string)
	Balance(string) (string, error)
	Inc(string)
	Done(string)
}

// Factory is the factory that generates Balancer
type Factory func([]string) Balancer

var factories = make(map[string]Factory)

// Build generates the corresponding Balancer according to the algorithm
func Build(algorithm string, hosts []string) (Balancer, error) {
	factory, ok := factories[algorithm]
	if !ok {
		return nil, ErrAlgoNotSupported
	}
	return factory(hosts), nil
}
