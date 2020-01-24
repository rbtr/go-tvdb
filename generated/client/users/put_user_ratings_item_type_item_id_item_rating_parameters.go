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

// NewPutUserRatingsItemTypeItemIDItemRatingParams creates a new PutUserRatingsItemTypeItemIDItemRatingParams object
// with the default values initialized.
func NewPutUserRatingsItemTypeItemIDItemRatingParams() *PutUserRatingsItemTypeItemIDItemRatingParams {
	var ()
	return &PutUserRatingsItemTypeItemIDItemRatingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutUserRatingsItemTypeItemIDItemRatingParamsWithTimeout creates a new PutUserRatingsItemTypeItemIDItemRatingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutUserRatingsItemTypeItemIDItemRatingParamsWithTimeout(timeout time.Duration) *PutUserRatingsItemTypeItemIDItemRatingParams {
	var ()
	return &PutUserRatingsItemTypeItemIDItemRatingParams{

		timeout: timeout,
	}
}

// NewPutUserRatingsItemTypeItemIDItemRatingParamsWithContext creates a new PutUserRatingsItemTypeItemIDItemRatingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutUserRatingsItemTypeItemIDItemRatingParamsWithContext(ctx context.Context) *PutUserRatingsItemTypeItemIDItemRatingParams {
	var ()
	return &PutUserRatingsItemTypeItemIDItemRatingParams{

		Context: ctx,
	}
}

// NewPutUserRatingsItemTypeItemIDItemRatingParamsWithHTTPClient creates a new PutUserRatingsItemTypeItemIDItemRatingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutUserRatingsItemTypeItemIDItemRatingParamsWithHTTPClient(client *http.Client) *PutUserRatingsItemTypeItemIDItemRatingParams {
	var ()
	return &PutUserRatingsItemTypeItemIDItemRatingParams{
		HTTPClient: client,
	}
}

/*PutUserRatingsItemTypeItemIDItemRatingParams contains all the parameters to send to the API endpoint
for the put user ratings item type item ID item rating operation typically these are written to a http.Request
*/
type PutUserRatingsItemTypeItemIDItemRatingParams struct {

	/*ItemID
	  ID of the ratings record that you wish to modify

	*/
	ItemID int64
	/*ItemRating
	  The updated rating number

	*/
	ItemRating int64
	/*ItemType
	  Item to update. Can be either 'series', 'episode', or 'image'

	*/
	ItemType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) WithTimeout(timeout time.Duration) *PutUserRatingsItemTypeItemIDItemRatingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) WithContext(ctx context.Context) *PutUserRatingsItemTypeItemIDItemRatingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) WithHTTPClient(client *http.Client) *PutUserRatingsItemTypeItemIDItemRatingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemID adds the itemID to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) WithItemID(itemID int64) *PutUserRatingsItemTypeItemIDItemRatingParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) SetItemID(itemID int64) {
	o.ItemID = itemID
}

// WithItemRating adds the itemRating to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) WithItemRating(itemRating int64) *PutUserRatingsItemTypeItemIDItemRatingParams {
	o.SetItemRating(itemRating)
	return o
}

// SetItemRating adds the itemRating to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) SetItemRating(itemRating int64) {
	o.ItemRating = itemRating
}

// WithItemType adds the itemType to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) WithItemType(itemType string) *PutUserRatingsItemTypeItemIDItemRatingParams {
	o.SetItemType(itemType)
	return o
}

// SetItemType adds the itemType to the put user ratings item type item ID item rating params
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) SetItemType(itemType string) {
	o.ItemType = itemType
}

// WriteToRequest writes these params to a swagger request
func (o *PutUserRatingsItemTypeItemIDItemRatingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param itemId
	if err := r.SetPathParam("itemId", swag.FormatInt64(o.ItemID)); err != nil {
		return err
	}

	// path param itemRating
	if err := r.SetPathParam("itemRating", swag.FormatInt64(o.ItemRating)); err != nil {
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
