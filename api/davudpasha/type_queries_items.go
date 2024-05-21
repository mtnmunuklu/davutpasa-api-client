package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type QueriesItem struct {
	ID                   *string                             `json:"ID,omitempty"`
	Name                 *string                             `json:"Name,omitempty"`
	Description          common.NullableString               `json:"Description,omitempty"`
	Query                *string                             `json:"Query,omitempty"`
	Columns              common.NullableList[SelectedColumn] `json:"Columns,omitempty"`
	Author               *string                             `json:"Author,omitempty"`
	InsertDate           *string                             `json:"InsertDate,omitempty"`
	LastUpdateDate       *string                             `json:"LastUpdateDate,omitempty"`
	QueryType            *string                             `json:"QueryType,omitempty"`
	DateTimeRange        *DateTimeRange                      `json:"DateTimeRange,omitempty"`
	Tags                 common.NullableList[string]         `json:"Tags,omitempty"`
	MitreTags            common.NullableList[string]         `json:"MitreTags,omitempty"`
	KillChainPhase       common.NullableString               `json:"KillChainPhase,omitempty"`
	FromMarket           *bool                               `json:"FromMarket,omitempty"`
	FromModules          *bool                               `json:"FromModules,omitempty"`
	HasUpdate            *bool                               `json:"HasUpdate,omitempty"`
	ModuleId             common.NullableString               `json:"ModuleId,omitempty"`
	ModuleGuid           common.NullableString               `json:"ModuleGuid,omitempty"`
	SharedUsersAndGroups common.NullableList[string]         `json:"SharedUsersAndGroups,omitempty"`
	Version              *float64                            `json:"Version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewQueriesItem creates a new QueriesItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueriesItem() *QueriesItem {
	this := QueriesItem{}
	return &this
}

// NewQueriesItemWithDefaults creates a new QueriesItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueriesItemWithDefaults() *QueriesItem {
	this := QueriesItem{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *QueriesItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *QueriesItem) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *QueriesItem) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueriesItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QueriesItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueriesItem) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a Description has been set.
func (o *QueriesItem) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *QueriesItem) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *QueriesItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *QueriesItem) UnsetDescription() {
	o.Description.UnSet()
}

// GetColumns returns the Columns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItem) GetColumns() []SelectedColumn {
	if o == nil || o.Columns.Get() == nil {
		var ret []SelectedColumn
		return ret
	}
	return *o.Columns.Get()
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItem) GetColumnsOk() (*[]SelectedColumn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns.Get(), o.Columns.IsSet()
}

// HasColumns returns a boolean if a Columns has been set.
func (o *QueriesItem) HasColumns() bool {
	return o != nil && o.Columns.IsSet()
}

// SetColumns gets a reference to the given datadog.Nullable[SelectedColumns] and assigns it to the Columns field.
func (o *QueriesItem) SetColumns(v []SelectedColumn) {
	o.Columns.Set(&v)
}

// SetColumnsNil sets the value for Columns to be an explicit nil.
func (o *QueriesItem) SetColumnsNil() {
	o.Columns.Set(nil)
}

// UnsetColumns ensures that no value is present for Columns, not even an explicit nil.
func (o *QueriesItem) UnSetColumns() {
	o.Columns.UnSet()
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *QueriesItem) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *QueriesItem) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *QueriesItem) SetAuthor(v string) {
	o.Author = &v
}

// GetInsertDate returns the InsertDate field value if set, zero value otherwise.
func (o *QueriesItem) GetInsertDate() string {
	if o == nil || o.InsertDate == nil {
		var ret string
		return ret
	}
	return *o.InsertDate
}

// GetInsertDateOk returns a tuple with the InsertDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetInsertDateOk() (*string, bool) {
	if o == nil || o.InsertDate == nil {
		return nil, false
	}
	return o.InsertDate, true
}

// HasInsertDate returns a boolean if a field has been set.
func (o *QueriesItem) HasInsertDate() bool {
	return o != nil && o.InsertDate != nil
}

// SetInsertDate gets a reference to the given string and assigns it to the InsertDate field.
func (o *QueriesItem) SetInsertDate(v string) {
	o.InsertDate = &v
}

// GetLastUpdateDate returns the LastUpdateDate field value if set, zero value otherwise.
func (o *QueriesItem) GetLastUpdateDate() string {
	if o == nil || o.LastUpdateDate == nil {
		var ret string
		return ret
	}
	return *o.LastUpdateDate
}

// GetLastUpdateDateOk returns a tuple with the LastUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetLastUpdateDateOk() (*string, bool) {
	if o == nil || o.LastUpdateDate == nil {
		return nil, false
	}
	return o.LastUpdateDate, true
}

// HasLastUpdateDate returns a boolean if a field has been set.
func (o *QueriesItem) HasLastUpdateDate() bool {
	return o != nil && o.LastUpdateDate != nil
}

// SetLastUpdateDate gets a reference to the given string and assigns it to the LastUpdateDate field.
func (o *QueriesItem) SetLastUpdateDate(v string) {
	o.LastUpdateDate = &v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *QueriesItem) GetQueryType() string {
	if o == nil || o.QueryType == nil {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetQueryTypeOk() (*string, bool) {
	if o == nil || o.QueryType == nil {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *QueriesItem) HasQueryType() bool {
	return o != nil && o.QueryType != nil
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *QueriesItem) SetQueryType(v string) {
	o.QueryType = &v
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *QueriesItem) GetDateTimeRange() DateTimeRange {
	if o == nil || o.DateTimeRange == nil {
		var ret DateTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || o.DateTimeRange == nil {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *QueriesItem) HasDateTimeRange() bool {
	return o != nil && o.DateTimeRange != nil
}

// SetDateTimeRange gets a reference to the given DateTimeRange and assigns it to the DateTimeRange field.
func (o *QueriesItem) SetDateTimeRange(v DateTimeRange) {
	o.DateTimeRange = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItem) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItem) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a Tags has been set.
func (o *QueriesItem) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.Nullable[]string and assigns it to the Tags field.
func (o *QueriesItem) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *QueriesItem) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *QueriesItem) UnSetTags() {
	o.Tags.UnSet()
}

// GetMitreTags returns the MitreTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItem) GetMitreTags() []string {
	if o == nil || o.MitreTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.MitreTags.Get()
}

// GetMitreTagsOk returns a tuple with the MitreTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItem) GetMitreTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitreTags.Get(), o.MitreTags.IsSet()
}

// HasMitreTags returns a boolean if a MitreTags has been set.
func (o *QueriesItem) HasMitreTags() bool {
	return o != nil && o.MitreTags.IsSet()
}

// SetMitreTags gets a reference to the given datadog.Nullable[]string and assigns it to the MitreTags field.
func (o *QueriesItem) SetMitreTags(v []string) {
	o.MitreTags.Set(&v)
}

// SetMitreTagsNil sets the value for MitreTags to be an explicit nil.
func (o *QueriesItem) SetMitreTagsNil() {
	o.MitreTags.Set(nil)
}

// UnsetMitreTags ensures that no value is present for MitreTags, not even an explicit nil.
func (o *QueriesItem) UnSetMitreTags() {
	o.MitreTags.UnSet()
}

// GetKillChainPhase returns the KillChainPhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItem) GetKillChainPhase() string {
	if o == nil || o.KillChainPhase.Get() == nil {
		var ret string
		return ret
	}
	return *o.KillChainPhase.Get()
}

// GetKillChainPhaseOk returns a tuple with the KillChainPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItem) GetKillChainPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KillChainPhase.Get(), o.KillChainPhase.IsSet()
}

// HasKillChainPhase returns a boolean if a KillChainPhase has been set.
func (o *QueriesItem) HasKillChainPhase() bool {
	return o != nil && o.KillChainPhase.IsSet()
}

// SetKillChainPhase gets a reference to the given datadog.NullableString and assigns it to the KillChainPhase field.
func (o *QueriesItem) SetKillChainPhase(v string) {
	o.KillChainPhase.Set(&v)
}

// SetKillChainPhaseNil sets the value for KillChainPhase to be an explicit nil.
func (o *QueriesItem) SetKillChainPhaseNil() {
	o.KillChainPhase.Set(nil)
}

// UnsetKillChainPhase ensures that no value is present for KillChainPhase, not even an explicit nil.
func (o *QueriesItem) UnsetKillChainPhase() {
	o.KillChainPhase.UnSet()
}

// GetFromMarket returns the FromMarket field value if set, zero value otherwise.
func (o *QueriesItem) GetFromMarket() bool {
	if o == nil || o.FromMarket == nil {
		var ret bool
		return ret
	}
	return *o.FromMarket
}

// GetFromMarketOk returns a tuple with the FromMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetFromMarketOk() (*bool, bool) {
	if o == nil || o.FromMarket == nil {
		return nil, false
	}
	return o.FromMarket, true
}

// HasFromMarket returns a boolean if a field has been set.
func (o *QueriesItem) HasFromMarket() bool {
	return o != nil && o.FromMarket != nil
}

// SetFromMarket gets a reference to the given bool and assigns it to the FromMarket field.
func (o *QueriesItem) SetFromMarket(v bool) {
	o.FromMarket = &v
}

// GetFromModules returns the FromModules field value if set, zero value otherwise.
func (o *QueriesItem) GetFromModules() bool {
	if o == nil || o.FromModules == nil {
		var ret bool
		return ret
	}
	return *o.FromModules
}

// GetFromModulesOk returns a tuple with the FromModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetFromModulesOk() (*bool, bool) {
	if o == nil || o.FromModules == nil {
		return nil, false
	}
	return o.FromModules, true
}

// HasFromModules returns a boolean if a field has been set.
func (o *QueriesItem) HasFromModules() bool {
	return o != nil && o.FromModules != nil
}

// SetFromModules gets a reference to the given bool and assigns it to the FromModules field.
func (o *QueriesItem) SetFromModules(v bool) {
	o.FromModules = &v
}

// GetHasUpdate returns the HasUpdate field value if set, zero value otherwise.
func (o *QueriesItem) GetHasUpdate() bool {
	if o == nil || o.HasUpdate == nil {
		var ret bool
		return ret
	}
	return *o.HasUpdate
}

// GetHasUpdateOk returns a tuple with the HasUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetHasUpdateOk() (*bool, bool) {
	if o == nil || o.HasUpdate == nil {
		return nil, false
	}
	return o.HasUpdate, true
}

// HasHasUpdate returns a boolean if a field has been set.
func (o *QueriesItem) HasHasUpdate() bool {
	return o != nil && o.HasUpdate != nil
}

// SetHasUpdate gets a reference to the given bool and assigns it to the HasUpdate field.
func (o *QueriesItem) SetHasUpdate(v bool) {
	o.HasUpdate = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItem) GetModuleId() string {
	if o == nil || o.ModuleId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleId.Get()
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItem) GetModuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleId.Get(), o.ModuleId.IsSet()
}

// HasModuleId returns a boolean if a ModuleId has been set.
func (o *QueriesItem) HasModuleId() bool {
	return o != nil && o.ModuleId.IsSet()
}

// SetModuleId gets a reference to the given datadog.NullableString and assigns it to the ModuleId field.
func (o *QueriesItem) SetModuleId(v string) {
	o.ModuleId.Set(&v)
}

// SetModuleIdNil sets the value for ModuleId to be an explicit nil.
func (o *QueriesItem) SetModuleIdNil() {
	o.ModuleId.Set(nil)
}

// UnsetModuleId ensures that no value is present for ModuleId, not even an explicit nil.
func (o *QueriesItem) UnsetModuleId() {
	o.ModuleId.UnSet()
}

// GetModuleGuid returns the ModuleGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItem) GetModuleGuid() string {
	if o == nil || o.ModuleGuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleGuid.Get()
}

// GetModuleGuidOk returns a tuple with the ModuleGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItem) GetModuleGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleGuid.Get(), o.ModuleGuid.IsSet()
}

// HasModuleGuid returns a boolean if a ModuleGuid has been set.
func (o *QueriesItem) HasModuleGuid() bool {
	return o != nil && o.ModuleGuid.IsSet()
}

// SetModuleGuid gets a reference to the given datadog.NullableString and assigns it to the ModuleGuid field.
func (o *QueriesItem) SetModuleGuid(v string) {
	o.ModuleGuid.Set(&v)
}

// SetModuleGuidNil sets the value for ModuleGuid to be an explicit nil.
func (o *QueriesItem) SetModuleGuidNil() {
	o.ModuleGuid.Set(nil)
}

// UnsetModuleGuid ensures that no value is present for ModuleGuid, not even an explicit nil.
func (o *QueriesItem) UnsetModuleGuid() {
	o.ModuleGuid.UnSet()
}

// GetSharedUsersAndGroups returns the SharedUsersAndGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueriesItem) GetSharedUsersAndGroups() []string {
	if o == nil || o.SharedUsersAndGroups.Get() == nil {
		var ret []string
		return ret
	}
	return *o.SharedUsersAndGroups.Get()
}

// GetSharedUsersAndGroupsOk returns a tuple with the SharedUsersAndGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QueriesItem) GetSharedUsersAndGroupsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedUsersAndGroups.Get(), o.SharedUsersAndGroups.IsSet()
}

// HasSharedUsersAndGroups returns a boolean if a SharedUsersAndGroups has been set.
func (o *QueriesItem) HasSharedUsersAndGroups() bool {
	return o != nil && o.SharedUsersAndGroups.IsSet()
}

// SetSharedUsersAndGroups gets a reference to the given datadog.Nullable[]string and assigns it to the SharedUsersAndGroups field.
func (o *QueriesItem) SetSharedUsersAndGroups(v []string) {
	o.SharedUsersAndGroups.Set(&v)
}

// SetSharedUsersAndGroupsNil sets the value for SharedUsersAndGroups to be an explicit nil.
func (o *QueriesItem) SetSharedUsersAndGroupsNil() {
	o.SharedUsersAndGroups.Set(nil)
}

// UnsetSharedUsersAndGroups ensures that no value is present for SharedUsersAndGroups, not even an explicit nil.
func (o *QueriesItem) UnSetSharedUsersAndGroups() {
	o.SharedUsersAndGroups.UnSet()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *QueriesItem) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesItem) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *QueriesItem) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *QueriesItem) SetVersion(v float64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueriesItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Query != nil {
		toSerialize["Query"] = o.Query
	}
	if o.Columns.IsSet() {
		toSerialize["Columns"] = o.Columns.Get()
	}
	if o.Author != nil {
		toSerialize["Author"] = o.Author
	}
	if o.InsertDate != nil {
		toSerialize["InsertDate"] = o.InsertDate
	}
	if o.LastUpdateDate != nil {
		toSerialize["LastUpdateDate"] = o.LastUpdateDate
	}
	if o.QueryType != nil {
		toSerialize["QueryType"] = o.QueryType
	}
	if o.DateTimeRange != nil {
		toSerialize["DateTimeRange"] = o.DateTimeRange
	}
	if o.Tags.IsSet() {
		toSerialize["Tags"] = o.Tags.Get()
	}
	if o.MitreTags.IsSet() {
		toSerialize["MitreTags"] = o.MitreTags.Get()
	}
	if o.KillChainPhase.IsSet() {
		toSerialize["KillChainPhase"] = o.KillChainPhase.Get()
	}
	if o.FromMarket != nil {
		toSerialize["FromMarket"] = o.FromMarket
	}
	if o.FromModules != nil {
		toSerialize["FromModules"] = o.FromModules
	}
	if o.HasUpdate != nil {
		toSerialize["HasUpdate"] = o.HasUpdate
	}
	if o.ModuleId.IsSet() {
		toSerialize["ModuleId"] = o.ModuleId.Get()
	}
	if o.ModuleGuid.IsSet() {
		toSerialize["ModuleGuid"] = o.ModuleGuid.Get()
	}
	if o.SharedUsersAndGroups.IsSet() {
		toSerialize["SharedUsersAndGroups"] = o.SharedUsersAndGroups.Get()
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *QueriesItem) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID                   *string                             `json:"ID,omitempty"`
		Name                 *string                             `json:"Name,omitempty"`
		Description          common.NullableString               `json:"Description,omitempty"`
		Query                *string                             `json:"Query,omitempty"`
		Columns              common.NullableList[SelectedColumn] `json:"Columns,omitempty"`
		Author               *string                             `json:"Author,omitempty"`
		InsertDate           *string                             `json:"InsertDate,omitempty"`
		LastUpdateDate       *string                             `json:"LastUpdateDate,omitempty"`
		QueryType            *string                             `json:"QueryType,omitempty"`
		DateTimeRange        *DateTimeRange                      `json:"DateTimeRange,omitempty"`
		Tags                 common.NullableList[string]         `json:"Tags,omitempty"`
		MitreTags            common.NullableList[string]         `json:"MitreTags,omitempty"`
		KillChainPhase       common.NullableString               `json:"KillChainPhase,omitempty"`
		FromMarket           *bool                               `json:"FromMarket,omitempty"`
		FromModules          *bool                               `json:"FromModules,omitempty"`
		HasUpdate            *bool                               `json:"HasUpdate,omitempty"`
		ModuleId             common.NullableString               `json:"ModuleId,omitempty"`
		ModuleGuid           common.NullableString               `json:"ModuleGuid,omitempty"`
		SharedUsersAndGroups common.NullableList[string]         `json:"SharedUsersAndGroups,omitempty"`
		Version              *float64                            `json:"Version,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ID == nil {
		return fmt.Errorf("requiered field ID is missing")
	}
	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}
	if all.Query == nil {
		return fmt.Errorf("requiered field Query is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ID", "Name", "Description", "Query", "Columns", "Author", "InsertDate", "LastUpdateDate", "QueryType", "DateTimeRange", "Tags", "MitreTags", "KillChainPhase", "FromMarket", "FromModules", "ModuleId", "ModuleGuid", "SharedUsersAndGroups", "Version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DateTimeRange != nil && all.DateTimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DateTimeRange = all.DateTimeRange
	o.ID = all.ID
	o.Name = all.Name
	o.Description = all.Description
	o.Query = all.Query
	o.Columns = all.Columns
	o.Author = all.Author
	o.InsertDate = all.InsertDate
	o.LastUpdateDate = all.LastUpdateDate
	o.QueryType = all.QueryType
	o.Tags = all.Tags
	o.MitreTags = all.MitreTags
	o.KillChainPhase = all.KillChainPhase
	o.FromMarket = all.FromMarket
	o.FromModules = all.FromModules
	o.HasUpdate = all.HasUpdate
	o.ModuleId = all.ModuleId
	o.ModuleGuid = all.ModuleGuid
	o.SharedUsersAndGroups = all.SharedUsersAndGroups
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
