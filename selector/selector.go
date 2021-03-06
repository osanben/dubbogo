package selector

import (
	"errors"
)

import (
	"github.com/AlexStocks/dubbogo/registry"
)

type NewSelector func(...Option) Selector

// Selector builds on the registry as a mechanism to pick nodes.
// This allows host pools and other things to be built using
// various algorithms.
type Selector interface {
	Options() Options
	// Select returns a function which should return the next node
	Select(conf registry.ServiceConfigIf) (Next, error)
	// Close renders the selector unusable
	Close() error
	// Name of the selector
	String() string
}

// Next is a function that returns the next node
// based on the selector's strategy
type Next func(ID int64) (*registry.ServiceURL, error)

var (
	ErrNotFound              = errors.New("not found")
	ErrNoneAvailable         = errors.New("none available")
	ErrRunOutAllServiceNodes = errors.New("has used out all provider nodes")
)
