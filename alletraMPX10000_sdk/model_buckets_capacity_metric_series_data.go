/*
Device Type-7 APIs

APIs to get information about the HPE Alletra Storage MP X10000 system

API version: 0.1.0
Contact: object-svc-team@hpe.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BucketsCapacityMetricSeriesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketsCapacityMetricSeriesData{}

// BucketsCapacityMetricSeriesData struct for BucketsCapacityMetricSeriesData
type BucketsCapacityMetricSeriesData struct {
	// epoch timestamp
	TimestampMs NullableInt64 `json:"timestampMs,omitempty"`
	// average logical usage value at particular timestamp
	UsageMiB NullableFloat32 `json:"usageMiB,omitempty"`
}

// NewBucketsCapacityMetricSeriesData instantiates a new BucketsCapacityMetricSeriesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketsCapacityMetricSeriesData() *BucketsCapacityMetricSeriesData {
	this := BucketsCapacityMetricSeriesData{}
	return &this
}

// NewBucketsCapacityMetricSeriesDataWithDefaults instantiates a new BucketsCapacityMetricSeriesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketsCapacityMetricSeriesDataWithDefaults() *BucketsCapacityMetricSeriesData {
	this := BucketsCapacityMetricSeriesData{}
	return &this
}

// GetTimestampMs returns the TimestampMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BucketsCapacityMetricSeriesData) GetTimestampMs() int64 {
	if o == nil || IsNil(o.TimestampMs.Get()) {
		var ret int64
		return ret
	}
	return *o.TimestampMs.Get()
}

// GetTimestampMsOk returns a tuple with the TimestampMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BucketsCapacityMetricSeriesData) GetTimestampMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimestampMs.Get(), o.TimestampMs.IsSet()
}

// HasTimestampMs returns a boolean if a field has been set.
func (o *BucketsCapacityMetricSeriesData) HasTimestampMs() bool {
	if o != nil && o.TimestampMs.IsSet() {
		return true
	}

	return false
}

// SetTimestampMs gets a reference to the given NullableInt64 and assigns it to the TimestampMs field.
func (o *BucketsCapacityMetricSeriesData) SetTimestampMs(v int64) {
	o.TimestampMs.Set(&v)
}
// SetTimestampMsNil sets the value for TimestampMs to be an explicit nil
func (o *BucketsCapacityMetricSeriesData) SetTimestampMsNil() {
	o.TimestampMs.Set(nil)
}

// UnsetTimestampMs ensures that no value is present for TimestampMs, not even an explicit nil
func (o *BucketsCapacityMetricSeriesData) UnsetTimestampMs() {
	o.TimestampMs.Unset()
}

// GetUsageMiB returns the UsageMiB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BucketsCapacityMetricSeriesData) GetUsageMiB() float32 {
	if o == nil || IsNil(o.UsageMiB.Get()) {
		var ret float32
		return ret
	}
	return *o.UsageMiB.Get()
}

// GetUsageMiBOk returns a tuple with the UsageMiB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BucketsCapacityMetricSeriesData) GetUsageMiBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageMiB.Get(), o.UsageMiB.IsSet()
}

// HasUsageMiB returns a boolean if a field has been set.
func (o *BucketsCapacityMetricSeriesData) HasUsageMiB() bool {
	if o != nil && o.UsageMiB.IsSet() {
		return true
	}

	return false
}

// SetUsageMiB gets a reference to the given NullableFloat32 and assigns it to the UsageMiB field.
func (o *BucketsCapacityMetricSeriesData) SetUsageMiB(v float32) {
	o.UsageMiB.Set(&v)
}
// SetUsageMiBNil sets the value for UsageMiB to be an explicit nil
func (o *BucketsCapacityMetricSeriesData) SetUsageMiBNil() {
	o.UsageMiB.Set(nil)
}

// UnsetUsageMiB ensures that no value is present for UsageMiB, not even an explicit nil
func (o *BucketsCapacityMetricSeriesData) UnsetUsageMiB() {
	o.UsageMiB.Unset()
}

func (o BucketsCapacityMetricSeriesData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketsCapacityMetricSeriesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TimestampMs.IsSet() {
		toSerialize["timestampMs"] = o.TimestampMs.Get()
	}
	if o.UsageMiB.IsSet() {
		toSerialize["usageMiB"] = o.UsageMiB.Get()
	}
	return toSerialize, nil
}

type NullableBucketsCapacityMetricSeriesData struct {
	value *BucketsCapacityMetricSeriesData
	isSet bool
}

func (v NullableBucketsCapacityMetricSeriesData) Get() *BucketsCapacityMetricSeriesData {
	return v.value
}

func (v *NullableBucketsCapacityMetricSeriesData) Set(val *BucketsCapacityMetricSeriesData) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketsCapacityMetricSeriesData) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketsCapacityMetricSeriesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketsCapacityMetricSeriesData(val *BucketsCapacityMetricSeriesData) *NullableBucketsCapacityMetricSeriesData {
	return &NullableBucketsCapacityMetricSeriesData{value: val, isSet: true}
}

func (v NullableBucketsCapacityMetricSeriesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketsCapacityMetricSeriesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


