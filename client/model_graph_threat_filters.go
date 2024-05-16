/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GraphThreatFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphThreatFilters{}

// GraphThreatFilters struct for GraphThreatFilters
type GraphThreatFilters struct {
	AwsFilter GraphCloudProviderFilter `json:"aws_filter"`
	AzureFilter GraphCloudProviderFilter `json:"azure_filter"`
	CloudResourceOnly bool `json:"cloud_resource_only"`
	GcpFilter GraphCloudProviderFilter `json:"gcp_filter"`
	Type string `json:"type"`
}

type _GraphThreatFilters GraphThreatFilters

// NewGraphThreatFilters instantiates a new GraphThreatFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphThreatFilters(awsFilter GraphCloudProviderFilter, azureFilter GraphCloudProviderFilter, cloudResourceOnly bool, gcpFilter GraphCloudProviderFilter, type_ string) *GraphThreatFilters {
	this := GraphThreatFilters{}
	this.AwsFilter = awsFilter
	this.AzureFilter = azureFilter
	this.CloudResourceOnly = cloudResourceOnly
	this.GcpFilter = gcpFilter
	this.Type = type_
	return &this
}

// NewGraphThreatFiltersWithDefaults instantiates a new GraphThreatFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphThreatFiltersWithDefaults() *GraphThreatFilters {
	this := GraphThreatFilters{}
	return &this
}

// GetAwsFilter returns the AwsFilter field value
func (o *GraphThreatFilters) GetAwsFilter() GraphCloudProviderFilter {
	if o == nil {
		var ret GraphCloudProviderFilter
		return ret
	}

	return o.AwsFilter
}

// GetAwsFilterOk returns a tuple with the AwsFilter field value
// and a boolean to check if the value has been set.
func (o *GraphThreatFilters) GetAwsFilterOk() (*GraphCloudProviderFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsFilter, true
}

// SetAwsFilter sets field value
func (o *GraphThreatFilters) SetAwsFilter(v GraphCloudProviderFilter) {
	o.AwsFilter = v
}

// GetAzureFilter returns the AzureFilter field value
func (o *GraphThreatFilters) GetAzureFilter() GraphCloudProviderFilter {
	if o == nil {
		var ret GraphCloudProviderFilter
		return ret
	}

	return o.AzureFilter
}

// GetAzureFilterOk returns a tuple with the AzureFilter field value
// and a boolean to check if the value has been set.
func (o *GraphThreatFilters) GetAzureFilterOk() (*GraphCloudProviderFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureFilter, true
}

// SetAzureFilter sets field value
func (o *GraphThreatFilters) SetAzureFilter(v GraphCloudProviderFilter) {
	o.AzureFilter = v
}

// GetCloudResourceOnly returns the CloudResourceOnly field value
func (o *GraphThreatFilters) GetCloudResourceOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CloudResourceOnly
}

// GetCloudResourceOnlyOk returns a tuple with the CloudResourceOnly field value
// and a boolean to check if the value has been set.
func (o *GraphThreatFilters) GetCloudResourceOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudResourceOnly, true
}

// SetCloudResourceOnly sets field value
func (o *GraphThreatFilters) SetCloudResourceOnly(v bool) {
	o.CloudResourceOnly = v
}

// GetGcpFilter returns the GcpFilter field value
func (o *GraphThreatFilters) GetGcpFilter() GraphCloudProviderFilter {
	if o == nil {
		var ret GraphCloudProviderFilter
		return ret
	}

	return o.GcpFilter
}

// GetGcpFilterOk returns a tuple with the GcpFilter field value
// and a boolean to check if the value has been set.
func (o *GraphThreatFilters) GetGcpFilterOk() (*GraphCloudProviderFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpFilter, true
}

// SetGcpFilter sets field value
func (o *GraphThreatFilters) SetGcpFilter(v GraphCloudProviderFilter) {
	o.GcpFilter = v
}

// GetType returns the Type field value
func (o *GraphThreatFilters) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GraphThreatFilters) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GraphThreatFilters) SetType(v string) {
	o.Type = v
}

func (o GraphThreatFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphThreatFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws_filter"] = o.AwsFilter
	toSerialize["azure_filter"] = o.AzureFilter
	toSerialize["cloud_resource_only"] = o.CloudResourceOnly
	toSerialize["gcp_filter"] = o.GcpFilter
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *GraphThreatFilters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aws_filter",
		"azure_filter",
		"cloud_resource_only",
		"gcp_filter",
		"type",
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

	varGraphThreatFilters := _GraphThreatFilters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGraphThreatFilters)

	if err != nil {
		return err
	}

	*o = GraphThreatFilters(varGraphThreatFilters)

	return err
}

type NullableGraphThreatFilters struct {
	value *GraphThreatFilters
	isSet bool
}

func (v NullableGraphThreatFilters) Get() *GraphThreatFilters {
	return v.value
}

func (v *NullableGraphThreatFilters) Set(val *GraphThreatFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphThreatFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphThreatFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphThreatFilters(val *GraphThreatFilters) *NullableGraphThreatFilters {
	return &NullableGraphThreatFilters{value: val, isSet: true}
}

func (v NullableGraphThreatFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphThreatFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


