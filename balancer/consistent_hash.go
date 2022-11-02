package balancer

import (
	"github.com/lafikl/consistent"
)

func init() {
	factories[ConsistentHashBalancer] = NewConsistent
}

// Consistent refers to consistent hash
type Consistent struct {
	BaseBalancer
	ch *consistent.Consistent
}

// NewConsistent create new Consistent balancer
func NewConsistent(hosts []string) Balancer {
	c := &Consistent{
		ch: consistent.New(),
	}
	for _, h := range hosts {
		c.ch.Add(h)
	}
	return c
}

// Add new host to the balancer
func (c *Consistent) Add(host string) {
	c.ch.Add(host)
}

// Remove new host from the balancer
func (c *Consistent) Remove(host string) {
	c.ch.Remove(host)
}

// Balance selects a suitable host according to the key value
func (c *Consistent) Balance(key string) (string, error) {
	if len(c.ch.Hosts()) == 0 {
		return "", ErrNoHost
	}
	return c.ch.Get(key)
}

// Inc refers to the number of connections to the server `+1`
func (b *Consistent) Inc(host string) {
	b.ch.Inc(host)
}

// Done refers to the number of connections to the server `-1`
func (b *Consistent) Done(host string) {
	b.ch.Done(host)
}
