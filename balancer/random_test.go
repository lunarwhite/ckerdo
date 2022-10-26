package balancer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom_Balance(t *testing.T) {
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
			NewRandom([]string{
				"http://127.0.0.1:1011",
			}),
			"",
			expect{
				"http://127.0.0.1:1011",
				nil,
			},
		},
		{
			"test-2",
			NewRandom([]string{}),
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
