package davudpasha

import (
	"encoding/json"
	"fmt"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// AlertsCorrelationData represents correlation data for alerts.
type AlertsCorrelationData struct {
	// ID of the alert correlation.
	ID *string `json:"Id,omitempty"`
	// Whether the correlation is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// Name of the correlation.
	Name *string `json:"Name,omitempty"`
	// Description of the correlation.
	Description common.NullableString `json:"Description,omitempty"`
	// Tags associated with the correlation.
	Tags common.NullableList[string] `json:"Tags,omitempty"`
	// Whether the correlation is from the market.
	FromMarket *bool `json:"FromMarket,omitempty"`
	// Whether the correlation is from modules.
	FromModules *bool `json:"FromModules,omitempty"`
	// Whether the correlation has an update.
	HasUpdate *bool `json:"HasUpdate,omitempty"`
	// ID of the module.
	ModuleID common.NullableString `json:"ModuleId,omitempty"`
	// GUID of the module.
	ModuleGUID common.NullableString `json:"ModuleGuid,omitempty"`
	// Risk level of the correlation.
	RiskLevel *int64 `json:"RiskLevel,omitempty"`
	// Maximum alert count.
	MaxAlertCount *int64 `json:"MaxAlertCount,omitempty"`
	// Time frame type for limiter.
	LimiterTimeFrameType common.NullableString `json:"LimiterTimeFrameType,omitempty"`
	// Time frame value for limiter.
	LimiterTimeFrameValue common.NullableInt64 `json:"LimiterTimeFrameValue,omitempty"`
	// Columns used for limiting.
	LimiterColumns common.NullableList[string] `json:"LimiterColumns,omitempty"`
	// Actions related to the correlation.
	Actions *Actions `json:"Actions,omitempty"`
	// Maximum send count.
	MaxSendCount *int64 `json:"MaxSendCount,omitempty"`
	// Raw data in JSON format.
	Data *AlertsQueryData `json:"Data,omitempty"`
	// Type of correlation.
	CorrelationType AlertsCorrelationType `json:"CorrelationType,omitempty"`
	// Message of the correlation.
	Message *string `json:"Message,omitempty"`
	// Columns grouped by the correlation.
	GroupedColumns []string `json:"GroupedColumns,omitempty"`
	// Options for columns that can be grouped.
	GroupedColumnsOptions []SelectedColumn `json:"GroupedColumnsOptions,omitempty"`
	// Version of the correlation.
	Version *float64 `json:"Version,omitempty"`
	// Raw value if deserialization fails.
	UnparsedObject map[string]interface{} `json:"-"`
	// Additional properties not defined in the struct.
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertsCorrelationData creates a new AlertsCorrelationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertsCorrelationData() *AlertsCorrelationData {
	this := AlertsCorrelationData{}
	return &this
}

// NewAlertsCorrelationDataWithDefaults creates a new AlertsCorrelationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertsCorrelationDataWithDefaults() *AlertsCorrelationData {
	this := AlertsCorrelationData{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasID() bool {
	return o != nil && o.ID != nil
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *AlertsCorrelationData) SetID(v string) {
	o.ID = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlertsCorrelationData) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertsCorrelationData) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsCorrelationData) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *AlertsCorrelationData) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *AlertsCorrelationData) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnSetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *AlertsCorrelationData) UnSetDescription() {
	o.Description.UnSet()
}

// GetTags returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetTags() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a Tags has been set.
func (o *AlertsCorrelationData) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given common.Nullable[]string and assigns it to the Tags field.
func (o *AlertsCorrelationData) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *AlertsCorrelationData) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnSetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *AlertsCorrelationData) UnSetTags() {
	o.Tags.UnSet()
}

// GetFromMarket returns the FromMarket field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetFromMarket() bool {
	if o == nil || o.FromMarket == nil {
		var ret bool
		return ret
	}
	return *o.FromMarket
}

// GetFromMarketOk returns a tuple with the FromMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetFromMarketOk() (*bool, bool) {
	if o == nil || o.FromMarket == nil {
		return nil, false
	}
	return o.FromMarket, true
}

// HasFromMarket returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasFromMarket() bool {
	return o != nil && o.FromMarket != nil
}

