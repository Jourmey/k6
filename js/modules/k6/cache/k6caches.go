package k6cache

import (
	"go.k6.io/k6/js/modules"
	"sync"
)

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	K6Cache struct {
		defaultCache *sync.Map
	}
)

var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &K6Cache{}
)

var (
	globalCache = sync.Map{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &K6Cache{
		defaultCache: &globalCache,
	}
}

// Exports returns the exports of the execution module.
func (k6cache *K6Cache) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"Store":  k6cache.Store,
			"Load":   k6cache.Load,
			"Delete": k6cache.Delete,
		},
	}
}
