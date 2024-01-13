// This file was auto-generated by Fern from our API Definition.

package devices

import (
	json "encoding/json"
	fmt "fmt"
	seamapigo "github.com/seamapi/go"
	core "github.com/seamapi/go/core"
	time "time"
)

type UnmanagedGetRequest struct {
	DeviceId *string `json:"device_id,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type UnmanagedListRequest struct {
	// List all devices owned by this connected account
	ConnectedAccountId  *string                 `json:"connected_account_id,omitempty"`
	ConnectedAccountIds []string                `json:"connected_account_ids,omitempty"`
	ConnectWebviewId    *string                 `json:"connect_webview_id,omitempty"`
	DeviceType          *seamapigo.DeviceType   `json:"device_type,omitempty"`
	DeviceTypes         []seamapigo.DeviceType  `json:"device_types,omitempty"`
	Manufacturer        *seamapigo.Manufacturer `json:"manufacturer,omitempty"`
	DeviceIds           []string                `json:"device_ids,omitempty"`
	Limit               *float64                `json:"limit,omitempty"`
	CreatedBefore       *time.Time              `json:"created_before,omitempty"`
	UserIdentifierKey   *string                 `json:"user_identifier_key,omitempty"`
}

type UnmanagedGetResponse struct {
	Device *seamapigo.UnmanagedDevice `json:"device,omitempty"`
	Ok     bool                       `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UnmanagedGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UnmanagedGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UnmanagedGetResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UnmanagedGetResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UnmanagedListResponse struct {
	Devices []*seamapigo.UnmanagedDevice `json:"devices,omitempty"`
	Ok      bool                         `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UnmanagedListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UnmanagedListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UnmanagedListResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UnmanagedListResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UnmanagedUpdateResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UnmanagedUpdateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UnmanagedUpdateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UnmanagedUpdateResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UnmanagedUpdateResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UnmanagedUpdateRequest struct {
	DeviceId  string `json:"device_id"`
	IsManaged bool   `json:"is_managed"`
}
