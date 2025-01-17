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

// NewGetSeriesIDImagesQueryParamsParams creates a new GetSeriesIDImagesQueryParamsParams object
// with the default values initialized.
func NewGetSeriesIDImagesQueryParamsParams() *GetSeriesIDImagesQueryParamsParams {
	var ()
	return &GetSeriesIDImagesQueryParamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSeriesIDImagesQueryParamsParamsWithTimeout creates a new GetSeriesIDImagesQueryParamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSeriesIDImagesQueryParamsParamsWithTimeout(timeout time.Duration) *GetSeriesIDImagesQueryParamsParams {
	var ()
	return &GetSeriesIDImagesQueryParamsParams{

		timeout: timeout,
	}
}

// NewGetSeriesIDImagesQueryParamsParamsWithContext creates a new GetSeriesIDImagesQueryParamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSeriesIDImagesQueryParamsParamsWithContext(ctx context.Context) *GetSeriesIDImagesQueryParamsParams {
	var ()
	return &GetSeriesIDImagesQueryParamsParams{

		Context: ctx,
	}
}

// NewGetSeriesIDImagesQueryParamsParamsWithHTTPClient creates a new GetSeriesIDImagesQueryParamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSeriesIDImagesQueryParamsParamsWithHTTPClient(client *http.Client) *GetSeriesIDImagesQueryParamsParams {
	var ()
	return &GetSeriesIDImagesQueryParamsParams{
		HTTPClient: client,
	}
}

/*GetSeriesIDImagesQueryParamsParams contains all the parameters to send to the API endpoint
for the get series ID images query params operation typically these are written to a http.Request
*/
type GetSeriesIDImagesQueryParamsParams struct {

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

// WithTimeout adds the timeout to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) WithTimeout(timeout time.Duration) *GetSeriesIDImagesQueryParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) WithContext(ctx context.Context) *GetSeriesIDImagesQueryParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) WithHTTPClient(client *http.Client) *GetSeriesIDImagesQueryParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) WithAcceptLanguage(acceptLanguage *string) *GetSeriesIDImagesQueryParamsParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithID adds the id to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) WithID(id int64) *GetSeriesIDImagesQueryParamsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get series ID images query params params
func (o *GetSeriesIDImagesQueryParamsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSeriesIDImagesQueryParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
