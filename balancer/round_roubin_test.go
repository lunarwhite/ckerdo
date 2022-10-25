package balancer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundRobin_Add(t *testing.T) {
	cases := []struct {
		name   string
		lb     Balancer
		args   string
		expect Balancer
	}{
		{
			"test-1",
			NewRoundRobin([]string{
				"http://127.0.0.1:1011",
				"http://127.0.0.1:1012",
			}),
			"http://127.0.0.1:1013",
			&RoundRobin{
				BaseBalancer: BaseBalancer{
					hosts: []string{
						"http://127.0.0.1:1011",
						"http://127.0.0.1:1012",
						"http://127.0.0.1:1013",
					},
				},
				i: 0,
			},
		},
		{
			"test-2",
			NewRoundRobin([]string{
				"http://127.0.0.1:1011",
				"http://127.0.0.1:1012",
			}),
			"http://127.0.0.1:1011",
			&RoundRobin{
				BaseBalancer: BaseBalancer{
					hosts: []string{
						"http://127.0.0.1:1011",
						"http://127.0.0.1:1012",
					},
				},
				i: 0,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.lb.Add(c.args)
			assert.Equal(t, c.expect, c.lb)
		})
	}
}

func TestRoundRobin_Remove(t *testing.T) {
	cases := []struct {
		name   string
		lb     Balancer
		args   string
		expect Balancer
	}{
		{
			"test-1",
			NewRoundRobin([]string{
				"http://127.0.0.1:1011",
				"http://127.0.0.1:1012",
			}),
			"http://127.0.0.1:1013",
			&RoundRobin{
				BaseBalancer: BaseBalancer{
					hosts: []string{
						"http://127.0.0.1:1011",
						"http://127.0.0.1:1012",
					},
				},
			},
		},
		{
			"test-2",
			NewRoundRobin([]string{
				"http://127.0.0.1:1011",
				"http://127.0.0.1:1012",
			}),
			"http://127.0.0.1:1012",
			&RoundRobin{
				BaseBalancer: BaseBalancer{
					hosts: []string{
						"http://127.0.0.1:1011",
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.lb.Remove(c.args)
			assert.Equal(t, c.expect, c.lb)
		})
	}
}

func TestRoundRobin_Balance(t *testing.T) {
	type expect struct {
		reply string
		err   error
	}
	cases := []struct {
		name   string
		lb     Balancer
		args   string
		expect expect
	}{
		{
			"test-1",
			NewRoundRobin([]string{
				"http://127.0.0.1:1011",
				"http://127.0.0.1:1012",
			}),
			"",
			expect{
				"http://127.0.0.1:1011",
				nil,
			},
		},
		{
			"test-2",
			NewRoundRobin([]string{}),
			"",
			expect{
				"",
				ErrNoHost,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			reply, err := c.lb.Balance(c.args)
			assert.Equal(t, c.expect.reply, reply)
			assert.Equal(t, c.expect.err, err)
		})
	}
}