// SetFromMarket gets a reference to the given bool and assigns it to the FromMarket field.
func (o *AlertsCorrelationData) SetFromMarket(v bool) {
	o.FromMarket = &v
}

// GetFromModules returns the FromModules field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetFromModules() bool {
	if o == nil || o.FromModules == nil {
		var ret bool
		return ret
	}
	return *o.FromModules
}

// GetFromModulesOk returns a tuple with the FromModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetFromModulesOk() (*bool, bool) {
	if o == nil || o.FromModules == nil {
		return nil, false
	}
	return o.FromModules, true
}

// HasFromModules returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasFromModules() bool {
	return o != nil && o.FromModules != nil
}

// SetFromModules gets a reference to the given bool and assigns it to the FromModules field.
func (o *AlertsCorrelationData) SetFromModules(v bool) {
	o.FromModules = &v
}

// GetHasUpdate returns the HasUpdate field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetHasUpdate() bool {
	if o == nil || o.HasUpdate == nil {
		var ret bool
		return ret
	}
	return *o.HasUpdate
}

// GetHasUpdateOk returns a tuple with the HasUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetHasUpdateOk() (*bool, bool) {
	if o == nil || o.HasUpdate == nil {
		return nil, false
	}
	return o.HasUpdate, true
}

// HasHasUpdate returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasHasUpdate() bool {
	return o != nil && o.HasUpdate != nil
}

// SetHasUpdate gets a reference to the given bool and assigns it to the HasUpdate field.
func (o *AlertsCorrelationData) SetHasUpdate(v bool) {
	o.HasUpdate = &v
}

// GetModuleID returns the ModuleID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsCorrelationData) GetModuleID() string {
	if o == nil || o.ModuleID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleID.Get()
}

// GetModuleIDOk returns a tuple with the ModuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetModuleIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleID.Get(), o.ModuleID.IsSet()
}

// HasModuleID returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasModuleID() bool {
	return o != nil && o.ModuleID.IsSet()
}

// SetModuleID gets a reference to the given common.NullableString and assigns it to the ModuleID field.
func (o *AlertsCorrelationData) SetModuleID(v string) {
	o.ModuleID.Set(&v)
}

// SetModuleIDNil sets the value for ModuleID to be an explicit nil.
func (o *AlertsCorrelationData) SetModuleIDNil() {
	o.ModuleID.Set(nil)
}

// UnSetModuleID ensures that no value is present for ModuleID, not even an explicit nil.
func (o *AlertsCorrelationData) UnSetModuleID() {
	o.ModuleID.UnSet()
}

// GetModuleGUID returns the ModuleGUID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsCorrelationData) GetModuleGUID() string {
	if o == nil || o.ModuleGUID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModuleGUID.Get()
}

// GetModuleGUIDOk returns a tuple with the ModuleGUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetModuleGUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleGUID.Get(), o.ModuleGUID.IsSet()
}

// HasModuleGUID returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasModuleGUID() bool {
	return o != nil && o.ModuleGUID.IsSet()
}

// SetModuleGUID gets a reference to the given common.NullableString and assigns it to the ModuleGUID field.
func (o *AlertsCorrelationData) SetModuleGUID(v string) {
	o.ModuleGUID.Set(&v)
}

// SetModuleGUIDNil sets the value for ModuleGUID to be an explicit nil.
func (o *AlertsCorrelationData) SetModuleGUIDNil() {
	o.ModuleGUID.Set(nil)
}

