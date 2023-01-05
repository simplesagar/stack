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

// Total struct for Total
type Total struct {
	Value *int64 `json:"value,omitempty"`
	Relation *string `json:"relation,omitempty"`
}

// NewTotal instantiates a new Total object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotal() *Total {
	this := Total{}
	return &this
}

// NewTotalWithDefaults instantiates a new Total object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotalWithDefaults() *Total {
	this := Total{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Total) GetValue() int64 {
	if o == nil || isNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Total) GetValueOk() (*int64, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Total) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *Total) SetValue(v int64) {
	o.Value = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *Total) GetRelation() string {
	if o == nil || isNil(o.Relation) {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Total) GetRelationOk() (*string, bool) {
	if o == nil || isNil(o.Relation) {
    return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *Total) HasRelation() bool {
	if o != nil && !isNil(o.Relation) {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *Total) SetRelation(v string) {
	o.Relation = &v
}

func (o Total) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Relation) {
		toSerialize["relation"] = o.Relation
	}
	return json.Marshal(toSerialize)
}

type NullableTotal struct {
	value *Total
	isSet bool
}

func (v NullableTotal) Get() *Total {
	return v.value
}

func (v *NullableTotal) Set(val *Total) {
	v.value = val
	v.isSet = true
}

func (v NullableTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotal(val *Total) *NullableTotal {
	return &NullableTotal{value: val, isSet: true}
}

func (v NullableTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


