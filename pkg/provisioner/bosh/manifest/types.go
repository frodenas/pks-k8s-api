/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package boshmanifest

// BOSHNetwork is a BOSH Network definition.
type BOSHNetwork struct {
	// Name is the name of the BOSH network.
	Name string `yaml:"name"`

	// Type is the type of the BOSH network (manual, dynamic, vip).
	Type string `yaml:"type"`

	// Subnets are the subnetworks associated with the BOSH network.
	// Only used when Type is "manual".
	Subnets []BOSHSubnet `yaml:"subnets,omitempty"`

	// CloudProperties describes any IaaS-specific properties for the network.
	// Only used when Type is "dynamic" or "vip".
	CloudProperties map[string]interface{} `yaml:"cloud_properties,omitempty"`
}

// BOSHSubnet represents a BOSH Subnet
type BOSHSubnet struct {
	// Range is the subnetwork IP range that includes all IPs from this subnetwork.
	Range string `yaml:"range"`

	// Gateway is the subnetwotk gateway IP.
	Gateway string `yaml:"gateway"`

	// DNS is a list of DNS servers for this subnetwork.
	DNS []string `yaml:"dns"`

	// Reserved is a list of reserved IPs and/or IP ranges. BOSH does not assign IPs from this range to any VM.
	Reserved []string `yaml:"reserved,omitempty"`

	// Static is a list of static IPs and/or IP ranges. BOSH assigns IPs from this range to jobs requesting static IPs.
	Static []string `yaml:"static,omitempty"`

	// AZS is a list of AZs associated with this subnetwork.
	AZS []string `yaml:"azs"`

	// CloudProperties describes any IaaS-specific properties for the subnetwork.
	CloudProperties map[string]interface{} `yaml:"cloud_properties"`
}

// BOSHAZs is a BOSH Availability Zones definition.
type BOSHAZs struct {
	// AZs is a list of BOSH availability zone
	AZs []BOSHAZ `yaml:"azs"`
}

// BOSHAZ is a BOSH Availability Zone definition.
type BOSHAZ struct {
	// Name is the name of the BOSH availability zone.
	Name string `yaml:"name"`

	// CloudProperties describes any IaaS-specific properties for the availability zone.
	CloudProperties map[string]interface{} `yaml:"cloud_properties"`
}

// BoshManifest is a BOSH deployment manifest.
type BoshManifest struct {
	Name           string                 `yaml:"name"`
	Features       BoshFeatures           `yaml:"features,omitempty"`
	Releases       []Release              `yaml:"releases"`
	Stemcells      []Stemcell             `yaml:"stemcells"`
	Update         *Update                `yaml:"update"`
	InstanceGroups []InstanceGroup        `yaml:"instance_groups"`
	Addons         []Addon                `yaml:"addons,omitempty"`
	Properties     map[string]interface{} `yaml:"properties,omitempty"`
	Variables      []Variable             `yaml:"variables,omitempty"`
	Tags           map[string]interface{} `yaml:"tags,omitempty"`
}

type BoshFeatures struct {
	ConvergeVariable     *bool `yaml:"converge_variables,omitempty"`
	RandomizeAZPlacement *bool `yaml:"randomize_az_placement,omitempty"`
	UseDNSAddresses      *bool `yaml:"use_dns_addresses,omitempty"`
	UseShortDNSAddresses *bool `yaml:"use_short_dns_addresses,omitempty"`
	UseTmpfsJobConfig    *bool `yaml:"use_tmpfs_job_config,omitempty"`
}

type Release struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	URL     string `yaml:"url,omitempty"`
	SHA1    string `yaml:"sha1,omitempty"`
}

type Stemcell struct {
	Alias   string `yaml:"alias"`
	OS      string `yaml:"os"`
	Version string `yaml:"version"`
	Name    string `yaml:"name,omitempty"`
}

type Update struct {
	Canaries        int              `yaml:"canaries"`
	MaxInFlight     MaxInFlightValue `yaml:"max_in_flight"`
	CanaryWatchTime string           `yaml:"canary_watch_time"`
	UpdateWatchTime string           `yaml:"update_watch_time"`
	Serial          *bool            `yaml:"serial,omitempty"`
	VmStrategy      string           `yaml:"vm_strategy,omitempty"`
}

type MaxInFlightValue interface{}

type InstanceGroup struct {
	Name               string                 `yaml:"name,omitempty"`
	AZs                []string               `yaml:"azs,omitempty"`
	Instances          int                    `yaml:"instances"`
	Jobs               []Job                  `yaml:"jobs,omitempty"`
	VMType             string                 `yaml:"vm_type"`
	VMExtensions       []string               `yaml:"vm_extensions,omitempty"`
	VMResources        VMResource             `yaml:"vm_resources,omitempty"`
	Stemcell           string                 `yaml:"stemcell"`
	PersistentDisk     int                    `yaml:"persistent_disk,omitempty"`
	PersistentDiskType string                 `yaml:"persistent_disk_type,omitempty"`
	Networks           []Network              `yaml:"networks"`
	Update             *Update                `yaml:"update,omitempty"`
	MigratedFrom       []Migration            `yaml:"migrated_from,omitempty"`
	Lifecycle          string                 `yaml:"lifecycle,omitempty"`
	Properties         map[string]interface{} `yaml:"properties,omitempty"`
	Env                map[string]interface{} `yaml:"env,omitempty"`
}

type Job struct {
	Name       string                  `yaml:"name"`
	Release    string                  `yaml:"release"`
	Consumes   map[string]ConsumesLink `yaml:"consumes,omitempty"`
	Provides   map[string]ProvidesLink `yaml:"provides,omitempty"`
	Properties map[string]interface{}  `yaml:"properties,omitempty"`
}
type ConsumesLink struct {
	From        string `yaml:"from,omitempty"`
	Deployment  string `yaml:"deployment,omitempty"`
	Network     string `yaml:"network,omitempty"`
	IPAddresses *bool  `yaml:"ip_addresses,omitempty"`
}

type ProvidesLink struct {
	As     string `yaml:"as,omitempty"`
	Shared bool   `yaml:"shared,omitempty"`
}

type VMResource struct {
	CPU               int `yaml:"cpu"`
	RAm               int `yaml:"ram"`
	EphemeralDiskSize int `yaml:"ephemeral_disk_size"`
}

type Network struct {
	Name      string   `yaml:"name"`
	StaticIPs []string `yaml:"static_ips,omitempty"`
	Default   []string `yaml:"default,omitempty"`
}

type Migration struct {
	Name string `yaml:"name"`
}

type Addon struct {
	Name string     `yaml:"name"`
	Jobs []AddonJob `yaml:"jobs"`
}

type AddonJob struct {
	Name       string                 `yaml:"name"`
	Release    string                 `release:"release"`
	Properties map[string]interface{} `yaml:"properties,omitempty"`
}

type Variable struct {
	Name    string                 `yaml:"name"`
	Type    string                 `yaml:"type"`
	Options map[string]interface{} `yaml:"options,omitempty"`
}
