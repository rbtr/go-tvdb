// Code generated by go-swagger; DO NOT EDIT.

package updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new updates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for updates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetUpdatedQuery Returns an array of series that have changed in a maximum of one week blocks since the provided `fromTime`.


The user may specify a `toTime` to grab results for less than a week. Any timespan larger than a week will be reduced down to one week automatically.
*/
func (a *Client) GetUpdatedQuery(params *GetUpdatedQueryParams, authInfo runtime.ClientAuthInfoWriter) (*GetUpdatedQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUpdatedQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUpdatedQuery",
		Method:             "GET",
		PathPattern:        "/updated/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUpdatedQueryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUpdatedQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUpdatedQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUpdatedQueryParams Returns an array of valid query keys for the `/updated/query/params` route.
*/
func (a *Client) GetUpdatedQueryParams(params *GetUpdatedQueryParamsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUpdatedQueryParamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUpdatedQueryParamsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUpdatedQueryParams",
		Method:             "GET",
		PathPattern:        "/updated/query/params",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUpdatedQueryParamsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUpdatedQueryParamsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUpdatedQueryParams: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
