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

// HeadSeriesIDReader is a Reader for the HeadSeriesID structure.
type HeadSeriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadSeriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadSeriesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHeadSeriesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHeadSeriesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHeadSeriesIDOK creates a HeadSeriesIDOK with default headers values
func NewHeadSeriesIDOK() *HeadSeriesIDOK {
	return &HeadSeriesIDOK{}
}

/*HeadSeriesIDOK handles this case with default header values.

Series header information. Good for getting the Last-Updated header to find out when the series was last modified.
*/
type HeadSeriesIDOK struct {
}

func (o *HeadSeriesIDOK) Error() string {
	return fmt.Sprintf("[HEAD /series/{id}][%d] headSeriesIdOK ", 200)
}

func (o *HeadSeriesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeadSeriesIDUnauthorized creates a HeadSeriesIDUnauthorized with default headers values
func NewHeadSeriesIDUnauthorized() *HeadSeriesIDUnauthorized {
	return &HeadSeriesIDUnauthorized{}
}

/*HeadSeriesIDUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type HeadSeriesIDUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *HeadSeriesIDUnauthorized) Error() string {
	return fmt.Sprintf("[HEAD /series/{id}][%d] headSeriesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *HeadSeriesIDUnauthorized) GetPayload() *models.NotAuthorized {
	return o.Payload
}

func (o *HeadSeriesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadSeriesIDNotFound creates a HeadSeriesIDNotFound with default headers values
func NewHeadSeriesIDNotFound() *HeadSeriesIDNotFound {
	return &HeadSeriesIDNotFound{}
}

/*HeadSeriesIDNotFound handles this case with default header values.

Returned if the given series ID does not exist
*/
type HeadSeriesIDNotFound struct {
	Payload *models.NotFound
}

func (o *HeadSeriesIDNotFound) Error() string {
	return fmt.Sprintf("[HEAD /series/{id}][%d] headSeriesIdNotFound  %+v", 404, o.Payload)
}

func (o *HeadSeriesIDNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *HeadSeriesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
