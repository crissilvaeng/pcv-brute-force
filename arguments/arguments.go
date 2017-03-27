package arguments

import (
	"errors"
	"io/ioutil"
)

const (
	filename = 1
)

// Argumentative returns the information passed as command line arguments.
type Argumentative interface {
	// Input returns a raw array of bytes from file input.
	Input() ([]byte, error)
}

type arguments struct {
	values []string
}

func (a *arguments) Input() ([]byte, error) {
	buffer, err := ioutil.ReadFile(a.values[filename])
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// New valid the argumens and returns an Argumentative instance.
func New(args []string) (Argumentative, error) {
	if len(args) < 2 {
		return nil, errors.New("usase: go run main.go routes.csv")
	}
	return &arguments{values: args}, nil
}
