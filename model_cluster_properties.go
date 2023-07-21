/*
IONOS DBaaS MongoDB REST API

With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.  MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ClusterProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterProperties{}

// ClusterProperties Properties of a database cluster.
type ClusterProperties struct {
	// The name of your cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// The MongoDB version of your cluster.
	MongoDBVersion *string `json:"mongoDBVersion,omitempty"`
	// The physical location where the cluster will be created. This is the location where all your instances will be located. This property is immutable. 
	Location *string `json:"location,omitempty"`
	// The total number of instances in the cluster (one primary and n-1 secondaries). 
	Instances *int32 `json:"instances,omitempty"`
	Connections []Connection `json:"connections,omitempty"`
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`
	// The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the changes use the /templates endpoint. 
	TemplateID *string `json:"templateID,omitempty"`
	// The connection string for your cluster.
	ConnectionString *string `json:"connectionString,omitempty"`
}

// NewClusterProperties instantiates a new ClusterProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterProperties() *ClusterProperties {
	this := ClusterProperties{}
	return &this
}

// NewClusterPropertiesWithDefaults instantiates a new ClusterProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterPropertiesWithDefaults() *ClusterProperties {
	this := ClusterProperties{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ClusterProperties) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterProperties) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ClusterProperties) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetMongoDBVersion returns the MongoDBVersion field value if set, zero value otherwise.
func (o *ClusterProperties) GetMongoDBVersion() string {
	if o == nil || IsNil(o.MongoDBVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBVersion
}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetMongoDBVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBVersion) {
		return nil, false
	}
	return o.MongoDBVersion, true
}

// HasMongoDBVersion returns a boolean if a field has been set.
func (o *ClusterProperties) HasMongoDBVersion() bool {
	if o != nil && !IsNil(o.MongoDBVersion) {
		return true
	}

	return false
}

// SetMongoDBVersion gets a reference to the given string and assigns it to the MongoDBVersion field.
func (o *ClusterProperties) SetMongoDBVersion(v string) {
	o.MongoDBVersion = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ClusterProperties) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ClusterProperties) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ClusterProperties) SetLocation(v string) {
	o.Location = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *ClusterProperties) GetInstances() int32 {
	if o == nil || IsNil(o.Instances) {
		var ret int32
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetInstancesOk() (*int32, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *ClusterProperties) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given int32 and assigns it to the Instances field.
func (o *ClusterProperties) SetInstances(v int32) {
	o.Instances = &v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *ClusterProperties) GetConnections() []Connection {
	if o == nil || IsNil(o.Connections) {
		var ret []Connection
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetConnectionsOk() ([]Connection, bool) {
	if o == nil || IsNil(o.Connections) {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *ClusterProperties) HasConnections() bool {
	if o != nil && !IsNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []Connection and assigns it to the Connections field.
func (o *ClusterProperties) SetConnections(v []Connection) {
	o.Connections = v
}

// GetMaintenanceWindow returns the MaintenanceWindow field value if set, zero value otherwise.
func (o *ClusterProperties) GetMaintenanceWindow() MaintenanceWindow {
	if o == nil || IsNil(o.MaintenanceWindow) {
		var ret MaintenanceWindow
		return ret
	}
	return *o.MaintenanceWindow
}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool) {
	if o == nil || IsNil(o.MaintenanceWindow) {
		return nil, false
	}
	return o.MaintenanceWindow, true
}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *ClusterProperties) HasMaintenanceWindow() bool {
	if o != nil && !IsNil(o.MaintenanceWindow) {
		return true
	}

	return false
}

// SetMaintenanceWindow gets a reference to the given MaintenanceWindow and assigns it to the MaintenanceWindow field.
func (o *ClusterProperties) SetMaintenanceWindow(v MaintenanceWindow) {
	o.MaintenanceWindow = &v
}

// GetTemplateID returns the TemplateID field value if set, zero value otherwise.
func (o *ClusterProperties) GetTemplateID() string {
	if o == nil || IsNil(o.TemplateID) {
		var ret string
		return ret
	}
	return *o.TemplateID
}

// GetTemplateIDOk returns a tuple with the TemplateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetTemplateIDOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateID) {
		return nil, false
	}
	return o.TemplateID, true
}

// HasTemplateID returns a boolean if a field has been set.
func (o *ClusterProperties) HasTemplateID() bool {
	if o != nil && !IsNil(o.TemplateID) {
		return true
	}

	return false
}

// SetTemplateID gets a reference to the given string and assigns it to the TemplateID field.
func (o *ClusterProperties) SetTemplateID(v string) {
	o.TemplateID = &v
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise.
func (o *ClusterProperties) GetConnectionString() string {
	if o == nil || IsNil(o.ConnectionString) {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetConnectionStringOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionString) {
		return nil, false
	}
	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *ClusterProperties) HasConnectionString() bool {
	if o != nil && !IsNil(o.ConnectionString) {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *ClusterProperties) SetConnectionString(v string) {
	o.ConnectionString = &v
}

func (o ClusterProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.MongoDBVersion) {
		toSerialize["mongoDBVersion"] = o.MongoDBVersion
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !IsNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	if !IsNil(o.MaintenanceWindow) {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}
	if !IsNil(o.TemplateID) {
		toSerialize["templateID"] = o.TemplateID
	}
	if !IsNil(o.ConnectionString) {
		toSerialize["connectionString"] = o.ConnectionString
	}
	return toSerialize, nil
}

type NullableClusterProperties struct {
	value *ClusterProperties
	isSet bool
}

func (v NullableClusterProperties) Get() *ClusterProperties {
	return v.value
}

func (v *NullableClusterProperties) Set(val *ClusterProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterProperties(val *ClusterProperties) *NullableClusterProperties {
	return &NullableClusterProperties{value: val, isSet: true}
}

func (v NullableClusterProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


