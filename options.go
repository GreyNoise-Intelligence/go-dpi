// Package godpi provides the main API interface for utilizing the go-dpi library.
package go_dpi

import (
	"github.com/GreyNoise-Intelligence/go-dpi/modules/classifiers"
	"github.com/GreyNoise-Intelligence/go-dpi/types"
)

// Options allow end users init module with custom options
// NOTE. it's necessary to check the module passed in Apply func
type Options interface {
	Apply(types.Module)
}

// ClassifierOption take classifier options to override default values
// for now this option was added for test
type ClassifierOption struct {
	// TODO
}

func (o ClassifierOption) Apply(mod types.Module) {
	// check module
	_, ok := mod.(*classifiers.ClassifierModule)
	if !ok {
		return
	}
	// TODO
}
