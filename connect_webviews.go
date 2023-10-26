// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type ConnectWebviewsCreateRequest struct {
	DeviceSelectionMode      *ConnectWebviewsCreateRequestDeviceSelectionMode            `json:"device_selection_mode,omitempty"`
	CustomRedirectUrl        *string                                                     `json:"custom_redirect_url,omitempty"`
	CustomRedirectFailureUrl *string                                                     `json:"custom_redirect_failure_url,omitempty"`
	AcceptedProviders        []ConnectWebviewsCreateRequestAcceptedProvidersItem         `json:"accepted_providers,omitempty"`
	ProviderCategory         *ConnectWebviewsCreateRequestProviderCategory               `json:"provider_category,omitempty"`
	CustomMetadata           map[string]*ConnectWebviewsCreateRequestCustomMetadataValue `json:"custom_metadata,omitempty"`
}

type ConnectWebviewsDeleteRequest struct {
	ConnectWebviewId string `json:"connect_webview_id"`
}

type ConnectWebviewsGetRequest struct {
	ConnectWebviewId string `json:"connect_webview_id"`
}

type ConnectWebviewsCreateRequestAcceptedProvidersItem string

const (
	ConnectWebviewsCreateRequestAcceptedProvidersItemAkuvox         ConnectWebviewsCreateRequestAcceptedProvidersItem = "akuvox"
	ConnectWebviewsCreateRequestAcceptedProvidersItemAugust         ConnectWebviewsCreateRequestAcceptedProvidersItem = "august"
	ConnectWebviewsCreateRequestAcceptedProvidersItemAvigilonAlta   ConnectWebviewsCreateRequestAcceptedProvidersItem = "avigilon_alta"
	ConnectWebviewsCreateRequestAcceptedProvidersItemBrivo          ConnectWebviewsCreateRequestAcceptedProvidersItem = "brivo"
	ConnectWebviewsCreateRequestAcceptedProvidersItemButterflymx    ConnectWebviewsCreateRequestAcceptedProvidersItem = "butterflymx"
	ConnectWebviewsCreateRequestAcceptedProvidersItemSchlage        ConnectWebviewsCreateRequestAcceptedProvidersItem = "schlage"
	ConnectWebviewsCreateRequestAcceptedProvidersItemSmartthings    ConnectWebviewsCreateRequestAcceptedProvidersItem = "smartthings"
	ConnectWebviewsCreateRequestAcceptedProvidersItemYale           ConnectWebviewsCreateRequestAcceptedProvidersItem = "yale"
	ConnectWebviewsCreateRequestAcceptedProvidersItemGenie          ConnectWebviewsCreateRequestAcceptedProvidersItem = "genie"
	ConnectWebviewsCreateRequestAcceptedProvidersItemDoorking       ConnectWebviewsCreateRequestAcceptedProvidersItem = "doorking"
	ConnectWebviewsCreateRequestAcceptedProvidersItemSalto          ConnectWebviewsCreateRequestAcceptedProvidersItem = "salto"
	ConnectWebviewsCreateRequestAcceptedProvidersItemLockly         ConnectWebviewsCreateRequestAcceptedProvidersItem = "lockly"
	ConnectWebviewsCreateRequestAcceptedProvidersItemTtlock         ConnectWebviewsCreateRequestAcceptedProvidersItem = "ttlock"
	ConnectWebviewsCreateRequestAcceptedProvidersItemLinear         ConnectWebviewsCreateRequestAcceptedProvidersItem = "linear"
	ConnectWebviewsCreateRequestAcceptedProvidersItemNoiseaware     ConnectWebviewsCreateRequestAcceptedProvidersItem = "noiseaware"
	ConnectWebviewsCreateRequestAcceptedProvidersItemNuki           ConnectWebviewsCreateRequestAcceptedProvidersItem = "nuki"
	ConnectWebviewsCreateRequestAcceptedProvidersItemSeamRelayAdmin ConnectWebviewsCreateRequestAcceptedProvidersItem = "seam_relay_admin"
	ConnectWebviewsCreateRequestAcceptedProvidersItemIgloo          ConnectWebviewsCreateRequestAcceptedProvidersItem = "igloo"
	ConnectWebviewsCreateRequestAcceptedProvidersItemKwikset        ConnectWebviewsCreateRequestAcceptedProvidersItem = "kwikset"
	ConnectWebviewsCreateRequestAcceptedProvidersItemMinut          ConnectWebviewsCreateRequestAcceptedProvidersItem = "minut"
	ConnectWebviewsCreateRequestAcceptedProvidersItemMy2N           ConnectWebviewsCreateRequestAcceptedProvidersItem = "my_2n"
	ConnectWebviewsCreateRequestAcceptedProvidersItemControlbyweb   ConnectWebviewsCreateRequestAcceptedProvidersItem = "controlbyweb"
	ConnectWebviewsCreateRequestAcceptedProvidersItemNest           ConnectWebviewsCreateRequestAcceptedProvidersItem = "nest"
	ConnectWebviewsCreateRequestAcceptedProvidersItemIgloohome      ConnectWebviewsCreateRequestAcceptedProvidersItem = "igloohome"
	ConnectWebviewsCreateRequestAcceptedProvidersItemEcobee         ConnectWebviewsCreateRequestAcceptedProvidersItem = "ecobee"
	ConnectWebviewsCreateRequestAcceptedProvidersItemHubitat        ConnectWebviewsCreateRequestAcceptedProvidersItem = "hubitat"
	ConnectWebviewsCreateRequestAcceptedProvidersItemYaleAccess     ConnectWebviewsCreateRequestAcceptedProvidersItem = "yale_access"
)

