/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// PaymentMethod struct for PaymentMethod
type PaymentMethod struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version *string `json:"version,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Status *PaymentMethodStatus `json:"status,omitempty"`
	Messages []string `json:"messages,omitempty"`
	StripePaymentMethodId *string `json:"stripe_payment_method_id,omitempty"`
	AuthorizationVerifiedAt *time.Time `json:"authorization_verified_at,omitempty"`
	AuthorizationCanceledAt *time.Time `json:"authorization_canceled_at,omitempty"`
	AuthorizationStripePaymentIntentId *string `json:"authorization_stripe_payment_intent_id,omitempty"`
	AuthorizationStripePaymentIntentClientSecret *string `json:"authorization_stripe_payment_intent_client_secret,omitempty"`
	CardBrand *string `json:"card_brand,omitempty"`
	CardCountry *string `json:"card_country,omitempty"`
	CardFunding *string `json:"card_funding,omitempty"`
	CardFingerprint *string `json:"card_fingerprint,omitempty"`
	CardLastDigits *string `json:"card_last_digits,omitempty"`
	CardExpirationMonth *int64 `json:"card_expiration_month,omitempty"`
	CardExpirationYear *int64 `json:"card_expiration_year,omitempty"`
}

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod() *PaymentMethod {
	this := PaymentMethod{}
	var status PaymentMethodStatus = PAYMENTMETHODSTATUS_INVALID
	this.Status = &status
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	var status PaymentMethodStatus = PAYMENTMETHODSTATUS_INVALID
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentMethod) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentMethod) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentMethod) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PaymentMethod) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PaymentMethod) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PaymentMethod) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PaymentMethod) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PaymentMethod) SetVersion(v string) {
	o.Version = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *PaymentMethod) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *PaymentMethod) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *PaymentMethod) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethod) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethod) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethod) SetType(v string) {
	o.Type = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PaymentMethod) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
    return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PaymentMethod) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *PaymentMethod) SetProvider(v string) {
	o.Provider = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentMethod) GetStatus() PaymentMethodStatus {
	if o == nil || isNil(o.Status) {
		var ret PaymentMethodStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentMethod) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PaymentMethodStatus and assigns it to the Status field.
func (o *PaymentMethod) SetStatus(v PaymentMethodStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *PaymentMethod) GetMessages() []string {
	if o == nil || isNil(o.Messages) {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetMessagesOk() ([]string, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *PaymentMethod) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *PaymentMethod) SetMessages(v []string) {
	o.Messages = v
}

// GetStripePaymentMethodId returns the StripePaymentMethodId field value if set, zero value otherwise.
func (o *PaymentMethod) GetStripePaymentMethodId() string {
	if o == nil || isNil(o.StripePaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StripePaymentMethodId
}

// GetStripePaymentMethodIdOk returns a tuple with the StripePaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetStripePaymentMethodIdOk() (*string, bool) {
	if o == nil || isNil(o.StripePaymentMethodId) {
    return nil, false
	}
	return o.StripePaymentMethodId, true
}

// HasStripePaymentMethodId returns a boolean if a field has been set.
func (o *PaymentMethod) HasStripePaymentMethodId() bool {
	if o != nil && !isNil(o.StripePaymentMethodId) {
		return true
	}

	return false
}

// SetStripePaymentMethodId gets a reference to the given string and assigns it to the StripePaymentMethodId field.
func (o *PaymentMethod) SetStripePaymentMethodId(v string) {
	o.StripePaymentMethodId = &v
}

// GetAuthorizationVerifiedAt returns the AuthorizationVerifiedAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetAuthorizationVerifiedAt() time.Time {
	if o == nil || isNil(o.AuthorizationVerifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.AuthorizationVerifiedAt
}

// GetAuthorizationVerifiedAtOk returns a tuple with the AuthorizationVerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetAuthorizationVerifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.AuthorizationVerifiedAt) {
    return nil, false
	}
	return o.AuthorizationVerifiedAt, true
}

// HasAuthorizationVerifiedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasAuthorizationVerifiedAt() bool {
	if o != nil && !isNil(o.AuthorizationVerifiedAt) {
		return true
	}

	return false
}

// SetAuthorizationVerifiedAt gets a reference to the given time.Time and assigns it to the AuthorizationVerifiedAt field.
func (o *PaymentMethod) SetAuthorizationVerifiedAt(v time.Time) {
	o.AuthorizationVerifiedAt = &v
}

// GetAuthorizationCanceledAt returns the AuthorizationCanceledAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetAuthorizationCanceledAt() time.Time {
	if o == nil || isNil(o.AuthorizationCanceledAt) {
		var ret time.Time
		return ret
	}
	return *o.AuthorizationCanceledAt
}

// GetAuthorizationCanceledAtOk returns a tuple with the AuthorizationCanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetAuthorizationCanceledAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.AuthorizationCanceledAt) {
    return nil, false
	}
	return o.AuthorizationCanceledAt, true
}

// HasAuthorizationCanceledAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasAuthorizationCanceledAt() bool {
	if o != nil && !isNil(o.AuthorizationCanceledAt) {
		return true
	}

	return false
}

// SetAuthorizationCanceledAt gets a reference to the given time.Time and assigns it to the AuthorizationCanceledAt field.
func (o *PaymentMethod) SetAuthorizationCanceledAt(v time.Time) {
	o.AuthorizationCanceledAt = &v
}

// GetAuthorizationStripePaymentIntentId returns the AuthorizationStripePaymentIntentId field value if set, zero value otherwise.
func (o *PaymentMethod) GetAuthorizationStripePaymentIntentId() string {
	if o == nil || isNil(o.AuthorizationStripePaymentIntentId) {
		var ret string
		return ret
	}
	return *o.AuthorizationStripePaymentIntentId
}

// GetAuthorizationStripePaymentIntentIdOk returns a tuple with the AuthorizationStripePaymentIntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetAuthorizationStripePaymentIntentIdOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationStripePaymentIntentId) {
    return nil, false
	}
	return o.AuthorizationStripePaymentIntentId, true
}

// HasAuthorizationStripePaymentIntentId returns a boolean if a field has been set.
func (o *PaymentMethod) HasAuthorizationStripePaymentIntentId() bool {
	if o != nil && !isNil(o.AuthorizationStripePaymentIntentId) {
		return true
	}

	return false
}

// SetAuthorizationStripePaymentIntentId gets a reference to the given string and assigns it to the AuthorizationStripePaymentIntentId field.
func (o *PaymentMethod) SetAuthorizationStripePaymentIntentId(v string) {
	o.AuthorizationStripePaymentIntentId = &v
}

// GetAuthorizationStripePaymentIntentClientSecret returns the AuthorizationStripePaymentIntentClientSecret field value if set, zero value otherwise.
func (o *PaymentMethod) GetAuthorizationStripePaymentIntentClientSecret() string {
	if o == nil || isNil(o.AuthorizationStripePaymentIntentClientSecret) {
		var ret string
		return ret
	}
	return *o.AuthorizationStripePaymentIntentClientSecret
}

// GetAuthorizationStripePaymentIntentClientSecretOk returns a tuple with the AuthorizationStripePaymentIntentClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetAuthorizationStripePaymentIntentClientSecretOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationStripePaymentIntentClientSecret) {
    return nil, false
	}
	return o.AuthorizationStripePaymentIntentClientSecret, true
}

// HasAuthorizationStripePaymentIntentClientSecret returns a boolean if a field has been set.
func (o *PaymentMethod) HasAuthorizationStripePaymentIntentClientSecret() bool {
	if o != nil && !isNil(o.AuthorizationStripePaymentIntentClientSecret) {
		return true
	}

	return false
}

// SetAuthorizationStripePaymentIntentClientSecret gets a reference to the given string and assigns it to the AuthorizationStripePaymentIntentClientSecret field.
func (o *PaymentMethod) SetAuthorizationStripePaymentIntentClientSecret(v string) {
	o.AuthorizationStripePaymentIntentClientSecret = &v
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardBrand() string {
	if o == nil || isNil(o.CardBrand) {
		var ret string
		return ret
	}
	return *o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardBrandOk() (*string, bool) {
	if o == nil || isNil(o.CardBrand) {
    return nil, false
	}
	return o.CardBrand, true
}

// HasCardBrand returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardBrand() bool {
	if o != nil && !isNil(o.CardBrand) {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given string and assigns it to the CardBrand field.
func (o *PaymentMethod) SetCardBrand(v string) {
	o.CardBrand = &v
}

// GetCardCountry returns the CardCountry field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardCountry() string {
	if o == nil || isNil(o.CardCountry) {
		var ret string
		return ret
	}
	return *o.CardCountry
}

// GetCardCountryOk returns a tuple with the CardCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardCountryOk() (*string, bool) {
	if o == nil || isNil(o.CardCountry) {
    return nil, false
	}
	return o.CardCountry, true
}

// HasCardCountry returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardCountry() bool {
	if o != nil && !isNil(o.CardCountry) {
		return true
	}

	return false
}

// SetCardCountry gets a reference to the given string and assigns it to the CardCountry field.
func (o *PaymentMethod) SetCardCountry(v string) {
	o.CardCountry = &v
}

// GetCardFunding returns the CardFunding field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardFunding() string {
	if o == nil || isNil(o.CardFunding) {
		var ret string
		return ret
	}
	return *o.CardFunding
}

// GetCardFundingOk returns a tuple with the CardFunding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardFundingOk() (*string, bool) {
	if o == nil || isNil(o.CardFunding) {
    return nil, false
	}
	return o.CardFunding, true
}

// HasCardFunding returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardFunding() bool {
	if o != nil && !isNil(o.CardFunding) {
		return true
	}

	return false
}

// SetCardFunding gets a reference to the given string and assigns it to the CardFunding field.
func (o *PaymentMethod) SetCardFunding(v string) {
	o.CardFunding = &v
}

// GetCardFingerprint returns the CardFingerprint field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardFingerprint() string {
	if o == nil || isNil(o.CardFingerprint) {
		var ret string
		return ret
	}
	return *o.CardFingerprint
}

// GetCardFingerprintOk returns a tuple with the CardFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardFingerprintOk() (*string, bool) {
	if o == nil || isNil(o.CardFingerprint) {
    return nil, false
	}
	return o.CardFingerprint, true
}

// HasCardFingerprint returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardFingerprint() bool {
	if o != nil && !isNil(o.CardFingerprint) {
		return true
	}

	return false
}

// SetCardFingerprint gets a reference to the given string and assigns it to the CardFingerprint field.
func (o *PaymentMethod) SetCardFingerprint(v string) {
	o.CardFingerprint = &v
}

// GetCardLastDigits returns the CardLastDigits field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardLastDigits() string {
	if o == nil || isNil(o.CardLastDigits) {
		var ret string
		return ret
	}
	return *o.CardLastDigits
}

// GetCardLastDigitsOk returns a tuple with the CardLastDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardLastDigitsOk() (*string, bool) {
	if o == nil || isNil(o.CardLastDigits) {
    return nil, false
	}
	return o.CardLastDigits, true
}

// HasCardLastDigits returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardLastDigits() bool {
	if o != nil && !isNil(o.CardLastDigits) {
		return true
	}

	return false
}

// SetCardLastDigits gets a reference to the given string and assigns it to the CardLastDigits field.
func (o *PaymentMethod) SetCardLastDigits(v string) {
	o.CardLastDigits = &v
}

// GetCardExpirationMonth returns the CardExpirationMonth field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardExpirationMonth() int64 {
	if o == nil || isNil(o.CardExpirationMonth) {
		var ret int64
		return ret
	}
	return *o.CardExpirationMonth
}

// GetCardExpirationMonthOk returns a tuple with the CardExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardExpirationMonthOk() (*int64, bool) {
	if o == nil || isNil(o.CardExpirationMonth) {
    return nil, false
	}
	return o.CardExpirationMonth, true
}

// HasCardExpirationMonth returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardExpirationMonth() bool {
	if o != nil && !isNil(o.CardExpirationMonth) {
		return true
	}

	return false
}

// SetCardExpirationMonth gets a reference to the given int64 and assigns it to the CardExpirationMonth field.
func (o *PaymentMethod) SetCardExpirationMonth(v int64) {
	o.CardExpirationMonth = &v
}

// GetCardExpirationYear returns the CardExpirationYear field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardExpirationYear() int64 {
	if o == nil || isNil(o.CardExpirationYear) {
		var ret int64
		return ret
	}
	return *o.CardExpirationYear
}

// GetCardExpirationYearOk returns a tuple with the CardExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardExpirationYearOk() (*int64, bool) {
	if o == nil || isNil(o.CardExpirationYear) {
    return nil, false
	}
	return o.CardExpirationYear, true
}

// HasCardExpirationYear returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardExpirationYear() bool {
	if o != nil && !isNil(o.CardExpirationYear) {
		return true
	}

	return false
}

// SetCardExpirationYear gets a reference to the given int64 and assigns it to the CardExpirationYear field.
func (o *PaymentMethod) SetCardExpirationYear(v int64) {
	o.CardExpirationYear = &v
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.StripePaymentMethodId) {
		toSerialize["stripe_payment_method_id"] = o.StripePaymentMethodId
	}
	if !isNil(o.AuthorizationVerifiedAt) {
		toSerialize["authorization_verified_at"] = o.AuthorizationVerifiedAt
	}
	if !isNil(o.AuthorizationCanceledAt) {
		toSerialize["authorization_canceled_at"] = o.AuthorizationCanceledAt
	}
	if !isNil(o.AuthorizationStripePaymentIntentId) {
		toSerialize["authorization_stripe_payment_intent_id"] = o.AuthorizationStripePaymentIntentId
	}
	if !isNil(o.AuthorizationStripePaymentIntentClientSecret) {
		toSerialize["authorization_stripe_payment_intent_client_secret"] = o.AuthorizationStripePaymentIntentClientSecret
	}
	if !isNil(o.CardBrand) {
		toSerialize["card_brand"] = o.CardBrand
	}
	if !isNil(o.CardCountry) {
		toSerialize["card_country"] = o.CardCountry
	}
	if !isNil(o.CardFunding) {
		toSerialize["card_funding"] = o.CardFunding
	}
	if !isNil(o.CardFingerprint) {
		toSerialize["card_fingerprint"] = o.CardFingerprint
	}
	if !isNil(o.CardLastDigits) {
		toSerialize["card_last_digits"] = o.CardLastDigits
	}
	if !isNil(o.CardExpirationMonth) {
		toSerialize["card_expiration_month"] = o.CardExpirationMonth
	}
	if !isNil(o.CardExpirationYear) {
		toSerialize["card_expiration_year"] = o.CardExpirationYear
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

