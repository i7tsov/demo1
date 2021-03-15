// Package generator implements Generator for data.
package generator

import (
	"fmt"
	"log"
	"time"
)

// Generator generates some data.
type Generator struct {
	o Opts
}

// Opts is options for Generator.
type Opts struct {
	Client Client
	Cycles int
	Delay  time.Duration
}

// Client ...
type Client interface {
	Set(key string, value interface{}) error
}

// New ...
func New(o Opts) (*Generator, error) {
	if o.Client == nil {
		return nil, fmt.Errorf("client is required")
	}
	return &Generator{
		o: o,
	}, nil
}

func Square(x int) int {
	return x * x
}

// Run ...
func (g *Generator) Run() error {
	counter := 0
	for {
		if g.o.Cycles != 0 && counter >= g.o.Cycles {
			return nil
		}
		counter++
		v := struct {
			Counter int
		}{
			Counter: Square(counter),
		}
		err := g.o.Client.Set("foo", v)
		if err != nil {
			return err
		}
		log.Print("Generated counter")
		time.Sleep(g.o.Delay)
	}
}
