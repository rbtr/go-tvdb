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

// GetMovieupdatesReader is a Reader for the GetMovieupdates structure.
type GetMovieupdatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMovieupdatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMovieupdatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMovieupdatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetMovieupdatesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetMovieupdatesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMovieupdatesOK creates a GetMovieupdatesOK with default headers values
func NewGetMovieupdatesOK() *GetMovieupdatesOK {
	return &GetMovieupdatesOK{}
}

/*GetMovieupdatesOK handles this case with default header values.

All movies ids updated since a given date
*/
type GetMovieupdatesOK struct {
	Payload *models.UpdatedMovies
}

func (o *GetMovieupdatesOK) Error() string {
	return fmt.Sprintf("[GET /movieupdates][%d] getMovieupdatesOK  %+v", 200, o.Payload)
}

func (o *GetMovieupdatesOK) GetPayload() *models.UpdatedMovies {
	return o.Payload
}

func (o *GetMovieupdatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdatedMovies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMovieupdatesUnauthorized creates a GetMovieupdatesUnauthorized with default headers values
func NewGetMovieupdatesUnauthorized() *GetMovieupdatesUnauthorized {
	return &GetMovieupdatesUnauthorized{}
}

/*GetMovieupdatesUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetMovieupdatesUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetMovieupdatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /movieupdates][%d] getMovieupdatesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMovieupdatesUnauthorized) GetPayload() *models.NotAuthorized {
	return o.Payload
}

func (o *GetMovieupdatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMovieupdatesMethodNotAllowed creates a GetMovieupdatesMethodNotAllowed with default headers values
func NewGetMovieupdatesMethodNotAllowed() *GetMovieupdatesMethodNotAllowed {
	return &GetMovieupdatesMethodNotAllowed{}
}

/*GetMovieupdatesMethodNotAllowed handles this case with default header values.

Missing query params are given
*/
type GetMovieupdatesMethodNotAllowed struct {
	Payload *models.InvalidQueryParams
}

func (o *GetMovieupdatesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /movieupdates][%d] getMovieupdatesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetMovieupdatesMethodNotAllowed) GetPayload() *models.InvalidQueryParams {
	return o.Payload
}

func (o *GetMovieupdatesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvalidQueryParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMovieupdatesUnprocessableEntity creates a GetMovieupdatesUnprocessableEntity with default headers values
func NewGetMovieupdatesUnprocessableEntity() *GetMovieupdatesUnprocessableEntity {
	return &GetMovieupdatesUnprocessableEntity{}
}

/*GetMovieupdatesUnprocessableEntity handles this case with default header values.

Invalid query params provided
*/
type GetMovieupdatesUnprocessableEntity struct {
	Payload *models.InvalidQueryParams
}

func (o *GetMovieupdatesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /movieupdates][%d] getMovieupdatesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetMovieupdatesUnprocessableEntity) GetPayload() *models.InvalidQueryParams {
	return o.Payload
}

func (o *GetMovieupdatesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvalidQueryParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