func NewConnectWebviewsCreateRequestAcceptedProvidersItemFromString(s string) (ConnectWebviewsCreateRequestAcceptedProvidersItem, error) {
	switch s {
	case "akuvox":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemAkuvox, nil
	case "august":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemAugust, nil
	case "avigilon_alta":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemAvigilonAlta, nil
	case "brivo":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemBrivo, nil
	case "butterflymx":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemButterflymx, nil
	case "schlage":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemSchlage, nil
	case "smartthings":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemSmartthings, nil
	case "yale":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemYale, nil
	case "genie":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemGenie, nil
	case "doorking":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemDoorking, nil
	case "salto":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemSalto, nil
	case "lockly":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemLockly, nil
	case "ttlock":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemTtlock, nil
	case "linear":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemLinear, nil
	case "noiseaware":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemNoiseaware, nil
	case "nuki":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemNuki, nil
	case "seam_relay_admin":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemSeamRelayAdmin, nil
	case "igloo":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemIgloo, nil
	case "kwikset":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemKwikset, nil
	case "minut":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemMinut, nil
	case "my_2n":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemMy2N, nil
	case "controlbyweb":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemControlbyweb, nil
	case "nest":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemNest, nil
	case "igloohome":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemIgloohome, nil
	case "ecobee":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemEcobee, nil
	case "hubitat":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemHubitat, nil
	case "yale_access":
		return ConnectWebviewsCreateRequestAcceptedProvidersItemYaleAccess, nil
	}
	var t ConnectWebviewsCreateRequestAcceptedProvidersItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c ConnectWebviewsCreateRequestAcceptedProvidersItem) Ptr() *ConnectWebviewsCreateRequestAcceptedProvidersItem {
	return &c
}

type ConnectWebviewsCreateRequestCustomMetadataValue struct {
	typeName       string
	String         string
	Double         float64
	StringOptional *string
	Boolean        bool
}

func NewConnectWebviewsCreateRequestCustomMetadataValueFromString(value string) *ConnectWebviewsCreateRequestCustomMetadataValue {
	return &ConnectWebviewsCreateRequestCustomMetadataValue{typeName: "string", String: value}
}

func NewConnectWebviewsCreateRequestCustomMetadataValueFromDouble(value float64) *ConnectWebviewsCreateRequestCustomMetadataValue {
	return &ConnectWebviewsCreateRequestCustomMetadataValue{typeName: "double", Double: value}
}

func NewConnectWebviewsCreateRequestCustomMetadataValueFromStringOptional(value *string) *ConnectWebviewsCreateRequestCustomMetadataValue {
	return &ConnectWebviewsCreateRequestCustomMetadataValue{typeName: "stringOptional", StringOptional: value}
}

func NewConnectWebviewsCreateRequestCustomMetadataValueFromBoolean(value bool) *ConnectWebviewsCreateRequestCustomMetadataValue {
	return &ConnectWebviewsCreateRequestCustomMetadataValue{typeName: "boolean", Boolean: value}
}

