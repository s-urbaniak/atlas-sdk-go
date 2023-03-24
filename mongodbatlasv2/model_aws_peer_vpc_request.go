/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the AWSPeerVpcRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSPeerVpcRequest{}

// AWSPeerVpcRequest struct for AWSPeerVpcRequest
type AWSPeerVpcRequest struct {
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection.
	ContainerId string `json:"containerId"`
	// Cloud service provider that determines the needed settings to configure the network connection for a virtual private connection.
	ProviderName string `json:"providerName"`
	// Amazon Web Services (AWS) region where the Virtual Peering Connection (VPC) that you peered with the MongoDB Cloud VPC resides. The resource returns `null` if your VPC and the MongoDB Cloud VPC reside in the same region.
	AccepterRegionName string `json:"accepterRegionName"`
	// Unique twelve-digit string that identifies the Amazon Web Services (AWS) account that owns the VPC that you peered with the MongoDB Cloud VPC.
	AwsAccountId string `json:"awsAccountId"`
	// Unique string that identifies the peering connection on AWS.
	ConnectionId *string `json:"connectionId,omitempty"`
	// Type of error that can be returned when requesting an Amazon Web Services (AWS) peering connection. The resource returns `null` if the request succeeded.
	ErrorStateName *string `json:"errorStateName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the network peering connection.
	Id *string `json:"id,omitempty"`
	// Internet Protocol (IP) addresses expressed in Classless Inter-Domain Routing (CIDR) notation of the VPC's subnet that you want to peer with the MongoDB Cloud VPC.
	RouteTableCidrBlock string `json:"routeTableCidrBlock"`
	// State of the network peering connection at the time you made the request.
	StatusName *string `json:"statusName,omitempty"`
	// Unique string that identifies the VPC on Amazon Web Services (AWS) that you want to peer with the MongoDB Cloud VPC.
	VpcId string `json:"vpcId"`
}

// NewAWSPeerVpcRequest instantiates a new AWSPeerVpcRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSPeerVpcRequest(containerId string, providerName string, accepterRegionName string, awsAccountId string, routeTableCidrBlock string, vpcId string) *AWSPeerVpcRequest {
	this := AWSPeerVpcRequest{}
	this.ContainerId = containerId
	this.ProviderName = providerName
	this.AccepterRegionName = accepterRegionName
	this.AwsAccountId = awsAccountId
	this.RouteTableCidrBlock = routeTableCidrBlock
	this.VpcId = vpcId
	return &this
}

// NewAWSPeerVpcRequestWithDefaults instantiates a new AWSPeerVpcRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSPeerVpcRequestWithDefaults() *AWSPeerVpcRequest {
	this := AWSPeerVpcRequest{}
	return &this
}

// GetContainerId returns the ContainerId field value
func (o *AWSPeerVpcRequest) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *AWSPeerVpcRequest) SetContainerId(v string) {
	o.ContainerId = v
}

// GetProviderName returns the ProviderName field value
func (o *AWSPeerVpcRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AWSPeerVpcRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetAccepterRegionName returns the AccepterRegionName field value
func (o *AWSPeerVpcRequest) GetAccepterRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccepterRegionName
}

// GetAccepterRegionNameOk returns a tuple with the AccepterRegionName field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetAccepterRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccepterRegionName, true
}

// SetAccepterRegionName sets field value
func (o *AWSPeerVpcRequest) SetAccepterRegionName(v string) {
	o.AccepterRegionName = v
}

// GetAwsAccountId returns the AwsAccountId field value
func (o *AWSPeerVpcRequest) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value
func (o *AWSPeerVpcRequest) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *AWSPeerVpcRequest) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *AWSPeerVpcRequest) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *AWSPeerVpcRequest) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetErrorStateName returns the ErrorStateName field value if set, zero value otherwise.
func (o *AWSPeerVpcRequest) GetErrorStateName() string {
	if o == nil || IsNil(o.ErrorStateName) {
		var ret string
		return ret
	}
	return *o.ErrorStateName
}

// GetErrorStateNameOk returns a tuple with the ErrorStateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetErrorStateNameOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorStateName) {
		return nil, false
	}
	return o.ErrorStateName, true
}

// HasErrorStateName returns a boolean if a field has been set.
func (o *AWSPeerVpcRequest) HasErrorStateName() bool {
	if o != nil && !IsNil(o.ErrorStateName) {
		return true
	}

	return false
}

// SetErrorStateName gets a reference to the given string and assigns it to the ErrorStateName field.
func (o *AWSPeerVpcRequest) SetErrorStateName(v string) {
	o.ErrorStateName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AWSPeerVpcRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AWSPeerVpcRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AWSPeerVpcRequest) SetId(v string) {
	o.Id = &v
}

// GetRouteTableCidrBlock returns the RouteTableCidrBlock field value
func (o *AWSPeerVpcRequest) GetRouteTableCidrBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RouteTableCidrBlock
}

// GetRouteTableCidrBlockOk returns a tuple with the RouteTableCidrBlock field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetRouteTableCidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteTableCidrBlock, true
}

// SetRouteTableCidrBlock sets field value
func (o *AWSPeerVpcRequest) SetRouteTableCidrBlock(v string) {
	o.RouteTableCidrBlock = v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *AWSPeerVpcRequest) GetStatusName() string {
	if o == nil || IsNil(o.StatusName) {
		var ret string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetStatusNameOk() (*string, bool) {
	if o == nil || IsNil(o.StatusName) {
		return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *AWSPeerVpcRequest) HasStatusName() bool {
	if o != nil && !IsNil(o.StatusName) {
		return true
	}

	return false
}

// SetStatusName gets a reference to the given string and assigns it to the StatusName field.
func (o *AWSPeerVpcRequest) SetStatusName(v string) {
	o.StatusName = &v
}

// GetVpcId returns the VpcId field value
func (o *AWSPeerVpcRequest) GetVpcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpcRequest) GetVpcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpcId, true
}

// SetVpcId sets field value
func (o *AWSPeerVpcRequest) SetVpcId(v string) {
	o.VpcId = v
}

func (o AWSPeerVpcRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AWSPeerVpcRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containerId"] = o.ContainerId
	toSerialize["providerName"] = o.ProviderName
	toSerialize["accepterRegionName"] = o.AccepterRegionName
	toSerialize["awsAccountId"] = o.AwsAccountId
	toSerialize["routeTableCidrBlock"] = o.RouteTableCidrBlock
	toSerialize["vpcId"] = o.VpcId
	return toSerialize, nil
}

type NullableAWSPeerVpcRequest struct {
	value *AWSPeerVpcRequest
	isSet bool
}

func (v NullableAWSPeerVpcRequest) Get() *AWSPeerVpcRequest {
	return v.value
}

func (v *NullableAWSPeerVpcRequest) Set(val *AWSPeerVpcRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSPeerVpcRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSPeerVpcRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSPeerVpcRequest(val *AWSPeerVpcRequest) *NullableAWSPeerVpcRequest {
	return &NullableAWSPeerVpcRequest{value: val, isSet: true}
}

func (v NullableAWSPeerVpcRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSPeerVpcRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


