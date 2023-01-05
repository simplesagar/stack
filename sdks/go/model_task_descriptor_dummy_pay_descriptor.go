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

// TaskDescriptorDummyPayDescriptor struct for TaskDescriptorDummyPayDescriptor
type TaskDescriptorDummyPayDescriptor struct {
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
	FileName *string `json:"fileName,omitempty"`
}

// NewTaskDescriptorDummyPayDescriptor instantiates a new TaskDescriptorDummyPayDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskDescriptorDummyPayDescriptor() *TaskDescriptorDummyPayDescriptor {
	this := TaskDescriptorDummyPayDescriptor{}
	return &this
}

// NewTaskDescriptorDummyPayDescriptorWithDefaults instantiates a new TaskDescriptorDummyPayDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskDescriptorDummyPayDescriptorWithDefaults() *TaskDescriptorDummyPayDescriptor {
	this := TaskDescriptorDummyPayDescriptor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskDescriptorDummyPayDescriptor) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorDummyPayDescriptor) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskDescriptorDummyPayDescriptor) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskDescriptorDummyPayDescriptor) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TaskDescriptorDummyPayDescriptor) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorDummyPayDescriptor) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TaskDescriptorDummyPayDescriptor) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TaskDescriptorDummyPayDescriptor) SetKey(v string) {
	o.Key = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *TaskDescriptorDummyPayDescriptor) GetFileName() string {
	if o == nil || isNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorDummyPayDescriptor) GetFileNameOk() (*string, bool) {
	if o == nil || isNil(o.FileName) {
    return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *TaskDescriptorDummyPayDescriptor) HasFileName() bool {
	if o != nil && !isNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *TaskDescriptorDummyPayDescriptor) SetFileName(v string) {
	o.FileName = &v
}

func (o TaskDescriptorDummyPayDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	return json.Marshal(toSerialize)
}

type NullableTaskDescriptorDummyPayDescriptor struct {
	value *TaskDescriptorDummyPayDescriptor
	isSet bool
}

func (v NullableTaskDescriptorDummyPayDescriptor) Get() *TaskDescriptorDummyPayDescriptor {
	return v.value
}

func (v *NullableTaskDescriptorDummyPayDescriptor) Set(val *TaskDescriptorDummyPayDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskDescriptorDummyPayDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskDescriptorDummyPayDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskDescriptorDummyPayDescriptor(val *TaskDescriptorDummyPayDescriptor) *NullableTaskDescriptorDummyPayDescriptor {
	return &NullableTaskDescriptorDummyPayDescriptor{value: val, isSet: true}
}

func (v NullableTaskDescriptorDummyPayDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskDescriptorDummyPayDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


