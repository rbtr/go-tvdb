// Code generated by go-swagger; DO NOT EDIT.

package episodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEpisodesIDParams creates a new GetEpisodesIDParams object
// with the default values initialized.
func NewGetEpisodesIDParams() *GetEpisodesIDParams {
	var ()
	return &GetEpisodesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEpisodesIDParamsWithTimeout creates a new GetEpisodesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEpisodesIDParamsWithTimeout(timeout time.Duration) *GetEpisodesIDParams {
	var ()
	return &GetEpisodesIDParams{

		timeout: timeout,
	}
}

// NewGetEpisodesIDParamsWithContext creates a new GetEpisodesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEpisodesIDParamsWithContext(ctx context.Context) *GetEpisodesIDParams {
	var ()
	return &GetEpisodesIDParams{

		Context: ctx,
	}
}

// NewGetEpisodesIDParamsWithHTTPClient creates a new GetEpisodesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEpisodesIDParamsWithHTTPClient(client *http.Client) *GetEpisodesIDParams {
	var ()
	return &GetEpisodesIDParams{
		HTTPClient: client,
	}
}

/*GetEpisodesIDParams contains all the parameters to send to the API endpoint
for the get episodes ID operation typically these are written to a http.Request
*/
type GetEpisodesIDParams struct {

	/*AcceptLanguage
	  Records are returned with the some fields in the desired language, if it exists. If there is no translation for the given language, then the record is still returned but with empty values for the translated fields.

	*/
	AcceptLanguage *string
	/*ID
	  ID of the episode

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get episodes ID params
func (o *GetEpisodesIDParams) WithTimeout(timeout time.Duration) *GetEpisodesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get episodes ID params
func (o *GetEpisodesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get episodes ID params
func (o *GetEpisodesIDParams) WithContext(ctx context.Context) *GetEpisodesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get episodes ID params
func (o *GetEpisodesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get episodes ID params
func (o *GetEpisodesIDParams) WithHTTPClient(client *http.Client) *GetEpisodesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get episodes ID params
func (o *GetEpisodesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get episodes ID params
func (o *GetEpisodesIDParams) WithAcceptLanguage(acceptLanguage *string) *GetEpisodesIDParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get episodes ID params
func (o *GetEpisodesIDParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithID adds the id to the get episodes ID params
func (o *GetEpisodesIDParams) WithID(id int64) *GetEpisodesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get episodes ID params
func (o *GetEpisodesIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEpisodesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
