// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/seamapi/go/core"
	http "net/http"
)

// WithBaseURL sets the client's base URL, overriding the
// default environment, if any.
func WithBaseURL(baseURL string) core.ClientOption {
	return func(opts *core.ClientOptions) {
		opts.BaseURL = baseURL
	}
}

// WithHTTPClient uses the given HTTPClient to issue all HTTP requests.
func WithHTTPClient(httpClient core.HTTPClient) core.ClientOption {
	return func(opts *core.ClientOptions) {
		opts.HTTPClient = httpClient
	}
}

// WithHTTPHeader adds the given http.Header to all requests
// issued by the client.
func WithHTTPHeader(httpHeader http.Header) core.ClientOption {
	return func(opts *core.ClientOptions) {
		// Clone the headers so they can't be modified after the option call.
		opts.HTTPHeader = httpHeader.Clone()
	}
}

// WithAuthApiKey sets the 'Authorization: Bearer <apiKey>' header on every request.
func WithAuthApiKey(apiKey string) core.ClientOption {
	return func(opts *core.ClientOptions) {
		opts.ApiKey = apiKey
	}
}