// UnSetModuleGUID ensures that no value is present for ModuleGUID, not even an explicit nil.
func (o *AlertsCorrelationData) UnSetModuleGUID() {
	o.ModuleGUID.UnSet()
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetRiskLevel() int64 {
	if o == nil || o.RiskLevel == nil {
		var ret int64
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetRiskLevelOk() (*int64, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasRiskLevel() bool {
	return o != nil && o.RiskLevel != nil
}

// SetRiskLevel gets a reference to the given int64 and assigns it to the RiskLevel field.
func (o *AlertsCorrelationData) SetRiskLevel(v int64) {
	o.RiskLevel = &v
}

// GetMaxAlertCount returns the MaxAlertCount field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetMaxAlertCount() int64 {
	if o == nil || o.MaxAlertCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxAlertCount
}

// GetMaxAlertCountOk returns a tuple with the MaxAlertCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetMaxAlertCountOk() (*int64, bool) {
	if o == nil || o.MaxAlertCount == nil {
		return nil, false
	}
	return o.MaxAlertCount, true
}

// HasMaxAlertCount returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasMaxAlertCount() bool {
	return o != nil && o.MaxAlertCount != nil
}

// SetMaxAlertCount gets a reference to the given int64 and assigns it to the MaxAlertCount field.
func (o *AlertsCorrelationData) SetMaxAlertCount(v int64) {
	o.MaxAlertCount = &v
}

// GetLimiterTimeFrameType returns the LimiterTimeFrameType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsCorrelationData) GetLimiterTimeFrameType() string {
	if o == nil || o.LimiterTimeFrameType.Get() == nil {
		var ret string
		return ret
	}
	return *o.LimiterTimeFrameType.Get()
}

// GetLimiterTimeFrameTypeOk returns a tuple with the LimiterTimeFrameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetLimiterTimeFrameTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimiterTimeFrameType.Get(), o.LimiterTimeFrameType.IsSet()
}

// HasLimiterTimeFrameType returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasLimiterTimeFrameType() bool {
	return o != nil && o.LimiterTimeFrameType.IsSet()
}

// SetLimiterTimeFrameType gets a reference to the given common.NullableString and assigns it to the LimiterTimeFrameType field.
func (o *AlertsCorrelationData) SetLimiterTimeFrameType(v string) {
	o.LimiterTimeFrameType.Set(&v)
}

// SetLimiterTimeFrameTypeNil sets the value for LimiterTimeFrameType to be an explicit nil.
func (o *AlertsCorrelationData) SetLimiterTimeFrameTypeNil() {
	o.LimiterTimeFrameType.Set(nil)
}

// UnSetLimiterTimeFrameType ensures that no value is present for LimiterTimeFrameType, not even an explicit nil.
func (o *AlertsCorrelationData) UnSetLimiterTimeFrameType() {
	o.LimiterTimeFrameType.UnSet()
}

// GetLimiterColumns returns a tuple with the LimiterColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetLimiterColumns() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimiterColumns.Get(), o.LimiterColumns.IsSet()
}

// GetLimiterColumnsOk returns a tuple with the LimiterColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AlertsCorrelationData) GetLimiterColumnsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimiterColumns.Get(), o.LimiterColumns.IsSet()
}

// HasLimiterColumns returns a boolean if a LimiterColumns has been set.
func (o *AlertsCorrelationData) HasLimiterColumns() bool {
	return o != nil && o.LimiterColumns.IsSet()
}

// SetLimiterColumns gets a reference to the given common.Nullable[]string and assigns it to the LimiterColumns field.
func (o *AlertsCorrelationData) SetLimiterColumns(v []string) {
	o.LimiterColumns.Set(&v)
}

// SetLimiterColumnsNil sets the value for LimiterColumns to be an explicit nil.
func (o *AlertsCorrelationData) SetLimiterColumnsNil() {
	o.LimiterColumns.Set(nil)
}