func (c *ConnectWebviewsCreateRequestCustomMetadataValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typeName = "string"
		c.String = valueString
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		c.typeName = "double"
		c.Double = valueDouble
		return nil
	}
	var valueStringOptional *string
	if err := json.Unmarshal(data, &valueStringOptional); err == nil {
		c.typeName = "stringOptional"
		c.StringOptional = valueStringOptional
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

func (c ConnectWebviewsCreateRequestCustomMetadataValue) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return json.Marshal(c.String)
	case "double":
		return json.Marshal(c.Double)
	case "stringOptional":
		return json.Marshal(c.StringOptional)
	case "boolean":
		return json.Marshal(c.Boolean)
	}
}

type ConnectWebviewsCreateRequestCustomMetadataValueVisitor interface {
	VisitString(string) error
	VisitDouble(float64) error
	VisitStringOptional(*string) error
	VisitBoolean(bool) error
}

func (c *ConnectWebviewsCreateRequestCustomMetadataValue) Accept(visitor ConnectWebviewsCreateRequestCustomMetadataValueVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return visitor.VisitString(c.String)
	case "double":
		return visitor.VisitDouble(c.Double)
	case "stringOptional":
		return visitor.VisitStringOptional(c.StringOptional)
	case "boolean":
		return visitor.VisitBoolean(c.Boolean)
	}
}

type ConnectWebviewsCreateRequestDeviceSelectionMode string

const (
	ConnectWebviewsCreateRequestDeviceSelectionModeNone     ConnectWebviewsCreateRequestDeviceSelectionMode = "none"
	ConnectWebviewsCreateRequestDeviceSelectionModeSingle   ConnectWebviewsCreateRequestDeviceSelectionMode = "single"
	ConnectWebviewsCreateRequestDeviceSelectionModeMultiple ConnectWebviewsCreateRequestDeviceSelectionMode = "multiple"
)

func NewConnectWebviewsCreateRequestDeviceSelectionModeFromString(s string) (ConnectWebviewsCreateRequestDeviceSelectionMode, error) {
	switch s {
	case "none":
		return ConnectWebviewsCreateRequestDeviceSelectionModeNone, nil
	case "single":
		return ConnectWebviewsCreateRequestDeviceSelectionModeSingle, nil
	case "multiple":
		return ConnectWebviewsCreateRequestDeviceSelectionModeMultiple, nil
	}
	var t ConnectWebviewsCreateRequestDeviceSelectionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c ConnectWebviewsCreateRequestDeviceSelectionMode) Ptr() *ConnectWebviewsCreateRequestDeviceSelectionMode {
	return &c
}

type ConnectWebviewsCreateRequestProviderCategory string

const (
	ConnectWebviewsCreateRequestProviderCategoryStable             ConnectWebviewsCreateRequestProviderCategory = "stable"
	ConnectWebviewsCreateRequestProviderCategoryConsumerSmartlocks ConnectWebviewsCreateRequestProviderCategory = "consumer_smartlocks"
	ConnectWebviewsCreateRequestProviderCategoryInternalBeta       ConnectWebviewsCreateRequestProviderCategory = "internal_beta"
)

func NewConnectWebviewsCreateRequestProviderCategoryFromString(s string) (ConnectWebviewsCreateRequestProviderCategory, error) {
	switch s {
	case "stable":
		return ConnectWebviewsCreateRequestProviderCategoryStable, nil
	case "consumer_smartlocks":
		return ConnectWebviewsCreateRequestProviderCategoryConsumerSmartlocks, nil
	case "internal_beta":
		return ConnectWebviewsCreateRequestProviderCategoryInternalBeta, nil
	}
	var t ConnectWebviewsCreateRequestProviderCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c ConnectWebviewsCreateRequestProviderCategory) Ptr() *ConnectWebviewsCreateRequestProviderCategory {
	return &c
}

type ConnectWebviewsCreateResponse struct {
	ConnectWebview *ConnectWebview `json:"connect_webview,omitempty"`
	Ok             bool            `json:"ok"`
}

type ConnectWebviewsDeleteResponse struct {
	Ok bool `json:"ok"`
}

type ConnectWebviewsGetResponse struct {
	ConnectWebview *ConnectWebview `json:"connect_webview,omitempty"`
	Ok             bool            `json:"ok"`
}

type ConnectWebviewsListResponse struct {
	ConnectWebviews []*ConnectWebview `json:"connect_webviews,omitempty"`
	Ok              bool              `json:"ok"`
}

type ConnectWebviewsViewRequest struct {
	ConnectWebviewId string `json:"-"`
	AuthToken        string `json:"-"`
}
