// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/seamapi/go/core"
	noisethresholds "github.com/seamapi/go/noisesensors/noisethresholds"
	simulate "github.com/seamapi/go/noisesensors/simulate"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header

	NoiseThresholds *noisethresholds.Client
	Simulate        *simulate.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:         options.BaseURL,
		httpClient:      options.HTTPClient,
		header:          options.ToHeader(),
		NoiseThresholds: noisethresholds.NewClient(opts...),
		Simulate:        simulate.NewClient(opts...),
	}
}
