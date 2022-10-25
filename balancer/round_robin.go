package balancer

// RoundRobin will select the server in turn from the server to proxy
type RoundRobin struct {
	BaseBalancer
	i uint64
}

func init() {
	factories[R2Balancer] = NewRoundRobin
}

// NewRoundRobin create new RoundRobin balancer
func NewRoundRobin(hosts []string) Balancer {
	return &RoundRobin{
		i: 0,
		BaseBalancer: BaseBalancer{
			hosts: hosts,
		},
	}
}

// Balance selects a suitable host according to the key value
func (r *RoundRobin) Balance(_ string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.hosts) == 0 {
		return "", ErrNoHost
	}
	host := r.hosts[r.i%uint64(len(r.hosts))]
	r.i++
	return host, nil
}
