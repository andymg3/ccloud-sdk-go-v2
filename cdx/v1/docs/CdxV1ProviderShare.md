# CdxV1ProviderShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ConsumerUserName** | Pointer to **string** | Name of the consumer | [optional] 
**ConsumerOrganizationName** | Pointer to **string** | Consumer organization name | [optional] 
**ProviderUserName** | Pointer to **string** | Name of the provider | [optional] 
**Status** | Pointer to **string** | Status of share | [optional] 
**DeliveryMethod** | Pointer to **string** | Method by which the invite will be delivered | [optional] 
**InvitedAt** | Pointer to **time.Time** | The date and time at which consumer was invited | [optional] 
**InviteExpiresAt** | Pointer to **time.Time** | The date and time at which the invitation will expire. Only for invited shares | [optional] 
**RedeemedAt** | Pointer to **time.Time** | The date and time at which the invite was redeemed | [optional] 
**SharedResource** | Pointer to [**ObjectReference**](ObjectReference.md) | The shared_resource to which this belongs. | [optional] 
**ServiceAccount** | Pointer to [**ObjectReference**](ObjectReference.md) | The service_account associated with this object. | [optional] 

## Methods

### NewCdxV1ProviderShare

`func NewCdxV1ProviderShare() *CdxV1ProviderShare`

NewCdxV1ProviderShare instantiates a new CdxV1ProviderShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1ProviderShareWithDefaults

`func NewCdxV1ProviderShareWithDefaults() *CdxV1ProviderShare`

NewCdxV1ProviderShareWithDefaults instantiates a new CdxV1ProviderShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1ProviderShare) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1ProviderShare) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1ProviderShare) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1ProviderShare) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1ProviderShare) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1ProviderShare) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1ProviderShare) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1ProviderShare) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1ProviderShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1ProviderShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1ProviderShare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1ProviderShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1ProviderShare) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1ProviderShare) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1ProviderShare) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1ProviderShare) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConsumerUserName

`func (o *CdxV1ProviderShare) GetConsumerUserName() string`

GetConsumerUserName returns the ConsumerUserName field if non-nil, zero value otherwise.

### GetConsumerUserNameOk

`func (o *CdxV1ProviderShare) GetConsumerUserNameOk() (*string, bool)`

GetConsumerUserNameOk returns a tuple with the ConsumerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUserName

`func (o *CdxV1ProviderShare) SetConsumerUserName(v string)`

SetConsumerUserName sets ConsumerUserName field to given value.

### HasConsumerUserName

`func (o *CdxV1ProviderShare) HasConsumerUserName() bool`

HasConsumerUserName returns a boolean if a field has been set.

### GetConsumerOrganizationName

`func (o *CdxV1ProviderShare) GetConsumerOrganizationName() string`

GetConsumerOrganizationName returns the ConsumerOrganizationName field if non-nil, zero value otherwise.

### GetConsumerOrganizationNameOk

`func (o *CdxV1ProviderShare) GetConsumerOrganizationNameOk() (*string, bool)`

GetConsumerOrganizationNameOk returns a tuple with the ConsumerOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerOrganizationName

`func (o *CdxV1ProviderShare) SetConsumerOrganizationName(v string)`

SetConsumerOrganizationName sets ConsumerOrganizationName field to given value.

### HasConsumerOrganizationName

`func (o *CdxV1ProviderShare) HasConsumerOrganizationName() bool`

HasConsumerOrganizationName returns a boolean if a field has been set.

### GetProviderUserName

`func (o *CdxV1ProviderShare) GetProviderUserName() string`

GetProviderUserName returns the ProviderUserName field if non-nil, zero value otherwise.

### GetProviderUserNameOk

`func (o *CdxV1ProviderShare) GetProviderUserNameOk() (*string, bool)`

GetProviderUserNameOk returns a tuple with the ProviderUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUserName

`func (o *CdxV1ProviderShare) SetProviderUserName(v string)`

SetProviderUserName sets ProviderUserName field to given value.

### HasProviderUserName

`func (o *CdxV1ProviderShare) HasProviderUserName() bool`

HasProviderUserName returns a boolean if a field has been set.

### GetStatus

`func (o *CdxV1ProviderShare) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CdxV1ProviderShare) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CdxV1ProviderShare) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CdxV1ProviderShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *CdxV1ProviderShare) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *CdxV1ProviderShare) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *CdxV1ProviderShare) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *CdxV1ProviderShare) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetInvitedAt

`func (o *CdxV1ProviderShare) GetInvitedAt() time.Time`

GetInvitedAt returns the InvitedAt field if non-nil, zero value otherwise.

### GetInvitedAtOk

`func (o *CdxV1ProviderShare) GetInvitedAtOk() (*time.Time, bool)`

GetInvitedAtOk returns a tuple with the InvitedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedAt

`func (o *CdxV1ProviderShare) SetInvitedAt(v time.Time)`

SetInvitedAt sets InvitedAt field to given value.

### HasInvitedAt

`func (o *CdxV1ProviderShare) HasInvitedAt() bool`

HasInvitedAt returns a boolean if a field has been set.

### GetInviteExpiresAt

`func (o *CdxV1ProviderShare) GetInviteExpiresAt() time.Time`

GetInviteExpiresAt returns the InviteExpiresAt field if non-nil, zero value otherwise.

### GetInviteExpiresAtOk

`func (o *CdxV1ProviderShare) GetInviteExpiresAtOk() (*time.Time, bool)`

GetInviteExpiresAtOk returns a tuple with the InviteExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteExpiresAt

`func (o *CdxV1ProviderShare) SetInviteExpiresAt(v time.Time)`

SetInviteExpiresAt sets InviteExpiresAt field to given value.

### HasInviteExpiresAt

`func (o *CdxV1ProviderShare) HasInviteExpiresAt() bool`

HasInviteExpiresAt returns a boolean if a field has been set.

### GetRedeemedAt

`func (o *CdxV1ProviderShare) GetRedeemedAt() time.Time`

GetRedeemedAt returns the RedeemedAt field if non-nil, zero value otherwise.

### GetRedeemedAtOk

`func (o *CdxV1ProviderShare) GetRedeemedAtOk() (*time.Time, bool)`

GetRedeemedAtOk returns a tuple with the RedeemedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedAt

`func (o *CdxV1ProviderShare) SetRedeemedAt(v time.Time)`

SetRedeemedAt sets RedeemedAt field to given value.

### HasRedeemedAt

`func (o *CdxV1ProviderShare) HasRedeemedAt() bool`

HasRedeemedAt returns a boolean if a field has been set.

### GetSharedResource

`func (o *CdxV1ProviderShare) GetSharedResource() ObjectReference`

GetSharedResource returns the SharedResource field if non-nil, zero value otherwise.

### GetSharedResourceOk

`func (o *CdxV1ProviderShare) GetSharedResourceOk() (*ObjectReference, bool)`

GetSharedResourceOk returns a tuple with the SharedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedResource

`func (o *CdxV1ProviderShare) SetSharedResource(v ObjectReference)`

SetSharedResource sets SharedResource field to given value.

### HasSharedResource

`func (o *CdxV1ProviderShare) HasSharedResource() bool`

HasSharedResource returns a boolean if a field has been set.

### GetServiceAccount

`func (o *CdxV1ProviderShare) GetServiceAccount() ObjectReference`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *CdxV1ProviderShare) GetServiceAccountOk() (*ObjectReference, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *CdxV1ProviderShare) SetServiceAccount(v ObjectReference)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *CdxV1ProviderShare) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


