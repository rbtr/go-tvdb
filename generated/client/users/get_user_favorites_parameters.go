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

// NewGetUserFavoritesParams creates a new GetUserFavoritesParams object
// with the default values initialized.
func NewGetUserFavoritesParams() *GetUserFavoritesParams {

	return &GetUserFavoritesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserFavoritesParamsWithTimeout creates a new GetUserFavoritesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserFavoritesParamsWithTimeout(timeout time.Duration) *GetUserFavoritesParams {

	return &GetUserFavoritesParams{

		timeout: timeout,
	}
}

// NewGetUserFavoritesParamsWithContext creates a new GetUserFavoritesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserFavoritesParamsWithContext(ctx context.Context) *GetUserFavoritesParams {

	return &GetUserFavoritesParams{

		Context: ctx,
	}
}

// NewGetUserFavoritesParamsWithHTTPClient creates a new GetUserFavoritesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserFavoritesParamsWithHTTPClient(client *http.Client) *GetUserFavoritesParams {

	return &GetUserFavoritesParams{
		HTTPClient: client,
	}
}

/*GetUserFavoritesParams contains all the parameters to send to the API endpoint
for the get user favorites operation typically these are written to a http.Request
*/
type GetUserFavoritesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user favorites params
func (o *GetUserFavoritesParams) WithTimeout(timeout time.Duration) *GetUserFavoritesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user favorites params
func (o *GetUserFavoritesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user favorites params
func (o *GetUserFavoritesParams) WithContext(ctx context.Context) *GetUserFavoritesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user favorites params
func (o *GetUserFavoritesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user favorites params
func (o *GetUserFavoritesParams) WithHTTPClient(client *http.Client) *GetUserFavoritesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user favorites params
func (o *GetUserFavoritesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserFavoritesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
