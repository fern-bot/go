// This file was auto-generated by Fern from our API Definition.

package client

import (
	accessgroups "github.com/seamapi/go/acs/accessgroups"
	credentialpools "github.com/seamapi/go/acs/credentialpools"
	credentialprovisioningautomations "github.com/seamapi/go/acs/credentialprovisioningautomations"
	credentials "github.com/seamapi/go/acs/credentials"
	entrances "github.com/seamapi/go/acs/entrances"
	systems "github.com/seamapi/go/acs/systems"
	users "github.com/seamapi/go/acs/users"
	core "github.com/seamapi/go/core"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	AccessGroups                      *accessgroups.Client
	CredentialPools                   *credentialpools.Client
	CredentialProvisioningAutomations *credentialprovisioningautomations.Client
	Credentials                       *credentials.Client
	Entrances                         *entrances.Client
	Systems                           *systems.Client
	Users                             *users.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:                           options.BaseURL,
		caller:                            core.NewCaller(options.HTTPClient),
		header:                            options.ToHeader(),
		AccessGroups:                      accessgroups.NewClient(opts...),
		CredentialPools:                   credentialpools.NewClient(opts...),
		CredentialProvisioningAutomations: credentialprovisioningautomations.NewClient(opts...),
		Credentials:                       credentials.NewClient(opts...),
		Entrances:                         entrances.NewClient(opts...),
		Systems:                           systems.NewClient(opts...),
		Users:                             users.NewClient(opts...),
	}
}
