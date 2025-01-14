// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
	time "time"
)

type ThermostatsCoolRequest struct {
	DeviceId                  string   `json:"device_id"`
	CoolingSetPointCelsius    *float64 `json:"cooling_set_point_celsius,omitempty"`
	CoolingSetPointFahrenheit *float64 `json:"cooling_set_point_fahrenheit,omitempty"`
	Sync                      *bool    `json:"sync,omitempty"`
}

type ThermostatsGetRequest struct {
	DeviceId *string `json:"device_id,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type ThermostatsHeatRequest struct {
	DeviceId                  string   `json:"device_id"`
	HeatingSetPointCelsius    *float64 `json:"heating_set_point_celsius,omitempty"`
	HeatingSetPointFahrenheit *float64 `json:"heating_set_point_fahrenheit,omitempty"`
	Sync                      *bool    `json:"sync,omitempty"`
}

type ThermostatsHeatCoolRequest struct {
	DeviceId                  string   `json:"device_id"`
	HeatingSetPointCelsius    *float64 `json:"heating_set_point_celsius,omitempty"`
	HeatingSetPointFahrenheit *float64 `json:"heating_set_point_fahrenheit,omitempty"`
	CoolingSetPointCelsius    *float64 `json:"cooling_set_point_celsius,omitempty"`
	CoolingSetPointFahrenheit *float64 `json:"cooling_set_point_fahrenheit,omitempty"`
	Sync                      *bool    `json:"sync,omitempty"`
}

type ThermostatsListRequest struct {
	// List all devices owned by this connected account
	ConnectedAccountId  *string                                                  `json:"connected_account_id,omitempty"`
	ConnectedAccountIds []string                                                 `json:"connected_account_ids,omitempty"`
	ConnectWebviewId    *string                                                  `json:"connect_webview_id,omitempty"`
	DeviceType          *DeviceType                                              `json:"device_type,omitempty"`
	DeviceTypes         []DeviceType                                             `json:"device_types,omitempty"`
	Manufacturer        *Manufacturer                                            `json:"manufacturer,omitempty"`
	DeviceIds           []string                                                 `json:"device_ids,omitempty"`
	Limit               *float64                                                 `json:"limit,omitempty"`
	CreatedBefore       *time.Time                                               `json:"created_before,omitempty"`
	UserIdentifierKey   *string                                                  `json:"user_identifier_key,omitempty"`
	CustomMetadataHas   map[string]*ThermostatsListRequestCustomMetadataHasValue `json:"custom_metadata_has,omitempty"`
}

type ThermostatsOffRequest struct {
	DeviceId string `json:"device_id"`
	Sync     *bool  `json:"sync,omitempty"`
}

type ThermostatsSetFanModeRequest struct {
	DeviceId       string                                      `json:"device_id"`
	FanMode        *FanMode                                    `json:"fan_mode,omitempty"`
	FanModeSetting *ThermostatsSetFanModeRequestFanModeSetting `json:"fan_mode_setting,omitempty"`
	Sync           *bool                                       `json:"sync,omitempty"`
}

type ThermostatsCoolResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (t *ThermostatsCoolResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ThermostatsCoolResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThermostatsCoolResponse(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThermostatsCoolResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type ThermostatsGetResponse struct {
	Thermostat *Device `json:"thermostat,omitempty"`
	Ok         bool    `json:"ok"`

	_rawJSON json.RawMessage
}

func (t *ThermostatsGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ThermostatsGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThermostatsGetResponse(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThermostatsGetResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type ThermostatsHeatCoolResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (t *ThermostatsHeatCoolResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ThermostatsHeatCoolResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThermostatsHeatCoolResponse(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThermostatsHeatCoolResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type ThermostatsHeatResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (t *ThermostatsHeatResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ThermostatsHeatResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThermostatsHeatResponse(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThermostatsHeatResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type ThermostatsListRequestCustomMetadataHasValue struct {
	typeName       string
	String         string
	Boolean        bool
	StringOptional *string
}

func NewThermostatsListRequestCustomMetadataHasValueFromString(value string) *ThermostatsListRequestCustomMetadataHasValue {
	return &ThermostatsListRequestCustomMetadataHasValue{typeName: "string", String: value}
}

func NewThermostatsListRequestCustomMetadataHasValueFromBoolean(value bool) *ThermostatsListRequestCustomMetadataHasValue {
	return &ThermostatsListRequestCustomMetadataHasValue{typeName: "boolean", Boolean: value}
}

func NewThermostatsListRequestCustomMetadataHasValueFromStringOptional(value *string) *ThermostatsListRequestCustomMetadataHasValue {
	return &ThermostatsListRequestCustomMetadataHasValue{typeName: "stringOptional", StringOptional: value}
}

func (t *ThermostatsListRequestCustomMetadataHasValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		t.typeName = "string"
		t.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		t.typeName = "boolean"
		t.Boolean = valueBoolean
		return nil
	}
	var valueStringOptional *string
	if err := json.Unmarshal(data, &valueStringOptional); err == nil {
		t.typeName = "stringOptional"
		t.StringOptional = valueStringOptional
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t ThermostatsListRequestCustomMetadataHasValue) MarshalJSON() ([]byte, error) {
	switch t.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", t.typeName, t)
	case "string":
		return json.Marshal(t.String)
	case "boolean":
		return json.Marshal(t.Boolean)
	case "stringOptional":
		return json.Marshal(t.StringOptional)
	}
}

type ThermostatsListRequestCustomMetadataHasValueVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
	VisitStringOptional(*string) error
}

func (t *ThermostatsListRequestCustomMetadataHasValue) Accept(visitor ThermostatsListRequestCustomMetadataHasValueVisitor) error {
	switch t.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", t.typeName, t)
	case "string":
		return visitor.VisitString(t.String)
	case "boolean":
		return visitor.VisitBoolean(t.Boolean)
	case "stringOptional":
		return visitor.VisitStringOptional(t.StringOptional)
	}
}

type ThermostatsListResponse struct {
	Thermostats []*Device `json:"thermostats,omitempty"`
	Ok          bool      `json:"ok"`

	_rawJSON json.RawMessage
}

func (t *ThermostatsListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ThermostatsListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThermostatsListResponse(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThermostatsListResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type ThermostatsOffResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (t *ThermostatsOffResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ThermostatsOffResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThermostatsOffResponse(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThermostatsOffResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type ThermostatsSetFanModeRequestFanModeSetting string

const (
	ThermostatsSetFanModeRequestFanModeSettingAuto ThermostatsSetFanModeRequestFanModeSetting = "auto"
	ThermostatsSetFanModeRequestFanModeSettingOn   ThermostatsSetFanModeRequestFanModeSetting = "on"
)

func NewThermostatsSetFanModeRequestFanModeSettingFromString(s string) (ThermostatsSetFanModeRequestFanModeSetting, error) {
	switch s {
	case "auto":
		return ThermostatsSetFanModeRequestFanModeSettingAuto, nil
	case "on":
		return ThermostatsSetFanModeRequestFanModeSettingOn, nil
	}
	var t ThermostatsSetFanModeRequestFanModeSetting
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t ThermostatsSetFanModeRequestFanModeSetting) Ptr() *ThermostatsSetFanModeRequestFanModeSetting {
	return &t
}

type ThermostatsSetFanModeResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (t *ThermostatsSetFanModeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ThermostatsSetFanModeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThermostatsSetFanModeResponse(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThermostatsSetFanModeResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type ThermostatsUpdateRequestDefaultClimateSetting struct {
	AutomaticHeatingEnabled   *bool                                                         `json:"automatic_heating_enabled,omitempty"`
	AutomaticCoolingEnabled   *bool                                                         `json:"automatic_cooling_enabled,omitempty"`
	HvacModeSetting           *ThermostatsUpdateRequestDefaultClimateSettingHvacModeSetting `json:"hvac_mode_setting,omitempty"`
	CoolingSetPointCelsius    *float64                                                      `json:"cooling_set_point_celsius,omitempty"`
	HeatingSetPointCelsius    *float64                                                      `json:"heating_set_point_celsius,omitempty"`
	CoolingSetPointFahrenheit *float64                                                      `json:"cooling_set_point_fahrenheit,omitempty"`
	HeatingSetPointFahrenheit *float64                                                      `json:"heating_set_point_fahrenheit,omitempty"`
	ManualOverrideAllowed     *bool                                                         `json:"manual_override_allowed,omitempty"`

	_rawJSON json.RawMessage
}

func (t *ThermostatsUpdateRequestDefaultClimateSetting) UnmarshalJSON(data []byte) error {
	type unmarshaler ThermostatsUpdateRequestDefaultClimateSetting
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThermostatsUpdateRequestDefaultClimateSetting(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThermostatsUpdateRequestDefaultClimateSetting) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type ThermostatsUpdateResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (t *ThermostatsUpdateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ThermostatsUpdateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThermostatsUpdateResponse(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThermostatsUpdateResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type ThermostatsUpdateRequest struct {
	DeviceId              string                                         `json:"device_id"`
	DefaultClimateSetting *ThermostatsUpdateRequestDefaultClimateSetting `json:"default_climate_setting,omitempty"`
}
