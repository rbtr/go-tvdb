// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUserRatingsParams creates a new GetUserRatingsParams object
// with the default values initialized.
func NewGetUserRatingsParams() *GetUserRatingsParams {

	return &GetUserRatingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRatingsParamsWithTimeout creates a new GetUserRatingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserRatingsParamsWithTimeout(timeout time.Duration) *GetUserRatingsParams {

	return &GetUserRatingsParams{

		timeout: timeout,
	}
}

// NewGetUserRatingsParamsWithContext creates a new GetUserRatingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserRatingsParamsWithContext(ctx context.Context) *GetUserRatingsParams {

	return &GetUserRatingsParams{

		Context: ctx,
	}
}

// NewGetUserRatingsParamsWithHTTPClient creates a new GetUserRatingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserRatingsParamsWithHTTPClient(client *http.Client) *GetUserRatingsParams {

	return &GetUserRatingsParams{
		HTTPClient: client,
	}
}

/*GetUserRatingsParams contains all the parameters to send to the API endpoint
for the get user ratings operation typically these are written to a http.Request
*/
type GetUserRatingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user ratings params
func (o *GetUserRatingsParams) WithTimeout(timeout time.Duration) *GetUserRatingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user ratings params
func (o *GetUserRatingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user ratings params
func (o *GetUserRatingsParams) WithContext(ctx context.Context) *GetUserRatingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user ratings params
func (o *GetUserRatingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user ratings params
func (o *GetUserRatingsParams) WithHTTPClient(client *http.Client) *GetUserRatingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user ratings params
func (o *GetUserRatingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRatingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
