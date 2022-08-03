/*
Auth API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: AUTH_VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authclient

import (
	"encoding/json"
)

// ScopeOptions struct for ScopeOptions
type ScopeOptions struct {
	Label string `json:"label"`
}

// NewScopeOptions instantiates a new ScopeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeOptions(label string) *ScopeOptions {
	this := ScopeOptions{}
	this.Label = label
	return &this
}

// NewScopeOptionsWithDefaults instantiates a new ScopeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeOptionsWithDefaults() *ScopeOptions {
	this := ScopeOptions{}
	return &this
}

// GetLabel returns the Label field value
func (o *ScopeOptions) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ScopeOptions) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ScopeOptions) SetLabel(v string) {
	o.Label = v
}

func (o ScopeOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableScopeOptions struct {
	value *ScopeOptions
	isSet bool
}

func (v NullableScopeOptions) Get() *ScopeOptions {
	return v.value
}

func (v *NullableScopeOptions) Set(val *ScopeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeOptions(val *ScopeOptions) *NullableScopeOptions {
	return &NullableScopeOptions{value: val, isSet: true}
}

func (v NullableScopeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

