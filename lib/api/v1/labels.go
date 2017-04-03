package v1

// A type that acts on an infrastructure level i.e. orchestrator hosts
type ContainerNetworkingLbl string

const (
	CNTypeLbl            ContainerNetworkingLbl = "cn.openebs.io/type"
	CNNetworkCIDRAddrLbl ContainerNetworkingLbl = "cn.openebs.io/network-cidr-addr"
	CNSubnetLbl          ContainerNetworkingLbl = "cn.openebs.io/subnet"
	CNInterfaceLbl       ContainerNetworkingLbl = "cn.openebs.io/interface"
)

// A type that acts on an infrastructure level i.e. orchestrator hosts
type ContainerStorageLbl string

const (
	CSPersistenceLocationLbl ContainerStorageLbl = "cs.openebs.io/persistence-location"
	CSReplicaCountLbl        ContainerStorageLbl = "cs.openebs.io/replica-count"
)

type RequestsLbl string

const (
	RegionLbl     RequestsLbl = "requests.openebs.io/region"
	DatacenterLbl RequestsLbl = "requests.openebs.io/dc"
)

const (
	VolumePluginNamePrefix string = "name.plugin.volume.openebs.io/"
)

const (
	DefaultOrchestratorConfigPath string = "/etc/mayaserver/orchprovider/"
)

const (
	JivaNomadPlacementSpecs string = "placement.specs.openebs.io/jnp-specs"
	JivaK8sPlacementSpecs   string = "placement.specs.openebs.io/jk8sp-specs"
)
