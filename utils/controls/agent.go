package controls

import (
	"encoding/json"
)

type ActionID int

const (
	StartVulnerabilityScan ActionID = iota
	StartSecretScan
	StartComplianceScan
	StartMalwareScan
	StartAgentUpgrade
)

type ScanResource int

const (
	Container ScanResource = iota
	Image
	Host
	CloudAccount
	KubernetesCluster
)

func ResourceTypeToNeo4j(t ScanResource) string {
	switch t {
	case Container:
		return "Container"
	case Image:
		return "ContainerImage"
	case Host:
		return "Node"
	case CloudAccount:
		return "Node"
	case KubernetesCluster:
		return "KubernetesCluster"
	}
	return ""
}

func ResourceTypeToString(t ScanResource) string {
	switch t {
	case Container:
		return "container"
	case Image:
		return "image"
	case Host:
		return "host"
	case CloudAccount:
		return "cloud_account"
	case KubernetesCluster:
		return "kubernetes_cluster"
	}
	return ""
}

func StringToResourceType(s string) ScanResource {
	switch s {
	case "container":
		return Container
	case "image":
		return Image
	case "host":
		return Host
	}
	return -1
}

type StartVulnerabilityScanRequest struct {
	NodeId   string            `json:"node_id" required:"true"`
	NodeType ScanResource      `json:"node_type" required:"true"`
	BinArgs  map[string]string `json:"bin_args" required:"true"`
}

type StartSecretScanRequest struct {
	NodeId   string            `json:"node_id" required:"true"`
	NodeType ScanResource      `json:"node_type" required:"true"`
	BinArgs  map[string]string `json:"bin_args" required:"true"`
}

type StartComplianceScanRequest struct {
	NodeId   string            `json:"node_id" required:"true"`
	NodeType ScanResource      `json:"node_type" required:"true"`
	BinArgs  map[string]string `json:"bin_args" required:"true"`
}

type StartMalwareScanRequest struct {
	NodeId   string            `json:"node_id" required:"true"`
	NodeType ScanResource      `json:"node_type" required:"true"`
	BinArgs  map[string]string `json:"bin_args" required:"true"`
}

type StartAgentUpgradeRequest struct {
	HomeDirectoryUrl string `json:"home_directory_url" required:"true"`
	Version          string `json:"version" required:"true"`
}

type Action struct {
	ID             ActionID `json:"id" required:"true"`
	RequestPayload string   `json:"request_payload" required:"true"`
}

type AgentControls struct {
	BeatRateSec int32    `json:"beatrate" required:"true"`
	Commands    []Action `json:"commands" required:"true"`
}

func (ac AgentControls) ToBytes() ([]byte, error) {
	return json.Marshal(ac)
}
