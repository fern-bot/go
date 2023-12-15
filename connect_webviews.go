// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
)

type ConnectWebviewsCreateRequest struct {
	DeviceSelectionMode           *SelectionMode                  `json:"device_selection_mode,omitempty"`
	CustomRedirectUrl             *string                         `json:"custom_redirect_url,omitempty"`
	CustomRedirectFailureUrl      *string                         `json:"custom_redirect_failure_url,omitempty"`
	AcceptedProviders             []AcceptedProvider              `json:"accepted_providers,omitempty"`
	ProviderCategory              *ProviderCategory               `json:"provider_category,omitempty"`
	CustomMetadata                map[string]*CustomMetadataValue `json:"custom_metadata,omitempty"`
	AutomaticallyManageNewDevices *bool                           `json:"automatically_manage_new_devices,omitempty"`
	WaitForDeviceCreation         *bool                           `json:"wait_for_device_creation,omitempty"`
}

type ConnectWebviewsDeleteRequest struct {
	ConnectWebviewId string `json:"connect_webview_id"`
}

type ConnectWebviewsGetRequest struct {
	ConnectWebviewId string `json:"connect_webview_id"`
}

type AcceptedProvider string

const (
	AcceptedProviderAkuvox         AcceptedProvider = "akuvox"
	AcceptedProviderAugust         AcceptedProvider = "august"
	AcceptedProviderAvigilonAlta   AcceptedProvider = "avigilon_alta"
	AcceptedProviderBrivo          AcceptedProvider = "brivo"
	AcceptedProviderButterflymx    AcceptedProvider = "butterflymx"
	AcceptedProviderSchlage        AcceptedProvider = "schlage"
	AcceptedProviderSmartthings    AcceptedProvider = "smartthings"
	AcceptedProviderYale           AcceptedProvider = "yale"
	AcceptedProviderGenie          AcceptedProvider = "genie"
	AcceptedProviderDoorking       AcceptedProvider = "doorking"
	AcceptedProviderSalto          AcceptedProvider = "salto"
	AcceptedProviderLockly         AcceptedProvider = "lockly"
	AcceptedProviderTtlock         AcceptedProvider = "ttlock"
	AcceptedProviderLinear         AcceptedProvider = "linear"
	AcceptedProviderNoiseaware     AcceptedProvider = "noiseaware"
	AcceptedProviderNuki           AcceptedProvider = "nuki"
	AcceptedProviderSeamRelayAdmin AcceptedProvider = "seam_relay_admin"
	AcceptedProviderIgloo          AcceptedProvider = "igloo"
	AcceptedProviderKwikset        AcceptedProvider = "kwikset"
	AcceptedProviderMinut          AcceptedProvider = "minut"
	AcceptedProviderMy2N           AcceptedProvider = "my_2n"
	AcceptedProviderControlbyweb   AcceptedProvider = "controlbyweb"
	AcceptedProviderNest           AcceptedProvider = "nest"
	AcceptedProviderIgloohome      AcceptedProvider = "igloohome"
	AcceptedProviderEcobee         AcceptedProvider = "ecobee"
	AcceptedProviderHubitat        AcceptedProvider = "hubitat"
	AcceptedProviderYaleAccess     AcceptedProvider = "yale_access"
)

func NewAcceptedProviderFromString(s string) (AcceptedProvider, error) {
	switch s {
	case "akuvox":
		return AcceptedProviderAkuvox, nil
	case "august":
		return AcceptedProviderAugust, nil
	case "avigilon_alta":
		return AcceptedProviderAvigilonAlta, nil
	case "brivo":
		return AcceptedProviderBrivo, nil
	case "butterflymx":
		return AcceptedProviderButterflymx, nil
	case "schlage":
		return AcceptedProviderSchlage, nil
	case "smartthings":
		return AcceptedProviderSmartthings, nil
	case "yale":
		return AcceptedProviderYale, nil
	case "genie":
		return AcceptedProviderGenie, nil
	case "doorking":
		return AcceptedProviderDoorking, nil
	case "salto":
		return AcceptedProviderSalto, nil
	case "lockly":
		return AcceptedProviderLockly, nil
	case "ttlock":
		return AcceptedProviderTtlock, nil
	case "linear":
		return AcceptedProviderLinear, nil
	case "noiseaware":
		return AcceptedProviderNoiseaware, nil
	case "nuki":
		return AcceptedProviderNuki, nil
	case "seam_relay_admin":
		return AcceptedProviderSeamRelayAdmin, nil
	case "igloo":
		return AcceptedProviderIgloo, nil
	case "kwikset":
		return AcceptedProviderKwikset, nil
	case "minut":
		return AcceptedProviderMinut, nil
	case "my_2n":
		return AcceptedProviderMy2N, nil
	case "controlbyweb":
		return AcceptedProviderControlbyweb, nil
	case "nest":
		return AcceptedProviderNest, nil
	case "igloohome":
		return AcceptedProviderIgloohome, nil
	case "ecobee":
		return AcceptedProviderEcobee, nil
	case "hubitat":
		return AcceptedProviderHubitat, nil
	case "yale_access":
		return AcceptedProviderYaleAccess, nil
	}
	var t AcceptedProvider
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AcceptedProvider) Ptr() *AcceptedProvider {
	return &a
}

