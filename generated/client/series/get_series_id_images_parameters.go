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

// NewGetSeriesIDImagesParams creates a new GetSeriesIDImagesParams object
// with the default values initialized.
func NewGetSeriesIDImagesParams() *GetSeriesIDImagesParams {
	var ()
	return &GetSeriesIDImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSeriesIDImagesParamsWithTimeout creates a new GetSeriesIDImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSeriesIDImagesParamsWithTimeout(timeout time.Duration) *GetSeriesIDImagesParams {
	var ()
	return &GetSeriesIDImagesParams{

		timeout: timeout,
	}
}

// NewGetSeriesIDImagesParamsWithContext creates a new GetSeriesIDImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSeriesIDImagesParamsWithContext(ctx context.Context) *GetSeriesIDImagesParams {
	var ()
	return &GetSeriesIDImagesParams{

		Context: ctx,
	}
}

// NewGetSeriesIDImagesParamsWithHTTPClient creates a new GetSeriesIDImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSeriesIDImagesParamsWithHTTPClient(client *http.Client) *GetSeriesIDImagesParams {
	var ()
	return &GetSeriesIDImagesParams{
		HTTPClient: client,
	}
}

/*GetSeriesIDImagesParams contains all the parameters to send to the API endpoint
for the get series ID images operation typically these are written to a http.Request
*/
type GetSeriesIDImagesParams struct {

	/*AcceptLanguage
	  Records are returned with the some fields in the desired language, if it exists. If there is no translation for the given language, then the record is still returned but with empty values for the translated fields.

	*/
	AcceptLanguage *string
	/*ID
	  ID of the series

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get series ID images params
func (o *GetSeriesIDImagesParams) WithTimeout(timeout time.Duration) *GetSeriesIDImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get series ID images params
func (o *GetSeriesIDImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get series ID images params
func (o *GetSeriesIDImagesParams) WithContext(ctx context.Context) *GetSeriesIDImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get series ID images params
func (o *GetSeriesIDImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get series ID images params
func (o *GetSeriesIDImagesParams) WithHTTPClient(client *http.Client) *GetSeriesIDImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get series ID images params
func (o *GetSeriesIDImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get series ID images params
func (o *GetSeriesIDImagesParams) WithAcceptLanguage(acceptLanguage *string) *GetSeriesIDImagesParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get series ID images params
func (o *GetSeriesIDImagesParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithID adds the id to the get series ID images params
func (o *GetSeriesIDImagesParams) WithID(id int64) *GetSeriesIDImagesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get series ID images params
func (o *GetSeriesIDImagesParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSeriesIDImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
