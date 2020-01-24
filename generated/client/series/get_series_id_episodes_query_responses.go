// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rbtr/go-tvdb/generated/models"
)

// GetSeriesIDEpisodesQueryReader is a Reader for the GetSeriesIDEpisodesQuery structure.
type GetSeriesIDEpisodesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSeriesIDEpisodesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSeriesIDEpisodesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSeriesIDEpisodesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSeriesIDEpisodesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSeriesIDEpisodesQueryOK creates a GetSeriesIDEpisodesQueryOK with default headers values
func NewGetSeriesIDEpisodesQueryOK() *GetSeriesIDEpisodesQueryOK {
	return &GetSeriesIDEpisodesQueryOK{}
}

/*GetSeriesIDEpisodesQueryOK handles this case with default header values.

An array of basic Episode results that matched the query
*/
type GetSeriesIDEpisodesQueryOK struct {
	Payload *models.SeriesEpisodesQuery
}

func (o *GetSeriesIDEpisodesQueryOK) Error() string {
	return fmt.Sprintf("[GET /series/{id}/episodes/query][%d] getSeriesIdEpisodesQueryOK  %+v", 200, o.Payload)
}

func (o *GetSeriesIDEpisodesQueryOK) GetPayload() *models.SeriesEpisodesQuery {
	return o.Payload
}

func (o *GetSeriesIDEpisodesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeriesEpisodesQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesIDEpisodesQueryUnauthorized creates a GetSeriesIDEpisodesQueryUnauthorized with default headers values
func NewGetSeriesIDEpisodesQueryUnauthorized() *GetSeriesIDEpisodesQueryUnauthorized {
	return &GetSeriesIDEpisodesQueryUnauthorized{}
}

/*GetSeriesIDEpisodesQueryUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetSeriesIDEpisodesQueryUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetSeriesIDEpisodesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /series/{id}/episodes/query][%d] getSeriesIdEpisodesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSeriesIDEpisodesQueryUnauthorized) GetPayload() *models.NotAuthorized {
	return o.Payload
}

func (o *GetSeriesIDEpisodesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesIDEpisodesQueryNotFound creates a GetSeriesIDEpisodesQueryNotFound with default headers values
func NewGetSeriesIDEpisodesQueryNotFound() *GetSeriesIDEpisodesQueryNotFound {
	return &GetSeriesIDEpisodesQueryNotFound{}
}

/*GetSeriesIDEpisodesQueryNotFound handles this case with default header values.

Returned if the given series ID does not exist or the query returns no results.
*/
type GetSeriesIDEpisodesQueryNotFound struct {
	Payload *models.NotFound
}

func (o *GetSeriesIDEpisodesQueryNotFound) Error() string {
	return fmt.Sprintf("[GET /series/{id}/episodes/query][%d] getSeriesIdEpisodesQueryNotFound  %+v", 404, o.Payload)
}

func (o *GetSeriesIDEpisodesQueryNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetSeriesIDEpisodesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
