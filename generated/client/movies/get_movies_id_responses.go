// Code generated by go-swagger; DO NOT EDIT.

package movies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rbtr/go-tvdb/generated/models"
)

// GetMoviesIDReader is a Reader for the GetMoviesID structure.
type GetMoviesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMoviesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMoviesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMoviesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMoviesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMoviesIDOK creates a GetMoviesIDOK with default headers values
func NewGetMoviesIDOK() *GetMoviesIDOK {
	return &GetMoviesIDOK{}
}

/*GetMoviesIDOK handles this case with default header values.

A movie record.
*/
type GetMoviesIDOK struct {
	Payload *models.Movie
}

func (o *GetMoviesIDOK) Error() string {
	return fmt.Sprintf("[GET /movies/{id}][%d] getMoviesIdOK  %+v", 200, o.Payload)
}

func (o *GetMoviesIDOK) GetPayload() *models.Movie {
	return o.Payload
}

func (o *GetMoviesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Movie)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMoviesIDUnauthorized creates a GetMoviesIDUnauthorized with default headers values
func NewGetMoviesIDUnauthorized() *GetMoviesIDUnauthorized {
	return &GetMoviesIDUnauthorized{}
}

/*GetMoviesIDUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetMoviesIDUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetMoviesIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /movies/{id}][%d] getMoviesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMoviesIDUnauthorized) GetPayload() *models.NotAuthorized {
	return o.Payload
}

func (o *GetMoviesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMoviesIDNotFound creates a GetMoviesIDNotFound with default headers values
func NewGetMoviesIDNotFound() *GetMoviesIDNotFound {
	return &GetMoviesIDNotFound{}
}

/*GetMoviesIDNotFound handles this case with default header values.

Returned if the given series ID does not exist
*/
type GetMoviesIDNotFound struct {
	Payload *models.NotFound
}

func (o *GetMoviesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /movies/{id}][%d] getMoviesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetMoviesIDNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetMoviesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
