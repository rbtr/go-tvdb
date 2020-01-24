// Code generated by go-swagger; DO NOT EDIT.

package languages

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

// NewGetLanguagesIDParams creates a new GetLanguagesIDParams object
// with the default values initialized.
func NewGetLanguagesIDParams() *GetLanguagesIDParams {
	var ()
	return &GetLanguagesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLanguagesIDParamsWithTimeout creates a new GetLanguagesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLanguagesIDParamsWithTimeout(timeout time.Duration) *GetLanguagesIDParams {
	var ()
	return &GetLanguagesIDParams{

		timeout: timeout,
	}
}

// NewGetLanguagesIDParamsWithContext creates a new GetLanguagesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLanguagesIDParamsWithContext(ctx context.Context) *GetLanguagesIDParams {
	var ()
	return &GetLanguagesIDParams{

		Context: ctx,
	}
}

// NewGetLanguagesIDParamsWithHTTPClient creates a new GetLanguagesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLanguagesIDParamsWithHTTPClient(client *http.Client) *GetLanguagesIDParams {
	var ()
	return &GetLanguagesIDParams{
		HTTPClient: client,
	}
}

/*GetLanguagesIDParams contains all the parameters to send to the API endpoint
for the get languages ID operation typically these are written to a http.Request
*/
type GetLanguagesIDParams struct {

	/*ID
	  ID of the language

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get languages ID params
func (o *GetLanguagesIDParams) WithTimeout(timeout time.Duration) *GetLanguagesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get languages ID params
func (o *GetLanguagesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get languages ID params
func (o *GetLanguagesIDParams) WithContext(ctx context.Context) *GetLanguagesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get languages ID params
func (o *GetLanguagesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get languages ID params
func (o *GetLanguagesIDParams) WithHTTPClient(client *http.Client) *GetLanguagesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get languages ID params
func (o *GetLanguagesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get languages ID params
func (o *GetLanguagesIDParams) WithID(id string) *GetLanguagesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get languages ID params
func (o *GetLanguagesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLanguagesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