// UnSetLimiterColumns ensures that no value is present for LimiterColumns, not even an explicit nil.
func (o *AlertsCorrelationData) UnSetLimiterColumns() {
	o.LimiterColumns.UnSet()
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetActions() Actions {
	if o == nil || o.Actions == nil {
		var ret Actions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetActionsOk() (*Actions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given Actions and assigns it to the Actions field.
func (o *AlertsCorrelationData) SetActions(v Actions) {
	o.Actions = &v
}

// GetMaxSendCount returns the MaxSendCount field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetMaxSendCount() int64 {
	if o == nil || o.MaxSendCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxSendCount
}

// GetMaxSendCountOk returns a tuple with the MaxSendCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetMaxSendCountOk() (*int64, bool) {
	if o == nil || o.MaxSendCount == nil {
		return nil, false
	}
	return o.MaxSendCount, true
}

// HasMaxSendCount returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasMaxSendCount() bool {
	return o != nil && o.MaxSendCount != nil
}

// SetMaxSendCount gets a reference to the given int64 and assigns it to the MaxSendCount field.
func (o *AlertsCorrelationData) SetMaxSendCount(v int64) {
	o.MaxSendCount = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetData() AlertsQueryData {
	if o == nil || o.Data == nil {
		var ret AlertsQueryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetDataOk() (*AlertsQueryData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given AlertsQueryData and assigns it to the Data field.
func (o *AlertsCorrelationData) SetData(v AlertsQueryData) {
	o.Data = &v
}

// GetCorrelationType returns the CorrelationType field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetCorrelationType() AlertsCorrelationType {
	if o == nil {
		var ret AlertsCorrelationType
		return ret
	}
	return o.CorrelationType
}

// GetCorrelationTypeOk returns a tuple with the CorrelationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetCorrelationTypeOk() (*AlertsCorrelationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorrelationType, true
}

// SetCorrelationType gets a reference to the given string and assigns it to the CorrelationType field.
func (o *AlertsCorrelationData) SetCorrelationType(v AlertsCorrelationType) {
	o.CorrelationType = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AlertsCorrelationData) SetMessage(v string) {
	o.Message = &v
}

// GetGroupedColumns returns the GroupedColumns field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetGroupedColumns() []string {
	if o == nil || o.GroupedColumns == nil {
		var ret []string
		return ret
	}
	return o.GroupedColumns
}

// GetGroupedColumnsOk returns a tuple with the GroupedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetGroupedColumnsOk() (*[]string, bool) {
	if o == nil || o.GroupedColumns == nil {
		return nil, false
	}
	return &o.GroupedColumns, true
}

// HasGroupedColumns returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasGroupedColumns() bool {
	return o != nil && o.GroupedColumns != nil
}

// SetGroupedColumns gets a reference to the given []string and assigns it to the GroupedColumns field.
func (o *AlertsCorrelationData) SetGroupedColumns(v []string) {
	o.GroupedColumns = v
}

// GetGroupedColumnsOptions returns the GroupedColumnsOptions field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetGroupedColumnsOptions() []SelectedColumn {
	if o == nil || o.GroupedColumnsOptions == nil {
		var ret []SelectedColumn
		return ret
	}
	return o.GroupedColumnsOptions
}

// GetGroupedColumnsOptionsOk returns a tuple with the GroupedColumnsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetGroupedColumnsOptionsOk() (*[]SelectedColumn, bool) {
	if o == nil || o.GroupedColumnsOptions == nil {
		return nil, false
	}
	return &o.GroupedColumnsOptions, true
}

// HasGroupedColumnsOptions returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasGroupedColumnsOptions() bool {
	return o != nil && o.GroupedColumnsOptions != nil
}

