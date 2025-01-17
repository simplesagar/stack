/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the WiseConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WiseConfig{}

// WiseConfig struct for WiseConfig
type WiseConfig struct {
	ApiKey string `json:"apiKey"`
}

// NewWiseConfig instantiates a new WiseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWiseConfig(apiKey string) *WiseConfig {
	this := WiseConfig{}
	this.ApiKey = apiKey
	return &this
}

// NewWiseConfigWithDefaults instantiates a new WiseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWiseConfigWithDefaults() *WiseConfig {
	this := WiseConfig{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *WiseConfig) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *WiseConfig) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *WiseConfig) SetApiKey(v string) {
	o.ApiKey = v
}

func (o WiseConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WiseConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiKey"] = o.ApiKey
	return toSerialize, nil
}

type NullableWiseConfig struct {
	value *WiseConfig
	isSet bool
}

func (v NullableWiseConfig) Get() *WiseConfig {
	return v.value
}

func (v *NullableWiseConfig) Set(val *WiseConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableWiseConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableWiseConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWiseConfig(val *WiseConfig) *NullableWiseConfig {
	return &NullableWiseConfig{value: val, isSet: true}
}

func (v NullableWiseConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWiseConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


