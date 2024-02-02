// This file was auto-generated by Fern from our API Definition.

package acs

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
)

type EntrancesGetRequest struct {
	AcsEntranceId string `json:"acs_entrance_id"`
}

type EntrancesGrantAccessRequest struct {
	AcsEntranceId string `json:"acs_entrance_id"`
	AcsUserId     string `json:"acs_user_id"`
}

type EntrancesListRequest struct {
	AcsSystemId     *string `json:"acs_system_id,omitempty"`
	AcsCredentialId *string `json:"acs_credential_id,omitempty"`
}

type EntrancesListCredentialsWithAccessRequest struct {
	AcsEntranceId string   `json:"acs_entrance_id"`
	IncludeIf     []string `json:"include_if,omitempty"`
}

type EntrancesGetResponse struct {
	AcsEntrance *EntrancesGetResponseAcsEntrance `json:"acs_entrance,omitempty"`
	Ok          bool                             `json:"ok"`

	_rawJSON json.RawMessage
}

func (e *EntrancesGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EntrancesGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntrancesGetResponse(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntrancesGetResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EntrancesGrantAccessResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (e *EntrancesGrantAccessResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EntrancesGrantAccessResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntrancesGrantAccessResponse(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntrancesGrantAccessResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EntrancesListCredentialsWithAccessResponse struct {
	AcsCredentials []*EntrancesListCredentialsWithAccessResponseAcsCredentialsItem `json:"acs_credentials,omitempty"`
	Ok             bool                                                            `json:"ok"`

	_rawJSON json.RawMessage
}

func (e *EntrancesListCredentialsWithAccessResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EntrancesListCredentialsWithAccessResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntrancesListCredentialsWithAccessResponse(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntrancesListCredentialsWithAccessResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EntrancesListResponse struct {
	AcsEntrances []*EntrancesListResponseAcsEntrancesItem `json:"acs_entrances,omitempty"`
	Ok           bool                                     `json:"ok"`

	_rawJSON json.RawMessage
}

func (e *EntrancesListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EntrancesListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntrancesListResponse(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntrancesListResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}
