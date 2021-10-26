# TriggerDeploymentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TriggerDeploymentMetadataTriggerType**](TriggerDeploymentMetadataTriggerType.md) |  | [optional] [default to TRIGGERDEPLOYMENTMETADATATRIGGERTYPE_UNKNOWN]
**Git** | Pointer to [**GitDeploymentMetadata**](GitDeploymentMetadata.md) |  | [optional] 

## Methods

### NewTriggerDeploymentMetadata

`func NewTriggerDeploymentMetadata() *TriggerDeploymentMetadata`

NewTriggerDeploymentMetadata instantiates a new TriggerDeploymentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerDeploymentMetadataWithDefaults

`func NewTriggerDeploymentMetadataWithDefaults() *TriggerDeploymentMetadata`

NewTriggerDeploymentMetadataWithDefaults instantiates a new TriggerDeploymentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TriggerDeploymentMetadata) GetType() TriggerDeploymentMetadataTriggerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerDeploymentMetadata) GetTypeOk() (*TriggerDeploymentMetadataTriggerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerDeploymentMetadata) SetType(v TriggerDeploymentMetadataTriggerType)`

SetType sets Type field to given value.

### HasType

`func (o *TriggerDeploymentMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGit

`func (o *TriggerDeploymentMetadata) GetGit() GitDeploymentMetadata`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *TriggerDeploymentMetadata) GetGitOk() (*GitDeploymentMetadata, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *TriggerDeploymentMetadata) SetGit(v GitDeploymentMetadata)`

SetGit sets Git field to given value.

### HasGit

`func (o *TriggerDeploymentMetadata) HasGit() bool`

HasGit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

