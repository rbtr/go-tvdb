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

// GetSeriesIDFilterReader is a Reader for the GetSeriesIDFilter structure.
type GetSeriesIDFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSeriesIDFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSeriesIDFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSeriesIDFilterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSeriesIDFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSeriesIDFilterOK creates a GetSeriesIDFilterOK with default headers values
func NewGetSeriesIDFilterOK() *GetSeriesIDFilterOK {
	return &GetSeriesIDFilterOK{}
}

/*GetSeriesIDFilterOK handles this case with default header values.

A filtered series record
*/
type GetSeriesIDFilterOK struct {
	Payload *models.SeriesData
}

func (o *GetSeriesIDFilterOK) Error() string {
	return fmt.Sprintf("[GET /series/{id}/filter][%d] getSeriesIdFilterOK  %+v", 200, o.Payload)
}

func (o *GetSeriesIDFilterOK) GetPayload() *models.SeriesData {
	return o.Payload
}

func (o *GetSeriesIDFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeriesData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesIDFilterUnauthorized creates a GetSeriesIDFilterUnauthorized with default headers values
func NewGetSeriesIDFilterUnauthorized() *GetSeriesIDFilterUnauthorized {
	return &GetSeriesIDFilterUnauthorized{}
}

/*GetSeriesIDFilterUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetSeriesIDFilterUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetSeriesIDFilterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /series/{id}/filter][%d] getSeriesIdFilterUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSeriesIDFilterUnauthorized) GetPayload() *models.NotAuthorized {
	return o.Payload
}

func (o *GetSeriesIDFilterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesIDFilterNotFound creates a GetSeriesIDFilterNotFound with default headers values
func NewGetSeriesIDFilterNotFound() *GetSeriesIDFilterNotFound {
	return &GetSeriesIDFilterNotFound{}
}

/*GetSeriesIDFilterNotFound handles this case with default header values.

Returned if the given series ID does not exist
*/
type GetSeriesIDFilterNotFound struct {
	Payload *models.NotFound
}

func (o *GetSeriesIDFilterNotFound) Error() string {
	return fmt.Sprintf("[GET /series/{id}/filter][%d] getSeriesIdFilterNotFound  %+v", 404, o.Payload)
}

func (o *GetSeriesIDFilterNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetSeriesIDFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
