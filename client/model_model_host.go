/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ModelHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelHost{}

// ModelHost struct for ModelHost
type ModelHost struct {
	AgentRunning bool `json:"agent_running"`
	AvailabilityZone string `json:"availability_zone"`
	CloudAccountId string `json:"cloud_account_id"`
	CloudProvider string `json:"cloud_provider"`
	CloudRegion string `json:"cloud_region"`
	ComplianceLatestScanId string `json:"compliance_latest_scan_id"`
	ComplianceScanStatus string `json:"compliance_scan_status"`
	CompliancesCount int32 `json:"compliances_count"`
	ContainerImages []ModelContainerImage `json:"container_images"`
	Containers []ModelContainer `json:"containers"`
	CpuMax float32 `json:"cpu_max"`
	CpuUsage float32 `json:"cpu_usage"`
	HostName string `json:"host_name"`
	InboundConnections []ModelConnection `json:"inbound_connections"`
	InstanceId string `json:"instance_id"`
	InstanceType string `json:"instance_type"`
	IsConsoleVm bool `json:"is_console_vm"`
	KernelId string `json:"kernel_id"`
	KernelVersion string `json:"kernel_version"`
	LocalCidr []interface{} `json:"local_cidr"`
	LocalNetworks []interface{} `json:"local_networks"`
	MalwareLatestScanId string `json:"malware_latest_scan_id"`
	MalwareScanStatus string `json:"malware_scan_status"`
	MalwaresCount int32 `json:"malwares_count"`
	MemoryMax int32 `json:"memory_max"`
	MemoryUsage int32 `json:"memory_usage"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	Os string `json:"os"`
	OutboundConnections []ModelConnection `json:"outbound_connections"`
	Pods []ModelPod `json:"pods"`
	PrivateIp []interface{} `json:"private_ip"`
	Processes []ModelProcess `json:"processes"`
	PublicIp []interface{} `json:"public_ip"`
	ResourceGroup string `json:"resource_group"`
	SecretLatestScanId string `json:"secret_latest_scan_id"`
	SecretScanStatus string `json:"secret_scan_status"`
	SecretsCount int32 `json:"secrets_count"`
	Uptime int32 `json:"uptime"`
	Version string `json:"version"`
	VulnerabilitiesCount int32 `json:"vulnerabilities_count"`
	VulnerabilityLatestScanId string `json:"vulnerability_latest_scan_id"`
	VulnerabilityScanStatus string `json:"vulnerability_scan_status"`
}

type _ModelHost ModelHost

// NewModelHost instantiates a new ModelHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelHost(agentRunning bool, availabilityZone string, cloudAccountId string, cloudProvider string, cloudRegion string, complianceLatestScanId string, complianceScanStatus string, compliancesCount int32, containerImages []ModelContainerImage, containers []ModelContainer, cpuMax float32, cpuUsage float32, hostName string, inboundConnections []ModelConnection, instanceId string, instanceType string, isConsoleVm bool, kernelId string, kernelVersion string, localCidr []interface{}, localNetworks []interface{}, malwareLatestScanId string, malwareScanStatus string, malwaresCount int32, memoryMax int32, memoryUsage int32, nodeId string, nodeName string, os string, outboundConnections []ModelConnection, pods []ModelPod, privateIp []interface{}, processes []ModelProcess, publicIp []interface{}, resourceGroup string, secretLatestScanId string, secretScanStatus string, secretsCount int32, uptime int32, version string, vulnerabilitiesCount int32, vulnerabilityLatestScanId string, vulnerabilityScanStatus string) *ModelHost {
	this := ModelHost{}
	this.AgentRunning = agentRunning
	this.AvailabilityZone = availabilityZone
	this.CloudAccountId = cloudAccountId
	this.CloudProvider = cloudProvider
	this.CloudRegion = cloudRegion
	this.ComplianceLatestScanId = complianceLatestScanId
	this.ComplianceScanStatus = complianceScanStatus
	this.CompliancesCount = compliancesCount
	this.ContainerImages = containerImages
	this.Containers = containers
	this.CpuMax = cpuMax
	this.CpuUsage = cpuUsage
	this.HostName = hostName
	this.InboundConnections = inboundConnections
	this.InstanceId = instanceId
	this.InstanceType = instanceType
	this.IsConsoleVm = isConsoleVm
	this.KernelId = kernelId
	this.KernelVersion = kernelVersion
	this.LocalCidr = localCidr
	this.LocalNetworks = localNetworks
	this.MalwareLatestScanId = malwareLatestScanId
	this.MalwareScanStatus = malwareScanStatus
	this.MalwaresCount = malwaresCount
	this.MemoryMax = memoryMax
	this.MemoryUsage = memoryUsage
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.Os = os
	this.OutboundConnections = outboundConnections
	this.Pods = pods
	this.PrivateIp = privateIp
	this.Processes = processes
	this.PublicIp = publicIp
	this.ResourceGroup = resourceGroup
	this.SecretLatestScanId = secretLatestScanId
	this.SecretScanStatus = secretScanStatus
	this.SecretsCount = secretsCount
	this.Uptime = uptime
	this.Version = version
	this.VulnerabilitiesCount = vulnerabilitiesCount
	this.VulnerabilityLatestScanId = vulnerabilityLatestScanId
	this.VulnerabilityScanStatus = vulnerabilityScanStatus
	return &this
}

// NewModelHostWithDefaults instantiates a new ModelHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelHostWithDefaults() *ModelHost {
	this := ModelHost{}
	return &this
}

// GetAgentRunning returns the AgentRunning field value
func (o *ModelHost) GetAgentRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgentRunning
}

// GetAgentRunningOk returns a tuple with the AgentRunning field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetAgentRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentRunning, true
}

// SetAgentRunning sets field value
func (o *ModelHost) SetAgentRunning(v bool) {
	o.AgentRunning = v
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *ModelHost) GetAvailabilityZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *ModelHost) SetAvailabilityZone(v string) {
	o.AvailabilityZone = v
}

// GetCloudAccountId returns the CloudAccountId field value
func (o *ModelHost) GetCloudAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudAccountId
}

// GetCloudAccountIdOk returns a tuple with the CloudAccountId field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetCloudAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudAccountId, true
}

// SetCloudAccountId sets field value
func (o *ModelHost) SetCloudAccountId(v string) {
	o.CloudAccountId = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *ModelHost) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ModelHost) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCloudRegion returns the CloudRegion field value
func (o *ModelHost) GetCloudRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudRegion
}

// GetCloudRegionOk returns a tuple with the CloudRegion field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetCloudRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudRegion, true
}

// SetCloudRegion sets field value
func (o *ModelHost) SetCloudRegion(v string) {
	o.CloudRegion = v
}

// GetComplianceLatestScanId returns the ComplianceLatestScanId field value
func (o *ModelHost) GetComplianceLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceLatestScanId
}

// GetComplianceLatestScanIdOk returns a tuple with the ComplianceLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetComplianceLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceLatestScanId, true
}

// SetComplianceLatestScanId sets field value
func (o *ModelHost) SetComplianceLatestScanId(v string) {
	o.ComplianceLatestScanId = v
}

// GetComplianceScanStatus returns the ComplianceScanStatus field value
func (o *ModelHost) GetComplianceScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceScanStatus
}

// GetComplianceScanStatusOk returns a tuple with the ComplianceScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetComplianceScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceScanStatus, true
}

// SetComplianceScanStatus sets field value
func (o *ModelHost) SetComplianceScanStatus(v string) {
	o.ComplianceScanStatus = v
}

// GetCompliancesCount returns the CompliancesCount field value
func (o *ModelHost) GetCompliancesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompliancesCount
}

// GetCompliancesCountOk returns a tuple with the CompliancesCount field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetCompliancesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompliancesCount, true
}

// SetCompliancesCount sets field value
func (o *ModelHost) SetCompliancesCount(v int32) {
	o.CompliancesCount = v
}

// GetContainerImages returns the ContainerImages field value
// If the value is explicit nil, the zero value for []ModelContainerImage will be returned
func (o *ModelHost) GetContainerImages() []ModelContainerImage {
	if o == nil {
		var ret []ModelContainerImage
		return ret
	}

	return o.ContainerImages
}

// GetContainerImagesOk returns a tuple with the ContainerImages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetContainerImagesOk() ([]ModelContainerImage, bool) {
	if o == nil || IsNil(o.ContainerImages) {
		return nil, false
	}
	return o.ContainerImages, true
}

// SetContainerImages sets field value
func (o *ModelHost) SetContainerImages(v []ModelContainerImage) {
	o.ContainerImages = v
}

// GetContainers returns the Containers field value
// If the value is explicit nil, the zero value for []ModelContainer will be returned
func (o *ModelHost) GetContainers() []ModelContainer {
	if o == nil {
		var ret []ModelContainer
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetContainersOk() ([]ModelContainer, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *ModelHost) SetContainers(v []ModelContainer) {
	o.Containers = v
}

// GetCpuMax returns the CpuMax field value
func (o *ModelHost) GetCpuMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuMax
}

// GetCpuMaxOk returns a tuple with the CpuMax field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetCpuMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuMax, true
}

// SetCpuMax sets field value
func (o *ModelHost) SetCpuMax(v float32) {
	o.CpuMax = v
}

// GetCpuUsage returns the CpuUsage field value
func (o *ModelHost) GetCpuUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetCpuUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuUsage, true
}

// SetCpuUsage sets field value
func (o *ModelHost) SetCpuUsage(v float32) {
	o.CpuUsage = v
}

// GetHostName returns the HostName field value
func (o *ModelHost) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelHost) SetHostName(v string) {
	o.HostName = v
}

// GetInboundConnections returns the InboundConnections field value
// If the value is explicit nil, the zero value for []ModelConnection will be returned
func (o *ModelHost) GetInboundConnections() []ModelConnection {
	if o == nil {
		var ret []ModelConnection
		return ret
	}

	return o.InboundConnections
}

// GetInboundConnectionsOk returns a tuple with the InboundConnections field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetInboundConnectionsOk() ([]ModelConnection, bool) {
	if o == nil || IsNil(o.InboundConnections) {
		return nil, false
	}
	return o.InboundConnections, true
}

// SetInboundConnections sets field value
func (o *ModelHost) SetInboundConnections(v []ModelConnection) {
	o.InboundConnections = v
}

// GetInstanceId returns the InstanceId field value
func (o *ModelHost) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *ModelHost) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetInstanceType returns the InstanceType field value
func (o *ModelHost) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *ModelHost) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetIsConsoleVm returns the IsConsoleVm field value
func (o *ModelHost) GetIsConsoleVm() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsConsoleVm
}

// GetIsConsoleVmOk returns a tuple with the IsConsoleVm field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetIsConsoleVmOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsConsoleVm, true
}

// SetIsConsoleVm sets field value
func (o *ModelHost) SetIsConsoleVm(v bool) {
	o.IsConsoleVm = v
}

// GetKernelId returns the KernelId field value
func (o *ModelHost) GetKernelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KernelId
}

// GetKernelIdOk returns a tuple with the KernelId field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetKernelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KernelId, true
}

// SetKernelId sets field value
func (o *ModelHost) SetKernelId(v string) {
	o.KernelId = v
}

// GetKernelVersion returns the KernelVersion field value
func (o *ModelHost) GetKernelVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KernelVersion
}

// GetKernelVersionOk returns a tuple with the KernelVersion field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetKernelVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KernelVersion, true
}

// SetKernelVersion sets field value
func (o *ModelHost) SetKernelVersion(v string) {
	o.KernelVersion = v
}

// GetLocalCidr returns the LocalCidr field value
// If the value is explicit nil, the zero value for []interface{} will be returned
func (o *ModelHost) GetLocalCidr() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.LocalCidr
}

// GetLocalCidrOk returns a tuple with the LocalCidr field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetLocalCidrOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.LocalCidr) {
		return nil, false
	}
	return o.LocalCidr, true
}

// SetLocalCidr sets field value
func (o *ModelHost) SetLocalCidr(v []interface{}) {
	o.LocalCidr = v
}

// GetLocalNetworks returns the LocalNetworks field value
// If the value is explicit nil, the zero value for []interface{} will be returned
func (o *ModelHost) GetLocalNetworks() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.LocalNetworks
}

// GetLocalNetworksOk returns a tuple with the LocalNetworks field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetLocalNetworksOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.LocalNetworks) {
		return nil, false
	}
	return o.LocalNetworks, true
}

// SetLocalNetworks sets field value
func (o *ModelHost) SetLocalNetworks(v []interface{}) {
	o.LocalNetworks = v
}

// GetMalwareLatestScanId returns the MalwareLatestScanId field value
func (o *ModelHost) GetMalwareLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MalwareLatestScanId
}

// GetMalwareLatestScanIdOk returns a tuple with the MalwareLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetMalwareLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwareLatestScanId, true
}

// SetMalwareLatestScanId sets field value
func (o *ModelHost) SetMalwareLatestScanId(v string) {
	o.MalwareLatestScanId = v
}

// GetMalwareScanStatus returns the MalwareScanStatus field value
func (o *ModelHost) GetMalwareScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MalwareScanStatus
}

// GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetMalwareScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwareScanStatus, true
}

// SetMalwareScanStatus sets field value
func (o *ModelHost) SetMalwareScanStatus(v string) {
	o.MalwareScanStatus = v
}

// GetMalwaresCount returns the MalwaresCount field value
func (o *ModelHost) GetMalwaresCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MalwaresCount
}

// GetMalwaresCountOk returns a tuple with the MalwaresCount field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetMalwaresCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwaresCount, true
}

// SetMalwaresCount sets field value
func (o *ModelHost) SetMalwaresCount(v int32) {
	o.MalwaresCount = v
}

// GetMemoryMax returns the MemoryMax field value
func (o *ModelHost) GetMemoryMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryMax
}

// GetMemoryMaxOk returns a tuple with the MemoryMax field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetMemoryMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryMax, true
}

// SetMemoryMax sets field value
func (o *ModelHost) SetMemoryMax(v int32) {
	o.MemoryMax = v
}

// GetMemoryUsage returns the MemoryUsage field value
func (o *ModelHost) GetMemoryUsage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetMemoryUsageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryUsage, true
}

// SetMemoryUsage sets field value
func (o *ModelHost) SetMemoryUsage(v int32) {
	o.MemoryUsage = v
}

// GetNodeId returns the NodeId field value
func (o *ModelHost) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelHost) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelHost) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelHost) SetNodeName(v string) {
	o.NodeName = v
}

// GetOs returns the Os field value
func (o *ModelHost) GetOs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Os
}

// GetOsOk returns a tuple with the Os field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Os, true
}

// SetOs sets field value
func (o *ModelHost) SetOs(v string) {
	o.Os = v
}

// GetOutboundConnections returns the OutboundConnections field value
// If the value is explicit nil, the zero value for []ModelConnection will be returned
func (o *ModelHost) GetOutboundConnections() []ModelConnection {
	if o == nil {
		var ret []ModelConnection
		return ret
	}

	return o.OutboundConnections
}

// GetOutboundConnectionsOk returns a tuple with the OutboundConnections field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetOutboundConnectionsOk() ([]ModelConnection, bool) {
	if o == nil || IsNil(o.OutboundConnections) {
		return nil, false
	}
	return o.OutboundConnections, true
}

// SetOutboundConnections sets field value
func (o *ModelHost) SetOutboundConnections(v []ModelConnection) {
	o.OutboundConnections = v
}

// GetPods returns the Pods field value
// If the value is explicit nil, the zero value for []ModelPod will be returned
func (o *ModelHost) GetPods() []ModelPod {
	if o == nil {
		var ret []ModelPod
		return ret
	}

	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetPodsOk() ([]ModelPod, bool) {
	if o == nil || IsNil(o.Pods) {
		return nil, false
	}
	return o.Pods, true
}

// SetPods sets field value
func (o *ModelHost) SetPods(v []ModelPod) {
	o.Pods = v
}

// GetPrivateIp returns the PrivateIp field value
// If the value is explicit nil, the zero value for []interface{} will be returned
func (o *ModelHost) GetPrivateIp() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetPrivateIpOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.PrivateIp) {
		return nil, false
	}
	return o.PrivateIp, true
}

// SetPrivateIp sets field value
func (o *ModelHost) SetPrivateIp(v []interface{}) {
	o.PrivateIp = v
}

// GetProcesses returns the Processes field value
// If the value is explicit nil, the zero value for []ModelProcess will be returned
func (o *ModelHost) GetProcesses() []ModelProcess {
	if o == nil {
		var ret []ModelProcess
		return ret
	}

	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetProcessesOk() ([]ModelProcess, bool) {
	if o == nil || IsNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// SetProcesses sets field value
func (o *ModelHost) SetProcesses(v []ModelProcess) {
	o.Processes = v
}

// GetPublicIp returns the PublicIp field value
// If the value is explicit nil, the zero value for []interface{} will be returned
func (o *ModelHost) GetPublicIp() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelHost) GetPublicIpOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// SetPublicIp sets field value
func (o *ModelHost) SetPublicIp(v []interface{}) {
	o.PublicIp = v
}

// GetResourceGroup returns the ResourceGroup field value
func (o *ModelHost) GetResourceGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetResourceGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroup, true
}

// SetResourceGroup sets field value
func (o *ModelHost) SetResourceGroup(v string) {
	o.ResourceGroup = v
}

// GetSecretLatestScanId returns the SecretLatestScanId field value
func (o *ModelHost) GetSecretLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretLatestScanId
}

// GetSecretLatestScanIdOk returns a tuple with the SecretLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetSecretLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretLatestScanId, true
}

// SetSecretLatestScanId sets field value
func (o *ModelHost) SetSecretLatestScanId(v string) {
	o.SecretLatestScanId = v
}

// GetSecretScanStatus returns the SecretScanStatus field value
func (o *ModelHost) GetSecretScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretScanStatus
}

// GetSecretScanStatusOk returns a tuple with the SecretScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetSecretScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretScanStatus, true
}

// SetSecretScanStatus sets field value
func (o *ModelHost) SetSecretScanStatus(v string) {
	o.SecretScanStatus = v
}

// GetSecretsCount returns the SecretsCount field value
func (o *ModelHost) GetSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsCount, true
}

// SetSecretsCount sets field value
func (o *ModelHost) SetSecretsCount(v int32) {
	o.SecretsCount = v
}

// GetUptime returns the Uptime field value
func (o *ModelHost) GetUptime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetUptimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uptime, true
}

// SetUptime sets field value
func (o *ModelHost) SetUptime(v int32) {
	o.Uptime = v
}

// GetVersion returns the Version field value
func (o *ModelHost) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ModelHost) SetVersion(v string) {
	o.Version = v
}

// GetVulnerabilitiesCount returns the VulnerabilitiesCount field value
func (o *ModelHost) GetVulnerabilitiesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VulnerabilitiesCount
}

// GetVulnerabilitiesCountOk returns a tuple with the VulnerabilitiesCount field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetVulnerabilitiesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilitiesCount, true
}

// SetVulnerabilitiesCount sets field value
func (o *ModelHost) SetVulnerabilitiesCount(v int32) {
	o.VulnerabilitiesCount = v
}

// GetVulnerabilityLatestScanId returns the VulnerabilityLatestScanId field value
func (o *ModelHost) GetVulnerabilityLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VulnerabilityLatestScanId
}

// GetVulnerabilityLatestScanIdOk returns a tuple with the VulnerabilityLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetVulnerabilityLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityLatestScanId, true
}

// SetVulnerabilityLatestScanId sets field value
func (o *ModelHost) SetVulnerabilityLatestScanId(v string) {
	o.VulnerabilityLatestScanId = v
}

// GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field value
func (o *ModelHost) GetVulnerabilityScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VulnerabilityScanStatus
}

// GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelHost) GetVulnerabilityScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityScanStatus, true
}

// SetVulnerabilityScanStatus sets field value
func (o *ModelHost) SetVulnerabilityScanStatus(v string) {
	o.VulnerabilityScanStatus = v
}

func (o ModelHost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelHost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent_running"] = o.AgentRunning
	toSerialize["availability_zone"] = o.AvailabilityZone
	toSerialize["cloud_account_id"] = o.CloudAccountId
	toSerialize["cloud_provider"] = o.CloudProvider
	toSerialize["cloud_region"] = o.CloudRegion
	toSerialize["compliance_latest_scan_id"] = o.ComplianceLatestScanId
	toSerialize["compliance_scan_status"] = o.ComplianceScanStatus
	toSerialize["compliances_count"] = o.CompliancesCount
	if o.ContainerImages != nil {
		toSerialize["container_images"] = o.ContainerImages
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	toSerialize["cpu_max"] = o.CpuMax
	toSerialize["cpu_usage"] = o.CpuUsage
	toSerialize["host_name"] = o.HostName
	if o.InboundConnections != nil {
		toSerialize["inbound_connections"] = o.InboundConnections
	}
	toSerialize["instance_id"] = o.InstanceId
	toSerialize["instance_type"] = o.InstanceType
	toSerialize["is_console_vm"] = o.IsConsoleVm
	toSerialize["kernel_id"] = o.KernelId
	toSerialize["kernel_version"] = o.KernelVersion
	if o.LocalCidr != nil {
		toSerialize["local_cidr"] = o.LocalCidr
	}
	if o.LocalNetworks != nil {
		toSerialize["local_networks"] = o.LocalNetworks
	}
	toSerialize["malware_latest_scan_id"] = o.MalwareLatestScanId
	toSerialize["malware_scan_status"] = o.MalwareScanStatus
	toSerialize["malwares_count"] = o.MalwaresCount
	toSerialize["memory_max"] = o.MemoryMax
	toSerialize["memory_usage"] = o.MemoryUsage
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["os"] = o.Os
	if o.OutboundConnections != nil {
		toSerialize["outbound_connections"] = o.OutboundConnections
	}
	if o.Pods != nil {
		toSerialize["pods"] = o.Pods
	}
	if o.PrivateIp != nil {
		toSerialize["private_ip"] = o.PrivateIp
	}
	if o.Processes != nil {
		toSerialize["processes"] = o.Processes
	}
	if o.PublicIp != nil {
		toSerialize["public_ip"] = o.PublicIp
	}
	toSerialize["resource_group"] = o.ResourceGroup
	toSerialize["secret_latest_scan_id"] = o.SecretLatestScanId
	toSerialize["secret_scan_status"] = o.SecretScanStatus
	toSerialize["secrets_count"] = o.SecretsCount
	toSerialize["uptime"] = o.Uptime
	toSerialize["version"] = o.Version
	toSerialize["vulnerabilities_count"] = o.VulnerabilitiesCount
	toSerialize["vulnerability_latest_scan_id"] = o.VulnerabilityLatestScanId
	toSerialize["vulnerability_scan_status"] = o.VulnerabilityScanStatus
	return toSerialize, nil
}

func (o *ModelHost) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent_running",
		"availability_zone",
		"cloud_account_id",
		"cloud_provider",
		"cloud_region",
		"compliance_latest_scan_id",
		"compliance_scan_status",
		"compliances_count",
		"container_images",
		"containers",
		"cpu_max",
		"cpu_usage",
		"host_name",
		"inbound_connections",
		"instance_id",
		"instance_type",
		"is_console_vm",
		"kernel_id",
		"kernel_version",
		"local_cidr",
		"local_networks",
		"malware_latest_scan_id",
		"malware_scan_status",
		"malwares_count",
		"memory_max",
		"memory_usage",
		"node_id",
		"node_name",
		"os",
		"outbound_connections",
		"pods",
		"private_ip",
		"processes",
		"public_ip",
		"resource_group",
		"secret_latest_scan_id",
		"secret_scan_status",
		"secrets_count",
		"uptime",
		"version",
		"vulnerabilities_count",
		"vulnerability_latest_scan_id",
		"vulnerability_scan_status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelHost := _ModelHost{}

	err = json.Unmarshal(bytes, &varModelHost)

	if err != nil {
		return err
	}

	*o = ModelHost(varModelHost)

	return err
}

type NullableModelHost struct {
	value *ModelHost
	isSet bool
}

func (v NullableModelHost) Get() *ModelHost {
	return v.value
}

func (v *NullableModelHost) Set(val *ModelHost) {
	v.value = val
	v.isSet = true
}

func (v NullableModelHost) IsSet() bool {
	return v.isSet
}

func (v *NullableModelHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelHost(val *ModelHost) *NullableModelHost {
	return &NullableModelHost{value: val, isSet: true}
}

func (v NullableModelHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


