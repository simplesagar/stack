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
	"time"
)

// checks if the TaskBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskBase{}

// TaskBase struct for TaskBase
type TaskBase struct {
	Id string `json:"id"`
	ConnectorId string `json:"connectorId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status PaymentStatus `json:"status"`
	State map[string]interface{} `json:"state"`
	Error *string `json:"error,omitempty"`
}

// NewTaskBase instantiates a new TaskBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskBase(id string, connectorId string, createdAt time.Time, updatedAt time.Time, status PaymentStatus, state map[string]interface{}) *TaskBase {
	this := TaskBase{}
	this.Id = id
	this.ConnectorId = connectorId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Status = status
	this.State = state
	return &this
}

// NewTaskBaseWithDefaults instantiates a new TaskBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskBaseWithDefaults() *TaskBase {
	this := TaskBase{}
	return &this
}

// GetId returns the Id field value
func (o *TaskBase) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskBase) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskBase) SetId(v string) {
	o.Id = v
}

// GetConnectorId returns the ConnectorId field value
func (o *TaskBase) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *TaskBase) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *TaskBase) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TaskBase) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskBase) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TaskBase) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TaskBase) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskBase) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TaskBase) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStatus returns the Status field value
func (o *TaskBase) GetStatus() PaymentStatus {
	if o == nil {
		var ret PaymentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TaskBase) GetStatusOk() (*PaymentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TaskBase) SetStatus(v PaymentStatus) {
	o.Status = v
}

// GetState returns the State field value
func (o *TaskBase) GetState() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TaskBase) GetStateOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.State, true
}

// SetState sets field value
func (o *TaskBase) SetState(v map[string]interface{}) {
	o.State = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *TaskBase) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBase) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TaskBase) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *TaskBase) SetError(v string) {
	o.Error = &v
}

func (o TaskBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["connectorId"] = o.ConnectorId
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["status"] = o.Status
	toSerialize["state"] = o.State
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableTaskBase struct {
	value *TaskBase
	isSet bool
}

func (v NullableTaskBase) Get() *TaskBase {
	return v.value
}

func (v *NullableTaskBase) Set(val *TaskBase) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskBase) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskBase(val *TaskBase) *NullableTaskBase {
	return &NullableTaskBase{value: val, isSet: true}
}

func (v NullableTaskBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


