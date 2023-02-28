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

// checks if the AssetHolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetHolder{}

// AssetHolder struct for AssetHolder
type AssetHolder struct {
	Assets map[string]float32 `json:"assets"`
}

// NewAssetHolder instantiates a new AssetHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetHolder(assets map[string]float32) *AssetHolder {
	this := AssetHolder{}
	this.Assets = assets
	return &this
}

// NewAssetHolderWithDefaults instantiates a new AssetHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetHolderWithDefaults() *AssetHolder {
	this := AssetHolder{}
	return &this
}

// GetAssets returns the Assets field value
func (o *AssetHolder) GetAssets() map[string]float32 {
	if o == nil {
		var ret map[string]float32
		return ret
	}

	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value
// and a boolean to check if the value has been set.
func (o *AssetHolder) GetAssetsOk() (*map[string]float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assets, true
}

// SetAssets sets field value
func (o *AssetHolder) SetAssets(v map[string]float32) {
	o.Assets = v
}

func (o AssetHolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetHolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assets"] = o.Assets
	return toSerialize, nil
}

type NullableAssetHolder struct {
	value *AssetHolder
	isSet bool
}

func (v NullableAssetHolder) Get() *AssetHolder {
	return v.value
}

func (v *NullableAssetHolder) Set(val *AssetHolder) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetHolder) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetHolder(val *AssetHolder) *NullableAssetHolder {
	return &NullableAssetHolder{value: val, isSet: true}
}

func (v NullableAssetHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


