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

// NewGetSeriesIDEpisodesQueryParams creates a new GetSeriesIDEpisodesQueryParams object
// with the default values initialized.
func NewGetSeriesIDEpisodesQueryParams() *GetSeriesIDEpisodesQueryParams {
	var ()
	return &GetSeriesIDEpisodesQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSeriesIDEpisodesQueryParamsWithTimeout creates a new GetSeriesIDEpisodesQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSeriesIDEpisodesQueryParamsWithTimeout(timeout time.Duration) *GetSeriesIDEpisodesQueryParams {
	var ()
	return &GetSeriesIDEpisodesQueryParams{

		timeout: timeout,
	}
}

// NewGetSeriesIDEpisodesQueryParamsWithContext creates a new GetSeriesIDEpisodesQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSeriesIDEpisodesQueryParamsWithContext(ctx context.Context) *GetSeriesIDEpisodesQueryParams {
	var ()
	return &GetSeriesIDEpisodesQueryParams{

		Context: ctx,
	}
}

// NewGetSeriesIDEpisodesQueryParamsWithHTTPClient creates a new GetSeriesIDEpisodesQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSeriesIDEpisodesQueryParamsWithHTTPClient(client *http.Client) *GetSeriesIDEpisodesQueryParams {
	var ()
	return &GetSeriesIDEpisodesQueryParams{
		HTTPClient: client,
	}
}

/*GetSeriesIDEpisodesQueryParams contains all the parameters to send to the API endpoint
for the get series ID episodes query operation typically these are written to a http.Request
*/
type GetSeriesIDEpisodesQueryParams struct {

	/*AcceptLanguage
	  Records are returned with the some fields in the desired language, if it exists. If there is no translation for the given language, then the record is still returned but with empty values for the translated fields.

	*/
	AcceptLanguage *string
	/*AbsoluteNumber
	  Absolute number of the episode

	*/
	AbsoluteNumber *string
	/*AiredEpisode
	  Aired episode number

	*/
	AiredEpisode *string
	/*AiredSeason
	  Aired season number

	*/
	AiredSeason *string
	/*DvdEpisode
	  DVD episode number

	*/
	DvdEpisode *string
	/*DvdSeason
	  DVD season number

	*/
	DvdSeason *string
	/*ID
	  ID of the series

	*/
	ID int64
	/*ImdbID
	  IMDB id of the series

	*/
	ImdbID *string
	/*Page
	  Page of results to fetch. Defaults to page 1 if not provided.

	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithTimeout(timeout time.Duration) *GetSeriesIDEpisodesQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithContext(ctx context.Context) *GetSeriesIDEpisodesQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithHTTPClient(client *http.Client) *GetSeriesIDEpisodesQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithAcceptLanguage(acceptLanguage *string) *GetSeriesIDEpisodesQueryParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithAbsoluteNumber adds the absoluteNumber to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithAbsoluteNumber(absoluteNumber *string) *GetSeriesIDEpisodesQueryParams {
	o.SetAbsoluteNumber(absoluteNumber)
	return o
}

// SetAbsoluteNumber adds the absoluteNumber to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetAbsoluteNumber(absoluteNumber *string) {
	o.AbsoluteNumber = absoluteNumber
}

// WithAiredEpisode adds the airedEpisode to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithAiredEpisode(airedEpisode *string) *GetSeriesIDEpisodesQueryParams {
	o.SetAiredEpisode(airedEpisode)
	return o
}

// SetAiredEpisode adds the airedEpisode to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetAiredEpisode(airedEpisode *string) {
	o.AiredEpisode = airedEpisode
}

// WithAiredSeason adds the airedSeason to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithAiredSeason(airedSeason *string) *GetSeriesIDEpisodesQueryParams {
	o.SetAiredSeason(airedSeason)
	return o
}

// SetAiredSeason adds the airedSeason to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetAiredSeason(airedSeason *string) {
	o.AiredSeason = airedSeason
}

// WithDvdEpisode adds the dvdEpisode to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithDvdEpisode(dvdEpisode *string) *GetSeriesIDEpisodesQueryParams {
	o.SetDvdEpisode(dvdEpisode)
	return o
}

// SetDvdEpisode adds the dvdEpisode to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetDvdEpisode(dvdEpisode *string) {
	o.DvdEpisode = dvdEpisode
}

// WithDvdSeason adds the dvdSeason to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithDvdSeason(dvdSeason *string) *GetSeriesIDEpisodesQueryParams {
	o.SetDvdSeason(dvdSeason)
	return o
}

// SetDvdSeason adds the dvdSeason to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetDvdSeason(dvdSeason *string) {
	o.DvdSeason = dvdSeason
}

// WithID adds the id to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithID(id int64) *GetSeriesIDEpisodesQueryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetID(id int64) {
	o.ID = id
}

// WithImdbID adds the imdbID to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithImdbID(imdbID *string) *GetSeriesIDEpisodesQueryParams {
	o.SetImdbID(imdbID)
	return o
}

// SetImdbID adds the imdbId to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetImdbID(imdbID *string) {
	o.ImdbID = imdbID
}

// WithPage adds the page to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) WithPage(page *string) *GetSeriesIDEpisodesQueryParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get series ID episodes query params
func (o *GetSeriesIDEpisodesQueryParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetSeriesIDEpisodesQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AbsoluteNumber != nil {

		// query param absoluteNumber
		var qrAbsoluteNumber string
		if o.AbsoluteNumber != nil {
			qrAbsoluteNumber = *o.AbsoluteNumber
		}
		qAbsoluteNumber := qrAbsoluteNumber
		if qAbsoluteNumber != "" {
			if err := r.SetQueryParam("absoluteNumber", qAbsoluteNumber); err != nil {
				return err
			}
		}

	}

	if o.AiredEpisode != nil {

		// query param airedEpisode
		var qrAiredEpisode string
		if o.AiredEpisode != nil {
			qrAiredEpisode = *o.AiredEpisode
		}
		qAiredEpisode := qrAiredEpisode
		if qAiredEpisode != "" {
			if err := r.SetQueryParam("airedEpisode", qAiredEpisode); err != nil {
				return err
			}
		}

	}

	if o.AiredSeason != nil {

		// query param airedSeason
		var qrAiredSeason string
		if o.AiredSeason != nil {
			qrAiredSeason = *o.AiredSeason
		}
		qAiredSeason := qrAiredSeason
		if qAiredSeason != "" {
			if err := r.SetQueryParam("airedSeason", qAiredSeason); err != nil {
				return err
			}
		}

	}

	if o.DvdEpisode != nil {

		// query param dvdEpisode
		var qrDvdEpisode string
		if o.DvdEpisode != nil {
			qrDvdEpisode = *o.DvdEpisode
		}
		qDvdEpisode := qrDvdEpisode
		if qDvdEpisode != "" {
			if err := r.SetQueryParam("dvdEpisode", qDvdEpisode); err != nil {
				return err
			}
		}

	}

	if o.DvdSeason != nil {

		// query param dvdSeason
		var qrDvdSeason string
		if o.DvdSeason != nil {
			qrDvdSeason = *o.DvdSeason
		}
		qDvdSeason := qrDvdSeason
		if qDvdSeason != "" {
			if err := r.SetQueryParam("dvdSeason", qDvdSeason); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.ImdbID != nil {

		// query param imdbId
		var qrImdbID string
		if o.ImdbID != nil {
			qrImdbID = *o.ImdbID
		}
		qImdbID := qrImdbID
		if qImdbID != "" {
			if err := r.SetQueryParam("imdbId", qImdbID); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
