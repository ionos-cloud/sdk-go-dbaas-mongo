/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.   MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Connection The network connection  details for your cluster.
type Connection struct {
	// The datacenter to which your cluster will be connected.
	DatacenterId *string `json:"datacenterId"`
	// The numeric LAN ID with which you connect your cluster.
	LanId *string `json:"lanId"`
	// The list of IPs for your cluster. All IPs must be in a /24 network. Note the following unavailable IP ranges: 10.233.114.0/24
	CidrList *[]string `json:"cidrList"`
}

// NewConnection instantiates a new Connection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnection(datacenterId string, lanId string, cidrList []string) *Connection {
	this := Connection{}

	this.DatacenterId = &datacenterId
	this.LanId = &lanId
	this.CidrList = &cidrList

	return &this
}

// NewConnectionWithDefaults instantiates a new Connection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWithDefaults() *Connection {
	this := Connection{}
	return &this
}

// GetDatacenterId returns the DatacenterId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Connection) GetDatacenterId() *string {
	if o == nil {
		return nil
	}

	return o.DatacenterId

}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetDatacenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DatacenterId, true
}

// SetDatacenterId sets field value
func (o *Connection) SetDatacenterId(v string) {

	o.DatacenterId = &v

}

// HasDatacenterId returns a boolean if a field has been set.
func (o *Connection) HasDatacenterId() bool {
	if o != nil && o.DatacenterId != nil {
		return true
	}

	return false
}

// GetLanId returns the LanId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Connection) GetLanId() *string {
	if o == nil {
		return nil
	}

	return o.LanId

}

// GetLanIdOk returns a tuple with the LanId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetLanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.LanId, true
}

// SetLanId sets field value
func (o *Connection) SetLanId(v string) {

	o.LanId = &v

}

// HasLanId returns a boolean if a field has been set.
func (o *Connection) HasLanId() bool {
	if o != nil && o.LanId != nil {
		return true
	}

	return false
}

// GetCidrList returns the CidrList field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *Connection) GetCidrList() *[]string {
	if o == nil {
		return nil
	}

	return o.CidrList

}

// GetCidrListOk returns a tuple with the CidrList field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetCidrListOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CidrList, true
}

// SetCidrList sets field value
func (o *Connection) SetCidrList(v []string) {

	o.CidrList = &v

}

// HasCidrList returns a boolean if a field has been set.
func (o *Connection) HasCidrList() bool {
	if o != nil && o.CidrList != nil {
		return true
	}

	return false
}

func (o Connection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatacenterId != nil {
		toSerialize["datacenterId"] = o.DatacenterId
	}

	if o.LanId != nil {
		toSerialize["lanId"] = o.LanId
	}

	if o.CidrList != nil {
		toSerialize["cidrList"] = o.CidrList
	}

	return json.Marshal(toSerialize)
}

type NullableConnection struct {
	value *Connection
	isSet bool
}

func (v NullableConnection) Get() *Connection {
	return v.value
}

func (v *NullableConnection) Set(val *Connection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnection(val *Connection) *NullableConnection {
	return &NullableConnection{value: val, isSet: true}
}

func (v NullableConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
