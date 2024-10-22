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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteUserRatingsItemTypeItemIDParams creates a new DeleteUserRatingsItemTypeItemIDParams object
// with the default values initialized.
func NewDeleteUserRatingsItemTypeItemIDParams() *DeleteUserRatingsItemTypeItemIDParams {
	var ()
	return &DeleteUserRatingsItemTypeItemIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserRatingsItemTypeItemIDParamsWithTimeout creates a new DeleteUserRatingsItemTypeItemIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserRatingsItemTypeItemIDParamsWithTimeout(timeout time.Duration) *DeleteUserRatingsItemTypeItemIDParams {
	var ()
	return &DeleteUserRatingsItemTypeItemIDParams{

		timeout: timeout,
	}
}

// NewDeleteUserRatingsItemTypeItemIDParamsWithContext creates a new DeleteUserRatingsItemTypeItemIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserRatingsItemTypeItemIDParamsWithContext(ctx context.Context) *DeleteUserRatingsItemTypeItemIDParams {
	var ()
	return &DeleteUserRatingsItemTypeItemIDParams{

		Context: ctx,
	}
}

// NewDeleteUserRatingsItemTypeItemIDParamsWithHTTPClient creates a new DeleteUserRatingsItemTypeItemIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserRatingsItemTypeItemIDParamsWithHTTPClient(client *http.Client) *DeleteUserRatingsItemTypeItemIDParams {
	var ()
	return &DeleteUserRatingsItemTypeItemIDParams{
		HTTPClient: client,
	}
}

/*DeleteUserRatingsItemTypeItemIDParams contains all the parameters to send to the API endpoint
for the delete user ratings item type item ID operation typically these are written to a http.Request
*/
type DeleteUserRatingsItemTypeItemIDParams struct {

	/*ItemID
	  ID of the ratings record that you wish to modify

	*/
	ItemID int64
	/*ItemType
	  Item to update. Can be either 'series', 'episode', or 'image'

	*/
	ItemType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) WithTimeout(timeout time.Duration) *DeleteUserRatingsItemTypeItemIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) WithContext(ctx context.Context) *DeleteUserRatingsItemTypeItemIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) WithHTTPClient(client *http.Client) *DeleteUserRatingsItemTypeItemIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemID adds the itemID to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) WithItemID(itemID int64) *DeleteUserRatingsItemTypeItemIDParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) SetItemID(itemID int64) {
	o.ItemID = itemID
}

// WithItemType adds the itemType to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) WithItemType(itemType string) *DeleteUserRatingsItemTypeItemIDParams {
	o.SetItemType(itemType)
	return o
}

// SetItemType adds the itemType to the delete user ratings item type item ID params
func (o *DeleteUserRatingsItemTypeItemIDParams) SetItemType(itemType string) {
	o.ItemType = itemType
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserRatingsItemTypeItemIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param itemId
	if err := r.SetPathParam("itemId", swag.FormatInt64(o.ItemID)); err != nil {
		return err
	}

	// path param itemType
	if err := r.SetPathParam("itemType", o.ItemType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
