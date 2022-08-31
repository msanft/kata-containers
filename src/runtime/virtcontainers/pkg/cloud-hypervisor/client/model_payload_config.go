/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PayloadConfig Payloads to boot in guest
type PayloadConfig struct {
	Kernel    *string `json:"kernel,omitempty"`
	Cmdline   *string `json:"cmdline,omitempty"`
	Initramfs *string `json:"initramfs,omitempty"`
}

// NewPayloadConfig instantiates a new PayloadConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayloadConfig() *PayloadConfig {
	this := PayloadConfig{}
	return &this
}

// NewPayloadConfigWithDefaults instantiates a new PayloadConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadConfigWithDefaults() *PayloadConfig {
	this := PayloadConfig{}
	return &this
}

// GetKernel returns the Kernel field value if set, zero value otherwise.
func (o *PayloadConfig) GetKernel() string {
	if o == nil || o.Kernel == nil {
		var ret string
		return ret
	}
	return *o.Kernel
}

// GetKernelOk returns a tuple with the Kernel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadConfig) GetKernelOk() (*string, bool) {
	if o == nil || o.Kernel == nil {
		return nil, false
	}
	return o.Kernel, true
}

// HasKernel returns a boolean if a field has been set.
func (o *PayloadConfig) HasKernel() bool {
	if o != nil && o.Kernel != nil {
		return true
	}

	return false
}

// SetKernel gets a reference to the given string and assigns it to the Kernel field.
func (o *PayloadConfig) SetKernel(v string) {
	o.Kernel = &v
}

// GetCmdline returns the Cmdline field value if set, zero value otherwise.
func (o *PayloadConfig) GetCmdline() string {
	if o == nil || o.Cmdline == nil {
		var ret string
		return ret
	}
	return *o.Cmdline
}

// GetCmdlineOk returns a tuple with the Cmdline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadConfig) GetCmdlineOk() (*string, bool) {
	if o == nil || o.Cmdline == nil {
		return nil, false
	}
	return o.Cmdline, true
}

// HasCmdline returns a boolean if a field has been set.
func (o *PayloadConfig) HasCmdline() bool {
	if o != nil && o.Cmdline != nil {
		return true
	}

	return false
}

// SetCmdline gets a reference to the given string and assigns it to the Cmdline field.
func (o *PayloadConfig) SetCmdline(v string) {
	o.Cmdline = &v
}

// GetInitramfs returns the Initramfs field value if set, zero value otherwise.
func (o *PayloadConfig) GetInitramfs() string {
	if o == nil || o.Initramfs == nil {
		var ret string
		return ret
	}
	return *o.Initramfs
}

// GetInitramfsOk returns a tuple with the Initramfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadConfig) GetInitramfsOk() (*string, bool) {
	if o == nil || o.Initramfs == nil {
		return nil, false
	}
	return o.Initramfs, true
}

// HasInitramfs returns a boolean if a field has been set.
func (o *PayloadConfig) HasInitramfs() bool {
	if o != nil && o.Initramfs != nil {
		return true
	}

	return false
}

// SetInitramfs gets a reference to the given string and assigns it to the Initramfs field.
func (o *PayloadConfig) SetInitramfs(v string) {
	o.Initramfs = &v
}

func (o PayloadConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kernel != nil {
		toSerialize["kernel"] = o.Kernel
	}
	if o.Cmdline != nil {
		toSerialize["cmdline"] = o.Cmdline
	}
	if o.Initramfs != nil {
		toSerialize["initramfs"] = o.Initramfs
	}
	return json.Marshal(toSerialize)
}

type NullablePayloadConfig struct {
	value *PayloadConfig
	isSet bool
}

func (v NullablePayloadConfig) Get() *PayloadConfig {
	return v.value
}

func (v *NullablePayloadConfig) Set(val *PayloadConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePayloadConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePayloadConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayloadConfig(val *PayloadConfig) *NullablePayloadConfig {
	return &NullablePayloadConfig{value: val, isSet: true}
}

func (v NullablePayloadConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayloadConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}