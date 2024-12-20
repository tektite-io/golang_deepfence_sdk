/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.2
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelCloudResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudResource{}

// ModelCloudResource struct for ModelCloudResource
type ModelCloudResource struct {
	AccountId string `json:"account_id"`
	CloudComplianceLatestScanId string `json:"cloud_compliance_latest_scan_id"`
	CloudComplianceScanStatus string `json:"cloud_compliance_scan_status"`
	CloudCompliancesCount int32 `json:"cloud_compliances_count"`
	CloudProvider string `json:"cloud_provider"`
	CloudRegion string `json:"cloud_region"`
	CloudWarnAlarmCount int32 `json:"cloud_warn_alarm_count"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	NodeType string `json:"node_type"`
	TypeLabel string `json:"type_label"`
}

type _ModelCloudResource ModelCloudResource

// NewModelCloudResource instantiates a new ModelCloudResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudResource(accountId string, cloudComplianceLatestScanId string, cloudComplianceScanStatus string, cloudCompliancesCount int32, cloudProvider string, cloudRegion string, cloudWarnAlarmCount int32, nodeId string, nodeName string, nodeType string, typeLabel string) *ModelCloudResource {
	this := ModelCloudResource{}
	this.AccountId = accountId
	this.CloudComplianceLatestScanId = cloudComplianceLatestScanId
	this.CloudComplianceScanStatus = cloudComplianceScanStatus
	this.CloudCompliancesCount = cloudCompliancesCount
	this.CloudProvider = cloudProvider
	this.CloudRegion = cloudRegion
	this.CloudWarnAlarmCount = cloudWarnAlarmCount
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.NodeType = nodeType
	this.TypeLabel = typeLabel
	return &this
}

// NewModelCloudResourceWithDefaults instantiates a new ModelCloudResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudResourceWithDefaults() *ModelCloudResource {
	this := ModelCloudResource{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *ModelCloudResource) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ModelCloudResource) SetAccountId(v string) {
	o.AccountId = v
}

// GetCloudComplianceLatestScanId returns the CloudComplianceLatestScanId field value
func (o *ModelCloudResource) GetCloudComplianceLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudComplianceLatestScanId
}

// GetCloudComplianceLatestScanIdOk returns a tuple with the CloudComplianceLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetCloudComplianceLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudComplianceLatestScanId, true
}

// SetCloudComplianceLatestScanId sets field value
func (o *ModelCloudResource) SetCloudComplianceLatestScanId(v string) {
	o.CloudComplianceLatestScanId = v
}

// GetCloudComplianceScanStatus returns the CloudComplianceScanStatus field value
func (o *ModelCloudResource) GetCloudComplianceScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudComplianceScanStatus
}

// GetCloudComplianceScanStatusOk returns a tuple with the CloudComplianceScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetCloudComplianceScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudComplianceScanStatus, true
}

// SetCloudComplianceScanStatus sets field value
func (o *ModelCloudResource) SetCloudComplianceScanStatus(v string) {
	o.CloudComplianceScanStatus = v
}

// GetCloudCompliancesCount returns the CloudCompliancesCount field value
func (o *ModelCloudResource) GetCloudCompliancesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudCompliancesCount
}

// GetCloudCompliancesCountOk returns a tuple with the CloudCompliancesCount field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetCloudCompliancesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudCompliancesCount, true
}

// SetCloudCompliancesCount sets field value
func (o *ModelCloudResource) SetCloudCompliancesCount(v int32) {
	o.CloudCompliancesCount = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *ModelCloudResource) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ModelCloudResource) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCloudRegion returns the CloudRegion field value
func (o *ModelCloudResource) GetCloudRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudRegion
}

// GetCloudRegionOk returns a tuple with the CloudRegion field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetCloudRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudRegion, true
}

// SetCloudRegion sets field value
func (o *ModelCloudResource) SetCloudRegion(v string) {
	o.CloudRegion = v
}

// GetCloudWarnAlarmCount returns the CloudWarnAlarmCount field value
func (o *ModelCloudResource) GetCloudWarnAlarmCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudWarnAlarmCount
}

// GetCloudWarnAlarmCountOk returns a tuple with the CloudWarnAlarmCount field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetCloudWarnAlarmCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudWarnAlarmCount, true
}

// SetCloudWarnAlarmCount sets field value
func (o *ModelCloudResource) SetCloudWarnAlarmCount(v int32) {
	o.CloudWarnAlarmCount = v
}

// GetNodeId returns the NodeId field value
func (o *ModelCloudResource) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelCloudResource) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelCloudResource) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelCloudResource) SetNodeName(v string) {
	o.NodeName = v
}

// GetNodeType returns the NodeType field value
func (o *ModelCloudResource) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelCloudResource) SetNodeType(v string) {
	o.NodeType = v
}

// GetTypeLabel returns the TypeLabel field value
func (o *ModelCloudResource) GetTypeLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeLabel
}

// GetTypeLabelOk returns a tuple with the TypeLabel field value
// and a boolean to check if the value has been set.
func (o *ModelCloudResource) GetTypeLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeLabel, true
}

// SetTypeLabel sets field value
func (o *ModelCloudResource) SetTypeLabel(v string) {
	o.TypeLabel = v
}

func (o ModelCloudResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["cloud_compliance_latest_scan_id"] = o.CloudComplianceLatestScanId
	toSerialize["cloud_compliance_scan_status"] = o.CloudComplianceScanStatus
	toSerialize["cloud_compliances_count"] = o.CloudCompliancesCount
	toSerialize["cloud_provider"] = o.CloudProvider
	toSerialize["cloud_region"] = o.CloudRegion
	toSerialize["cloud_warn_alarm_count"] = o.CloudWarnAlarmCount
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["node_type"] = o.NodeType
	toSerialize["type_label"] = o.TypeLabel
	return toSerialize, nil
}

func (o *ModelCloudResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"cloud_compliance_latest_scan_id",
		"cloud_compliance_scan_status",
		"cloud_compliances_count",
		"cloud_provider",
		"cloud_region",
		"cloud_warn_alarm_count",
		"node_id",
		"node_name",
		"node_type",
		"type_label",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelCloudResource := _ModelCloudResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelCloudResource)

	if err != nil {
		return err
	}

	*o = ModelCloudResource(varModelCloudResource)

	return err
}

type NullableModelCloudResource struct {
	value *ModelCloudResource
	isSet bool
}

func (v NullableModelCloudResource) Get() *ModelCloudResource {
	return v.value
}

func (v *NullableModelCloudResource) Set(val *ModelCloudResource) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudResource) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudResource(val *ModelCloudResource) *NullableModelCloudResource {
	return &NullableModelCloudResource{value: val, isSet: true}
}

func (v NullableModelCloudResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


