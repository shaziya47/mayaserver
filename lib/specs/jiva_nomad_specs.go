package specs

import (
	"github.com/openebs/mayaserver/lib/api/v1"
)

// jNPSpecs acts as a placeholder of jiva storage properties that can be
// transformed into Nomad's placement specs otherwise known as a
// job specs.
//
// It implements following interface(s):
//  1.  orchspecs.PlacementSpecs
type jNPSpecs struct {

	// Name of the jiva volume
	volName string

	// ID of the jiva volume
	volID string

	// Region where this jiva volume will be placed
	region string

	// The datacenter where this jiva volume will be placed
	dc string

	// jiva front end volume's size
	feVolSize string

	// jiva back end volume's size
	beVolSize string

	// Total no. of jiva front end volumes
	feCount int

	// Total no. of jiva back end volumes
	beCount int

	// The backing stor of jiva back end volume
	beVolumeStor string

	// jiva front end image version
	feVersion string

	// jiva back end image version
	beVersion string

	// jiva volume's container network type
	networkType string

	// comma separated list of front-end IPs
	feIPs []string

	// comma separated list of back-end IPs
	beIPs []string

	// jiva front end container's subnet
	feSubnet string

	// jiva front end container's network interface
	feInterface string
}

// NewJNPSpecs creates a jNPSpecs instance based on the provided
// persistent volume claim
func NewJNPSpecs(pvc *v1.PersistentVolumeClaim) *jNPSpecs {
	return &jNPSpecs{}
}

// Name provides the name of this placement specs instance
func (j *jNPSpecs) Name() string {
	return v1.JivaNomadPlacementSpecs
}

// NewPlacementSpecs will generate specs/intent that is understood by
// Nomad. This specs determine the scheduling & placement of jiva
// storage based containers
func (j *jNPSpecs) NewPlacementSpecs() interface{} {
	return nil
}
