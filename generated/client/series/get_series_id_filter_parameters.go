// Code generated by go-swagger; DO NOT EDIT.

package series

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

// NewGetSeriesIDFilterParams creates a new GetSeriesIDFilterParams object
// with the default values initialized.
func NewGetSeriesIDFilterParams() *GetSeriesIDFilterParams {
	var ()
	return &GetSeriesIDFilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSeriesIDFilterParamsWithTimeout creates a new GetSeriesIDFilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSeriesIDFilterParamsWithTimeout(timeout time.Duration) *GetSeriesIDFilterParams {
	var ()
	return &GetSeriesIDFilterParams{

		timeout: timeout,
	}
}

// NewGetSeriesIDFilterParamsWithContext creates a new GetSeriesIDFilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSeriesIDFilterParamsWithContext(ctx context.Context) *GetSeriesIDFilterParams {
	var ()
	return &GetSeriesIDFilterParams{

		Context: ctx,
	}
}

// NewGetSeriesIDFilterParamsWithHTTPClient creates a new GetSeriesIDFilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSeriesIDFilterParamsWithHTTPClient(client *http.Client) *GetSeriesIDFilterParams {
	var ()
	return &GetSeriesIDFilterParams{
		HTTPClient: client,
	}
}

/*GetSeriesIDFilterParams contains all the parameters to send to the API endpoint
for the get series ID filter operation typically these are written to a http.Request
*/
type GetSeriesIDFilterParams struct {

	/*AcceptLanguage
	  Records are returned with the some fields in the desired language, if it exists. If there is no translation for the given language, then the record is still returned but with empty values for the translated fields.

	*/
	AcceptLanguage *string
	/*ID
	  ID of the series

	*/
	ID int64
	/*Keys
	  Comma-separated list of keys to filter by

	*/
	Keys string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get series ID filter params
func (o *GetSeriesIDFilterParams) WithTimeout(timeout time.Duration) *GetSeriesIDFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get series ID filter params
func (o *GetSeriesIDFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get series ID filter params
func (o *GetSeriesIDFilterParams) WithContext(ctx context.Context) *GetSeriesIDFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get series ID filter params
func (o *GetSeriesIDFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get series ID filter params
func (o *GetSeriesIDFilterParams) WithHTTPClient(client *http.Client) *GetSeriesIDFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get series ID filter params
func (o *GetSeriesIDFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get series ID filter params
func (o *GetSeriesIDFilterParams) WithAcceptLanguage(acceptLanguage *string) *GetSeriesIDFilterParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get series ID filter params
func (o *GetSeriesIDFilterParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithID adds the id to the get series ID filter params
func (o *GetSeriesIDFilterParams) WithID(id int64) *GetSeriesIDFilterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get series ID filter params
func (o *GetSeriesIDFilterParams) SetID(id int64) {
	o.ID = id
}

// WithKeys adds the keys to the get series ID filter params
func (o *GetSeriesIDFilterParams) WithKeys(keys string) *GetSeriesIDFilterParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the get series ID filter params
func (o *GetSeriesIDFilterParams) SetKeys(keys string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *GetSeriesIDFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param keys
	qrKeys := o.Keys
	qKeys := qrKeys
	if qKeys != "" {
		if err := r.SetQueryParam("keys", qKeys); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
