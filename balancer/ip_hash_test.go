package balancer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPHash_Balance(t *testing.T) {
	type expect struct {
		reply string
		err   error
	}
	cases := []struct {
		name   string
		lb     Balancer
		key    string
		expect expect
	}{
		{"test-1-1",
			NewIPHash([]string{
				"http://127.0.0.1:1011",
				"http://127.0.0.1:1012",
				"http://127.0.0.1:1013",
			}),
			"192.168.1.1",
			expect{
				"http://127.0.0.1:1011",
				nil,
			},
		},
		{"test-1-2",
			NewIPHash([]string{
				"http://127.0.0.1:1011",
				"http://127.0.0.1:1012",
				"http://127.0.0.1:1013",
			}),
			"192.168.1.1",
			expect{
				"http://127.0.0.1:1011",
				nil,
			},
		},
		{
			"test-2",
			NewIPHash([]string{}),
			"192.168.1.1",
			expect{
				"",
				ErrNoHost,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			value, err := c.lb.Balance(c.key)
			assert.Equal(t, true, reflect.DeepEqual(c.expect.reply, value))
			assert.Equal(t, true, errors.Is(c.expect.err, err))
		})
	}
}
