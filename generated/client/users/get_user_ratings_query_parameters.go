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

// NewGetUserRatingsQueryParams creates a new GetUserRatingsQueryParams object
// with the default values initialized.
func NewGetUserRatingsQueryParams() *GetUserRatingsQueryParams {
	var ()
	return &GetUserRatingsQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRatingsQueryParamsWithTimeout creates a new GetUserRatingsQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserRatingsQueryParamsWithTimeout(timeout time.Duration) *GetUserRatingsQueryParams {
	var ()
	return &GetUserRatingsQueryParams{

		timeout: timeout,
	}
}

// NewGetUserRatingsQueryParamsWithContext creates a new GetUserRatingsQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserRatingsQueryParamsWithContext(ctx context.Context) *GetUserRatingsQueryParams {
	var ()
	return &GetUserRatingsQueryParams{

		Context: ctx,
	}
}

// NewGetUserRatingsQueryParamsWithHTTPClient creates a new GetUserRatingsQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserRatingsQueryParamsWithHTTPClient(client *http.Client) *GetUserRatingsQueryParams {
	var ()
	return &GetUserRatingsQueryParams{
		HTTPClient: client,
	}
}

/*GetUserRatingsQueryParams contains all the parameters to send to the API endpoint
for the get user ratings query operation typically these are written to a http.Request
*/
type GetUserRatingsQueryParams struct {

	/*ItemType
	  Item to query. Can be either 'series', 'episode', or 'banner'

	*/
	ItemType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user ratings query params
func (o *GetUserRatingsQueryParams) WithTimeout(timeout time.Duration) *GetUserRatingsQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user ratings query params
func (o *GetUserRatingsQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user ratings query params
func (o *GetUserRatingsQueryParams) WithContext(ctx context.Context) *GetUserRatingsQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user ratings query params
func (o *GetUserRatingsQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user ratings query params
func (o *GetUserRatingsQueryParams) WithHTTPClient(client *http.Client) *GetUserRatingsQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user ratings query params
func (o *GetUserRatingsQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemType adds the itemType to the get user ratings query params
func (o *GetUserRatingsQueryParams) WithItemType(itemType *string) *GetUserRatingsQueryParams {
	o.SetItemType(itemType)
	return o
}

// SetItemType adds the itemType to the get user ratings query params
func (o *GetUserRatingsQueryParams) SetItemType(itemType *string) {
	o.ItemType = itemType
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRatingsQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ItemType != nil {

		// query param itemType
		var qrItemType string
		if o.ItemType != nil {
			qrItemType = *o.ItemType
		}
		qItemType := qrItemType
		if qItemType != "" {
			if err := r.SetQueryParam("itemType", qItemType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