// SetGroupedColumnsOptions gets a reference to the given []SelectedColumn and assigns it to the GroupedColumnsOptions field.
func (o *AlertsCorrelationData) SetGroupedColumnsOptions(v []SelectedColumn) {
	o.GroupedColumnsOptions = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AlertsCorrelationData) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsCorrelationData) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AlertsCorrelationData) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *AlertsCorrelationData) SetVersion(v float64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertsCorrelationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ID != nil {
		toSerialize["Id"] = o.ID
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["Tags"] = o.Tags.Get()
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
	if o.ModuleID.IsSet() {
		toSerialize["ModuleId"] = o.ModuleID.Get()
	}
	if o.ModuleGUID.IsSet() {
		toSerialize["ModuleGuid"] = o.ModuleGUID.Get()
	}
	if o.RiskLevel != nil {
		toSerialize["RiskLevel"] = o.RiskLevel
	}
	if o.MaxAlertCount != nil {
		toSerialize["MaxAlertCount"] = o.MaxAlertCount
	}
	if o.LimiterTimeFrameType.IsSet() {
		toSerialize["LimiterTimeFrameType"] = o.LimiterTimeFrameType.Get()
	}
	if o.LimiterTimeFrameValue.IsSet() {
		toSerialize["LimiterTimeFrameValue"] = o.LimiterTimeFrameValue.Get()
	}
	if o.LimiterColumns.IsSet() {
		toSerialize["LimiterColumns"] = o.LimiterColumns.Get()
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if o.MaxSendCount != nil {
		toSerialize["MaxSendCount"] = o.MaxSendCount
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if o.CorrelationType.IsValid() {
		toSerialize["CorrelationType"] = o.CorrelationType
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.GroupedColumns != nil {
		toSerialize["GroupedColumns"] = o.GroupedColumns
	}
	if o.GroupedColumnsOptions != nil {
		toSerialize["GroupedColumnsOptions"] = o.GroupedColumnsOptions
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
func (o *AlertsCorrelationData) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		ID                    *string                     `json:"Id,omitempty"`
		Enabled               *bool                       `json:"Enabled,omitempty"`
		Name                  *string                     `json:"Name,omitempty"`
		Description           common.NullableString       `json:"Description,omitempty"`
		Tags                  common.NullableList[string] `json:"Tags,omitempty"`
		FromMarket            *bool                       `json:"FromMarket,omitempty"`
		FromModules           *bool                       `json:"FromModules,omitempty"`
		HasUpdate             *bool                       `json:"HasUpdate,omitempty"`
		ModuleID              common.NullableString       `json:"ModuleId,omitempty"`
		ModuleGUID            common.NullableString       `json:"ModuleGuid,omitempty"`
		RiskLevel             *int64                      `json:"RiskLevel,omitempty"`
		MaxAlertCount         *int64                      `json:"MaxAlertCount,omitempty"`
		LimiterTimeFrameType  common.NullableString       `json:"LimiterTimeFrameType,omitempty"`
		LimiterTimeFrameValue common.NullableInt64        `json:"LimiterTimeFrameValue,omitempty"`
		LimiterColumns        common.NullableList[string] `json:"LimiterColumns,omitempty"`
		Actions               *Actions                    `json:"Actions,omitempty"`
		MaxSendCount          *int64                      `json:"MaxSendCount,omitempty"`
		Data                  *AlertsQueryData            `json:"Data,omitempty"`
		CorrelationType       *AlertsCorrelationType      `json:"CorrelationType,omitempty"`
		Message               *string                     `json:"Message,omitempty"`
		GroupedColumns        []string                    `json:"GroupedColumns,omitempty"`
		GroupedColumnsOptions []SelectedColumn            `json:"GroupedColumnsOptions,omitempty"`
		Version               *float64                    `json:"Version,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("requiered field Name is missing")
	}
	if all.Message == nil {
		return fmt.Errorf("requiered field Message is missing")
	}
	if all.CorrelationType == nil {
		return fmt.Errorf("requiered field CorrelationType is missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Id", "Enabled", "Name", "Query", "Description", "Tags", "FromMarket", "FromModules", "HasUpdate", "ModuleId", "ModuleGuid", "RiskLevel", "MaxAlertCount", "LimiterTimeFrameType", "LimiterTimeFrameValue", "LimiterColumns", "Actions", "MaxSendCount", "Data", "CorrelationType", "Message", "GroupedColumns", "GroupedColumnsOptions", "Version"})
	} else {
		return err
	}

	o.ID = all.ID
	o.Enabled = all.Enabled
	o.Name = all.Name
	o.Description = all.Description
	o.Tags = all.Tags
	o.FromMarket = all.FromMarket
	o.FromModules = all.FromModules
	o.HasUpdate = all.HasUpdate
	o.ModuleID = all.ModuleID
	o.ModuleGUID = all.ModuleGUID
	o.RiskLevel = all.RiskLevel
	o.MaxAlertCount = all.MaxAlertCount
	o.LimiterTimeFrameType = all.LimiterTimeFrameType
	o.LimiterTimeFrameValue = all.LimiterTimeFrameValue
	o.LimiterColumns = all.LimiterColumns
	hasInvalidField := false
	if all.Actions != nil && all.Actions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Actions = all.Actions
	o.MaxSendCount = all.MaxSendCount
	o.Data = all.Data
	if !all.CorrelationType.IsValid() {
		hasInvalidField = true
	} else {
		o.CorrelationType = *all.CorrelationType
	}
	o.Message = all.Message
	o.GroupedColumns = all.GroupedColumns
	o.GroupedColumnsOptions = all.GroupedColumnsOptions
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
