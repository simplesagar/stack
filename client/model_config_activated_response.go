/*
Webhooks server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.2.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ConfigActivatedResponse struct for ConfigActivatedResponse
type ConfigActivatedResponse struct {
	Data *ConfigActivated `json:"data,omitempty"`
}

// NewConfigActivatedResponse instantiates a new ConfigActivatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigActivatedResponse() *ConfigActivatedResponse {
	this := ConfigActivatedResponse{}
	return &this
}

// NewConfigActivatedResponseWithDefaults instantiates a new ConfigActivatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigActivatedResponseWithDefaults() *ConfigActivatedResponse {
	this := ConfigActivatedResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ConfigActivatedResponse) GetData() ConfigActivated {
	if o == nil || isNil(o.Data) {
		var ret ConfigActivated
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigActivatedResponse) GetDataOk() (*ConfigActivated, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ConfigActivatedResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ConfigActivated and assigns it to the Data field.
func (o *ConfigActivatedResponse) SetData(v ConfigActivated) {
	o.Data = &v
}

func (o ConfigActivatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableConfigActivatedResponse struct {
	value *ConfigActivatedResponse
	isSet bool
}

func (v NullableConfigActivatedResponse) Get() *ConfigActivatedResponse {
	return v.value
}

func (v *NullableConfigActivatedResponse) Set(val *ConfigActivatedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigActivatedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigActivatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigActivatedResponse(val *ConfigActivatedResponse) *NullableConfigActivatedResponse {
	return &NullableConfigActivatedResponse{value: val, isSet: true}
}

func (v NullableConfigActivatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigActivatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


