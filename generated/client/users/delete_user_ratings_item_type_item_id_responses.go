// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rbtr/go-tvdb/generated/models"
)

// DeleteUserRatingsItemTypeItemIDReader is a Reader for the DeleteUserRatingsItemTypeItemID structure.
type DeleteUserRatingsItemTypeItemIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserRatingsItemTypeItemIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserRatingsItemTypeItemIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUserRatingsItemTypeItemIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserRatingsItemTypeItemIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserRatingsItemTypeItemIDOK creates a DeleteUserRatingsItemTypeItemIDOK with default headers values
func NewDeleteUserRatingsItemTypeItemIDOK() *DeleteUserRatingsItemTypeItemIDOK {
	return &DeleteUserRatingsItemTypeItemIDOK{}
}

/*DeleteUserRatingsItemTypeItemIDOK handles this case with default header values.

Returns OK if the delete was successful
*/
type DeleteUserRatingsItemTypeItemIDOK struct {
	Payload *models.UserRatingsDataNoLinksEmptyArray
}

func (o *DeleteUserRatingsItemTypeItemIDOK) Error() string {
	return fmt.Sprintf("[DELETE /user/ratings/{itemType}/{itemId}][%d] deleteUserRatingsItemTypeItemIdOK  %+v", 200, o.Payload)
}

func (o *DeleteUserRatingsItemTypeItemIDOK) GetPayload() *models.UserRatingsDataNoLinksEmptyArray {
	return o.Payload
}

func (o *DeleteUserRatingsItemTypeItemIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRatingsDataNoLinksEmptyArray)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRatingsItemTypeItemIDUnauthorized creates a DeleteUserRatingsItemTypeItemIDUnauthorized with default headers values
func NewDeleteUserRatingsItemTypeItemIDUnauthorized() *DeleteUserRatingsItemTypeItemIDUnauthorized {
	return &DeleteUserRatingsItemTypeItemIDUnauthorized{}
}

/*DeleteUserRatingsItemTypeItemIDUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type DeleteUserRatingsItemTypeItemIDUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *DeleteUserRatingsItemTypeItemIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /user/ratings/{itemType}/{itemId}][%d] deleteUserRatingsItemTypeItemIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserRatingsItemTypeItemIDUnauthorized) GetPayload() *models.NotAuthorized {
	return o.Payload
}

func (o *DeleteUserRatingsItemTypeItemIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRatingsItemTypeItemIDNotFound creates a DeleteUserRatingsItemTypeItemIDNotFound with default headers values
func NewDeleteUserRatingsItemTypeItemIDNotFound() *DeleteUserRatingsItemTypeItemIDNotFound {
	return &DeleteUserRatingsItemTypeItemIDNotFound{}
}

/*DeleteUserRatingsItemTypeItemIDNotFound handles this case with default header values.

Returned if no rating is found that matches your given parameters
*/
type DeleteUserRatingsItemTypeItemIDNotFound struct {
	Payload *models.NotFound
}

func (o *DeleteUserRatingsItemTypeItemIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user/ratings/{itemType}/{itemId}][%d] deleteUserRatingsItemTypeItemIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserRatingsItemTypeItemIDNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *DeleteUserRatingsItemTypeItemIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
