// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rbtr/go-tvdb/generated/models"
)

// GetRefreshTokenReader is a Reader for the GetRefreshToken structure.
type GetRefreshTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRefreshTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRefreshTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRefreshTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRefreshTokenOK creates a GetRefreshTokenOK with default headers values
func NewGetRefreshTokenOK() *GetRefreshTokenOK {
	return &GetRefreshTokenOK{}
}

/*GetRefreshTokenOK handles this case with default header values.

Returns a new token to use in your subsequent requests
*/
type GetRefreshTokenOK struct {
	Payload *models.Token
}

func (o *GetRefreshTokenOK) Error() string {
	return fmt.Sprintf("[GET /refresh_token][%d] getRefreshTokenOK  %+v", 200, o.Payload)
}

func (o *GetRefreshTokenOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *GetRefreshTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRefreshTokenUnauthorized creates a GetRefreshTokenUnauthorized with default headers values
func NewGetRefreshTokenUnauthorized() *GetRefreshTokenUnauthorized {
	return &GetRefreshTokenUnauthorized{}
}

/*GetRefreshTokenUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetRefreshTokenUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetRefreshTokenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /refresh_token][%d] getRefreshTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRefreshTokenUnauthorized) GetPayload() *models.NotAuthorized {
	return o.Payload
}

func (o *GetRefreshTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
