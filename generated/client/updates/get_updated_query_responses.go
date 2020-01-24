// Code generated by go-swagger; DO NOT EDIT.

package updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rbtr/go-tvdb/generated/models"
)

// GetUpdatedQueryReader is a Reader for the GetUpdatedQuery structure.
type GetUpdatedQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUpdatedQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUpdatedQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUpdatedQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUpdatedQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUpdatedQueryOK creates a GetUpdatedQueryOK with default headers values
func NewGetUpdatedQueryOK() *GetUpdatedQueryOK {
	return &GetUpdatedQueryOK{}
}

/*GetUpdatedQueryOK handles this case with default header values.

An array of Update objects that match the given timeframe.
*/
type GetUpdatedQueryOK struct {
	Payload *models.UpdateData
}

func (o *GetUpdatedQueryOK) Error() string {
	return fmt.Sprintf("[GET /updated/query][%d] getUpdatedQueryOK  %+v", 200, o.Payload)
}

func (o *GetUpdatedQueryOK) GetPayload() *models.UpdateData {
	return o.Payload
}

func (o *GetUpdatedQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpdatedQueryUnauthorized creates a GetUpdatedQueryUnauthorized with default headers values
func NewGetUpdatedQueryUnauthorized() *GetUpdatedQueryUnauthorized {
	return &GetUpdatedQueryUnauthorized{}
}

/*GetUpdatedQueryUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetUpdatedQueryUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetUpdatedQueryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /updated/query][%d] getUpdatedQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUpdatedQueryUnauthorized) GetPayload() *models.NotAuthorized {
	return o.Payload
}

func (o *GetUpdatedQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpdatedQueryNotFound creates a GetUpdatedQueryNotFound with default headers values
func NewGetUpdatedQueryNotFound() *GetUpdatedQueryNotFound {
	return &GetUpdatedQueryNotFound{}
}

/*GetUpdatedQueryNotFound handles this case with default header values.

Returned if no records exist for the given timespan
*/
type GetUpdatedQueryNotFound struct {
	Payload *models.NotFound
}

func (o *GetUpdatedQueryNotFound) Error() string {
	return fmt.Sprintf("[GET /updated/query][%d] getUpdatedQueryNotFound  %+v", 404, o.Payload)
}

func (o *GetUpdatedQueryNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetUpdatedQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
