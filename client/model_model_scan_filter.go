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

// checks if the ModelScanFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanFilter{}

// ModelScanFilter struct for ModelScanFilter
type ModelScanFilter struct {
	CloudAccountScanFilter ReportersContainsFilter `json:"cloud_account_scan_filter"`
	ContainerScanFilter ReportersContainsFilter `json:"container_scan_filter"`
	HostScanFilter ReportersContainsFilter `json:"host_scan_filter"`
	ImageScanFilter ReportersContainsFilter `json:"image_scan_filter"`
	KubernetesClusterScanFilter ReportersContainsFilter `json:"kubernetes_cluster_scan_filter"`
}

type _ModelScanFilter ModelScanFilter

// NewModelScanFilter instantiates a new ModelScanFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanFilter(cloudAccountScanFilter ReportersContainsFilter, containerScanFilter ReportersContainsFilter, hostScanFilter ReportersContainsFilter, imageScanFilter ReportersContainsFilter, kubernetesClusterScanFilter ReportersContainsFilter) *ModelScanFilter {
	this := ModelScanFilter{}
	this.CloudAccountScanFilter = cloudAccountScanFilter
	this.ContainerScanFilter = containerScanFilter
	this.HostScanFilter = hostScanFilter
	this.ImageScanFilter = imageScanFilter
	this.KubernetesClusterScanFilter = kubernetesClusterScanFilter
	return &this
}

// NewModelScanFilterWithDefaults instantiates a new ModelScanFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanFilterWithDefaults() *ModelScanFilter {
	this := ModelScanFilter{}
	return &this
}

// GetCloudAccountScanFilter returns the CloudAccountScanFilter field value
func (o *ModelScanFilter) GetCloudAccountScanFilter() ReportersContainsFilter {
	if o == nil {
		var ret ReportersContainsFilter
		return ret
	}

	return o.CloudAccountScanFilter
}

// GetCloudAccountScanFilterOk returns a tuple with the CloudAccountScanFilter field value
// and a boolean to check if the value has been set.
func (o *ModelScanFilter) GetCloudAccountScanFilterOk() (*ReportersContainsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudAccountScanFilter, true
}

// SetCloudAccountScanFilter sets field value
func (o *ModelScanFilter) SetCloudAccountScanFilter(v ReportersContainsFilter) {
	o.CloudAccountScanFilter = v
}

// GetContainerScanFilter returns the ContainerScanFilter field value
func (o *ModelScanFilter) GetContainerScanFilter() ReportersContainsFilter {
	if o == nil {
		var ret ReportersContainsFilter
		return ret
	}

	return o.ContainerScanFilter
}

// GetContainerScanFilterOk returns a tuple with the ContainerScanFilter field value
// and a boolean to check if the value has been set.
func (o *ModelScanFilter) GetContainerScanFilterOk() (*ReportersContainsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerScanFilter, true
}

// SetContainerScanFilter sets field value
func (o *ModelScanFilter) SetContainerScanFilter(v ReportersContainsFilter) {
	o.ContainerScanFilter = v
}

// GetHostScanFilter returns the HostScanFilter field value
func (o *ModelScanFilter) GetHostScanFilter() ReportersContainsFilter {
	if o == nil {
		var ret ReportersContainsFilter
		return ret
	}

	return o.HostScanFilter
}

// GetHostScanFilterOk returns a tuple with the HostScanFilter field value
// and a boolean to check if the value has been set.
func (o *ModelScanFilter) GetHostScanFilterOk() (*ReportersContainsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostScanFilter, true
}

// SetHostScanFilter sets field value
func (o *ModelScanFilter) SetHostScanFilter(v ReportersContainsFilter) {
	o.HostScanFilter = v
}

// GetImageScanFilter returns the ImageScanFilter field value
func (o *ModelScanFilter) GetImageScanFilter() ReportersContainsFilter {
	if o == nil {
		var ret ReportersContainsFilter
		return ret
	}

	return o.ImageScanFilter
}

// GetImageScanFilterOk returns a tuple with the ImageScanFilter field value
// and a boolean to check if the value has been set.
func (o *ModelScanFilter) GetImageScanFilterOk() (*ReportersContainsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageScanFilter, true
}

// SetImageScanFilter sets field value
func (o *ModelScanFilter) SetImageScanFilter(v ReportersContainsFilter) {
	o.ImageScanFilter = v
}

// GetKubernetesClusterScanFilter returns the KubernetesClusterScanFilter field value
func (o *ModelScanFilter) GetKubernetesClusterScanFilter() ReportersContainsFilter {
	if o == nil {
		var ret ReportersContainsFilter
		return ret
	}

	return o.KubernetesClusterScanFilter
}

// GetKubernetesClusterScanFilterOk returns a tuple with the KubernetesClusterScanFilter field value
// and a boolean to check if the value has been set.
func (o *ModelScanFilter) GetKubernetesClusterScanFilterOk() (*ReportersContainsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterScanFilter, true
}

// SetKubernetesClusterScanFilter sets field value
func (o *ModelScanFilter) SetKubernetesClusterScanFilter(v ReportersContainsFilter) {
	o.KubernetesClusterScanFilter = v
}

func (o ModelScanFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_account_scan_filter"] = o.CloudAccountScanFilter
	toSerialize["container_scan_filter"] = o.ContainerScanFilter
	toSerialize["host_scan_filter"] = o.HostScanFilter
	toSerialize["image_scan_filter"] = o.ImageScanFilter
	toSerialize["kubernetes_cluster_scan_filter"] = o.KubernetesClusterScanFilter
	return toSerialize, nil
}

func (o *ModelScanFilter) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_account_scan_filter",
		"container_scan_filter",
		"host_scan_filter",
		"image_scan_filter",
		"kubernetes_cluster_scan_filter",
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

	varModelScanFilter := _ModelScanFilter{}

	err = json.Unmarshal(bytes, &varModelScanFilter)

	if err != nil {
		return err
	}

	*o = ModelScanFilter(varModelScanFilter)

	return err
}

type NullableModelScanFilter struct {
	value *ModelScanFilter
	isSet bool
}

func (v NullableModelScanFilter) Get() *ModelScanFilter {
	return v.value
}

func (v *NullableModelScanFilter) Set(val *ModelScanFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanFilter(val *ModelScanFilter) *NullableModelScanFilter {
	return &NullableModelScanFilter{value: val, isSet: true}
}

func (v NullableModelScanFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


