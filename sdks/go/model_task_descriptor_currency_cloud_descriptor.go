/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: SDK_VERSION
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// TaskDescriptorCurrencyCloudDescriptor struct for TaskDescriptorCurrencyCloudDescriptor
type TaskDescriptorCurrencyCloudDescriptor struct {
	Name *string `json:"name,omitempty"`
}

// NewTaskDescriptorCurrencyCloudDescriptor instantiates a new TaskDescriptorCurrencyCloudDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskDescriptorCurrencyCloudDescriptor() *TaskDescriptorCurrencyCloudDescriptor {
	this := TaskDescriptorCurrencyCloudDescriptor{}
	return &this
}

// NewTaskDescriptorCurrencyCloudDescriptorWithDefaults instantiates a new TaskDescriptorCurrencyCloudDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskDescriptorCurrencyCloudDescriptorWithDefaults() *TaskDescriptorCurrencyCloudDescriptor {
	this := TaskDescriptorCurrencyCloudDescriptor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskDescriptorCurrencyCloudDescriptor) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorCurrencyCloudDescriptor) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskDescriptorCurrencyCloudDescriptor) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskDescriptorCurrencyCloudDescriptor) SetName(v string) {
	o.Name = &v
}

func (o TaskDescriptorCurrencyCloudDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTaskDescriptorCurrencyCloudDescriptor struct {
	value *TaskDescriptorCurrencyCloudDescriptor
	isSet bool
}

func (v NullableTaskDescriptorCurrencyCloudDescriptor) Get() *TaskDescriptorCurrencyCloudDescriptor {
	return v.value
}

func (v *NullableTaskDescriptorCurrencyCloudDescriptor) Set(val *TaskDescriptorCurrencyCloudDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskDescriptorCurrencyCloudDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskDescriptorCurrencyCloudDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskDescriptorCurrencyCloudDescriptor(val *TaskDescriptorCurrencyCloudDescriptor) *NullableTaskDescriptorCurrencyCloudDescriptor {
	return &NullableTaskDescriptorCurrencyCloudDescriptor{value: val, isSet: true}
}

func (v NullableTaskDescriptorCurrencyCloudDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskDescriptorCurrencyCloudDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


