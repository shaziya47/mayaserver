// This is a **Work In Progress** item.
//
// This is being currently developed to create Nomad specs that can be fed
// to Nomad APIs. These specs will be generated based on persistent storage
// properties i.e. scheduling, placement, QoS, etc.
//
// NOTE:
// This will be an effective replacement for lib/orchprovider/nomad/helper_funcs.go.
//
// This will be designed in a way that can cater to K8s specs as well.
package specs

type PlacementSpecs interface {
	Name() string

	// This will generate specs/intent that is understood by various
	// orchestrators e.g. Nomad/K8s in order to place storage containers
	NewPlacementSpecs() interface{}
}
