package davudpasha

import (
	"encoding/json"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

type ReportsQuery struct {
	Name               *string                    `json:"Nane,omitempty"`
	Data               *ReportsQueryData          `json:"Data,omitempty"`
	ShowTable          *bool                      `json:"ShowTable,omitempty"`
	TableVisualization *ReportsTableVisualization `json:"TableVisualization,omitempty"`
	ShowChart          *bool                      `json:"ShowChart,omitempty"`
	ChartVisualization *ReportsChartVisualization `json:"ChartVisualization,omitempty"`
	ExtData            *ReportsQueryExtData       `json:"ExtData,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}

// NewReportsQuery creates a new ReportsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportsQuery() *ReportsQuery {
	this := ReportsQuery{}
	return &this
}

// NewReportsQueryWithDefaults creates a new ReportsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportsQueryWithDefaults() *ReportsQuery {
	this := ReportsQuery{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReportsQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReportsQuery) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReportsQuery) SetName(v string) {
	o.Name = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReportsQuery) GetData() ReportsQueryData {
	if o == nil || o.Data == nil {
		var ret ReportsQueryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetDataOk() (*ReportsQueryData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReportsQuery) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given ReportsQueryData and assigns it to the Data field.
func (o *ReportsQuery) SetData(v ReportsQueryData) {
	o.Data = &v
}

// GetShowTable returns the ShowTable field value if set, zero value otherwise.
func (o *ReportsQuery) GetShowTable() bool {
	if o == nil || o.ShowTable == nil {
		var ret bool
		return ret
	}
	return *o.ShowTable
}

// GetShowTableOk returns a tuple with the ShowTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetShowTableOk() (*bool, bool) {
	if o == nil || o.ShowTable == nil {
		return nil, false
	}
	return o.ShowTable, true
}

// HasShowTable returns a boolean if a field has been set.
func (o *ReportsQuery) HasShowTable() bool {
	return o != nil && o.ShowTable != nil
}

// SetShowTable gets a reference to the given bool and assigns it to the ShowTable field.
func (o *ReportsQuery) SetShowTable(v bool) {
	o.ShowTable = &v
}

// GetTableVisualization returns the TableVisualization field value if set, zero value otherwise.
func (o *ReportsQuery) GetTableVisualization() ReportsTableVisualization {
	if o == nil || o.TableVisualization == nil {
		var ret ReportsTableVisualization
		return ret
	}
	return *o.TableVisualization
}

// GetTableVisualizationOk returns a tuple with the TableVisualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetTableVisualizationOk() (*ReportsTableVisualization, bool) {
	if o == nil || o.TableVisualization == nil {
		return nil, false
	}
	return o.TableVisualization, true
}

// HasTableVisualization returns a boolean if a field has been set.
func (o *ReportsQuery) HasTableVisualization() bool {
	return o != nil && o.TableVisualization != nil
}

// SetTableVisualization gets a reference to the given ReportsTableVisualization and assigns it to the TableVisualization field.
func (o *ReportsQuery) SetTableVisualization(v ReportsTableVisualization) {
	o.TableVisualization = &v
}

// GetShowChart returns the ShowChart field value if set, zero value otherwise.
func (o *ReportsQuery) GetShowChart() bool {
	if o == nil || o.ShowChart == nil {
		var ret bool
		return ret
	}
	return *o.ShowChart
}

// GetShowChartOk returns a tuple with the ShowChart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetShowChartOk() (*bool, bool) {
	if o == nil || o.ShowChart == nil {
		return nil, false
	}
	return o.ShowChart, true
}

// HasShowChart returns a boolean if a field has been set.
func (o *ReportsQuery) HasShowChart() bool {
	return o != nil && o.ShowChart != nil
}

// SetShowChart gets a reference to the given bool and assigns it to the ShowChart field.
func (o *ReportsQuery) SetShowChart(v bool) {
	o.ShowChart = &v
}

// GetChartVisualization returns the ChartVisualization field value if set, zero value otherwise.
func (o *ReportsQuery) GetChartVisualization() ReportsChartVisualization {
	if o == nil || o.ChartVisualization == nil {
		var ret ReportsChartVisualization
		return ret
	}
	return *o.ChartVisualization
}

// GetChartVisualizationOk returns a tuple with the ChartVisualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetChartVisualizationOk() (*ReportsChartVisualization, bool) {
	if o == nil || o.ChartVisualization == nil {
		return nil, false
	}
	return o.ChartVisualization, true
}

// HasChartVisualization returns a boolean if a field has been set.
func (o *ReportsQuery) HasChartVisualization() bool {
	return o != nil && o.ChartVisualization != nil
}

// SetChartVisualization gets a reference to the given ReportsChartVisualization and assigns it to the ChartVisualization field.
func (o *ReportsQuery) SetChartVisualization(v ReportsChartVisualization) {
	o.ChartVisualization = &v
}

// GetExtData returns the ExtData field value if set, zero value otherwise.
func (o *ReportsQuery) GetExtData() ReportsQueryExtData {
	if o == nil || o.ExtData == nil {
		var ret ReportsQueryExtData
		return ret
	}
	return *o.ExtData
}

// GetExtDataOk returns a tuple with the ExtData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsQuery) GetExtDataOk() (*ReportsQueryExtData, bool) {
	if o == nil || o.ExtData == nil {
		return nil, false
	}
	return o.ExtData, true
}

// HasExtData returns a boolean if a field has been set.
func (o *ReportsQuery) HasExtData() bool {
	return o != nil && o.ExtData != nil
}

// SetExtData gets a reference to the given ReportsQueryExtData and assigns it to the ExtData field.
func (o *ReportsQuery) SetExtData(v ReportsQueryExtData) {
	o.ExtData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if o.ShowTable != nil {
		toSerialize["ShowTable"] = o.ShowTable
	}
	if o.TableVisualization != nil {
		toSerialize["TableVisualization"] = o.TableVisualization
	}
	if o.ShowChart != nil {
		toSerialize["ShowChart"] = o.ShowChart
	}
	if o.ChartVisualization != nil {
		toSerialize["ChartVisualization"] = o.ChartVisualization
	}
	if o.ExtData != nil {
		toSerialize["ExtData"] = o.ExtData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnMarshalJSON deserializes the given payload.
func (o *ReportsQuery) UnMarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name               *string                    `json:"Nane,omitempty"`
		Data               *ReportsQueryData          `json:"Data,omitempty"`
		ShowTable          *bool                      `json:"ShowTable,omitempty"`
		TableVisualization *ReportsTableVisualization `json:"TableVisualization,omitempty"`
		ShowChart          *bool                      `json:"ShowChart,omitempty"`
		ChartVisualization *ReportsChartVisualization `json:"ChartVisualization,omitempty"`
		ExtData            *ReportsQueryExtData       `json:"ExtData,omitempty"`
	}{}

	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Name", "Data", "ShowTable", "TableVisualization", "ShowChart", "ChartVisualization", "ExtData"})
	} else {
		return err
	}

	o.Name = all.Name
	o.Data = all.Data
	o.ShowTable = all.ShowTable
	o.TableVisualization = all.TableVisualization
	o.ShowChart = all.ShowChart
	o.ChartVisualization = all.ChartVisualization
	o.ExtData = all.ExtData

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
