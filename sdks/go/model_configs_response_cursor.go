/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.20230228
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ConfigsResponseCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigsResponseCursor{}

// ConfigsResponseCursor struct for ConfigsResponseCursor
type ConfigsResponseCursor struct {
	HasMore bool `json:"hasMore"`
	Data []WebhooksConfig `json:"data"`
}

// NewConfigsResponseCursor instantiates a new ConfigsResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigsResponseCursor(hasMore bool, data []WebhooksConfig) *ConfigsResponseCursor {
	this := ConfigsResponseCursor{}
	this.HasMore = hasMore
	this.Data = data
	return &this
}

// NewConfigsResponseCursorWithDefaults instantiates a new ConfigsResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigsResponseCursorWithDefaults() *ConfigsResponseCursor {
	this := ConfigsResponseCursor{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *ConfigsResponseCursor) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ConfigsResponseCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ConfigsResponseCursor) SetHasMore(v bool) {
	o.HasMore = v
}

// GetData returns the Data field value
func (o *ConfigsResponseCursor) GetData() []WebhooksConfig {
	if o == nil {
		var ret []WebhooksConfig
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ConfigsResponseCursor) GetDataOk() ([]WebhooksConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ConfigsResponseCursor) SetData(v []WebhooksConfig) {
	o.Data = v
}

func (o ConfigsResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigsResponseCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hasMore"] = o.HasMore
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableConfigsResponseCursor struct {
	value *ConfigsResponseCursor
	isSet bool
}

func (v NullableConfigsResponseCursor) Get() *ConfigsResponseCursor {
	return v.value
}

func (v *NullableConfigsResponseCursor) Set(val *ConfigsResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigsResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigsResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigsResponseCursor(val *ConfigsResponseCursor) *NullableConfigsResponseCursor {
	return &NullableConfigsResponseCursor{value: val, isSet: true}
}

func (v NullableConfigsResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigsResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


