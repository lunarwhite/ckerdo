package balancer

import "sync"

type BaseBalancer struct {
	sync.RWMutex
	hosts []string
}

// Add new host to the balancer
func (b *BaseBalancer) Add(host string) {
	b.Lock()
	defer b.Unlock()
	for _, h := range b.hosts {
		if h == host {
			return
		}
	}
	b.hosts = append(b.hosts, host)
}

// Remove new host from the balancer
func (b *BaseBalancer) Remove(host string) {
	b.Lock()
	defer b.Unlock()
	for i, h := range b.hosts {
		if h == host {
			b.hosts = append(b.hosts[:i], b.hosts[i+1:]...)
			return
		}
	}
}

// Balance selects a suitable host according to the key value
func (b *BaseBalancer) Balance(key string) (string, error) {
	return "", nil
}

// Inc refers to the number of connections to the server `+1`
func (b *BaseBalancer) Inc(_ string) {}

// Done refers to the number of connections to the server `-1`
func (b *BaseBalancer) Done(_ string) {}
