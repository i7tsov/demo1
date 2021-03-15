package generator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "negative",
			input:    -2,
			expected: 4,
		},
		{
			name:     "large",
			input:    0xffff,
			expected: 0xfffe0001,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Square(c.input)
			assert.Equal(t, c.expected, actual, "square")
		})
	}
}

func TestGenerator(t *testing.T) {
	cases := []struct {
		name             string
		client           Client
		cycles           int
		expectedError    bool
		expectedCounters []int
	}{
		{
			name:          "nil client",
			client:        nil,
			expectedError: true,
		},
		{
			name:             "one",
			client:           &stubClient{},
			cycles:           1,
			expectedCounters: []int{1},
		},
		{
			name:             "5",
			client:           &stubClient{},
			cycles:           5,
			expectedCounters: []int{1, 4, 9, 16, 25},
		},
		{
			name:             "5",
			client:           &stubFaultyClient{},
			cycles:           5,
			expectedCounters: []int{1, 4, 9, 16, 25},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			o := Opts{
				Cycles: c.cycles,
				Client: c.client,
			}
			g, err := New(o)
			if c.expectedError {
				if err == nil {
					assert.FailNow(t, "Case is expecting error")
				} else {
					return
				}
			}
			if !c.expectedError && err != nil {
				assert.FailNowf(t, "Creating generator", "unexpected error: %v", err)
			}
			err = g.Run()
			//assert.NoError(t, err)
			if err == nil {
				assert.Equal(t, c.expectedCounters, c.client.(*stubClient).sets)
			}
			fmt.Println("Done.")
		})
	}
}

type stubClient struct {
	sets []int
}

func (c *stubClient) Set(key string, value interface{}) error {
	v, ok := value.(struct{ Counter int })
	if !ok {
		return fmt.Errorf("unexpected type %T, expected struct{ Counter int }", value)
	}
	c.sets = append(c.sets, v.Counter)
	return nil
}

type stubFaultyClient struct{}

func (c *stubFaultyClient) Set(key string, value interface{}) error {
	return fmt.Errorf("I am faulty")
}