type ConnectWebviewsCreateResponse struct {
	ConnectWebview *ConnectWebview `json:"connect_webview,omitempty"`
	Ok             bool            `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *ConnectWebviewsCreateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectWebviewsCreateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectWebviewsCreateResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectWebviewsCreateResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type ConnectWebviewsDeleteResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *ConnectWebviewsDeleteResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectWebviewsDeleteResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectWebviewsDeleteResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectWebviewsDeleteResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type ConnectWebviewsGetResponse struct {
	ConnectWebview *ConnectWebview `json:"connect_webview,omitempty"`
	Ok             bool            `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *ConnectWebviewsGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectWebviewsGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectWebviewsGetResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectWebviewsGetResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type ConnectWebviewsListResponse struct {
	ConnectWebviews []*ConnectWebview `json:"connect_webviews,omitempty"`
	Ok              bool              `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *ConnectWebviewsListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectWebviewsListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectWebviewsListResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectWebviewsListResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CustomMetadataValue struct {
	typeName       string
	StringOptional *string
	Double         float64
	Boolean        bool
}

func NewCustomMetadataValueFromStringOptional(value *string) *CustomMetadataValue {
	return &CustomMetadataValue{typeName: "stringOptional", StringOptional: value}
}

func NewCustomMetadataValueFromDouble(value float64) *CustomMetadataValue {
	return &CustomMetadataValue{typeName: "double", Double: value}
}

func NewCustomMetadataValueFromBoolean(value bool) *CustomMetadataValue {
	return &CustomMetadataValue{typeName: "boolean", Boolean: value}
}

func (c *CustomMetadataValue) UnmarshalJSON(data []byte) error {
	var valueStringOptional *string
	if err := json.Unmarshal(data, &valueStringOptional); err == nil {
		c.typeName = "stringOptional"
		c.StringOptional = valueStringOptional
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		c.typeName = "double"
		c.Double = valueDouble
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		c.typeName = "boolean"
		c.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CustomMetadataValue) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "stringOptional":
		return json.Marshal(c.StringOptional)
	case "double":
		return json.Marshal(c.Double)
	case "boolean":
		return json.Marshal(c.Boolean)
	}
}

type CustomMetadataValueVisitor interface {
	VisitStringOptional(*string) error
	VisitDouble(float64) error
	VisitBoolean(bool) error
}

func (c *CustomMetadataValue) Accept(visitor CustomMetadataValueVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "stringOptional":
		return visitor.VisitStringOptional(c.StringOptional)
	case "double":
		return visitor.VisitDouble(c.Double)
	case "boolean":
		return visitor.VisitBoolean(c.Boolean)
	}
}

type ProviderCategory string

const (
	ProviderCategoryStable             ProviderCategory = "stable"
	ProviderCategoryConsumerSmartlocks ProviderCategory = "consumer_smartlocks"
	ProviderCategoryInternalBeta       ProviderCategory = "internal_beta"
)

func NewProviderCategoryFromString(s string) (ProviderCategory, error) {
	switch s {
	case "stable":
		return ProviderCategoryStable, nil
	case "consumer_smartlocks":
		return ProviderCategoryConsumerSmartlocks, nil
	case "internal_beta":
		return ProviderCategoryInternalBeta, nil
	}
	var t ProviderCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p ProviderCategory) Ptr() *ProviderCategory {
	return &p
}

type SelectionMode string

const (
	SelectionModeNone     SelectionMode = "none"
	SelectionModeSingle   SelectionMode = "single"
	SelectionModeMultiple SelectionMode = "multiple"
)

func NewSelectionModeFromString(s string) (SelectionMode, error) {
	switch s {
	case "none":
		return SelectionModeNone, nil
	case "single":
		return SelectionModeSingle, nil
	case "multiple":
		return SelectionModeMultiple, nil
	}
	var t SelectionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SelectionMode) Ptr() *SelectionMode {
	return &s
}

type ConnectWebviewsViewRequest struct {
	ConnectWebviewId string `json:"-"`
	AuthToken        string `json:"-"`
}
