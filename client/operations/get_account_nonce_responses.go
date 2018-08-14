// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/models"
)

// GetAccountNonceReader is a Reader for the GetAccountNonce structure.
type GetAccountNonceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountNonceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountNonceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAccountNonceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAccountNonceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountNonceOK creates a GetAccountNonceOK with default headers values
func NewGetAccountNonceOK() *GetAccountNonceOK {
	return &GetAccountNonceOK{}
}

/*GetAccountNonceOK handles this case with default header values.

successful operation
*/
type GetAccountNonceOK struct {
	Payload *models.Nonce
}

func (o *GetAccountNonceOK) Error() string {
	return fmt.Sprintf("[GET /account/{account_pubkey}/nonce][%d] getAccountNonceOK  %+v", 200, o.Payload)
}

func (o *GetAccountNonceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Nonce)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountNonceBadRequest creates a GetAccountNonceBadRequest with default headers values
func NewGetAccountNonceBadRequest() *GetAccountNonceBadRequest {
	return &GetAccountNonceBadRequest{}
}

/*GetAccountNonceBadRequest handles this case with default header values.

Invalid account_pubkey provided
*/
type GetAccountNonceBadRequest struct {
	Payload *models.Error
}

func (o *GetAccountNonceBadRequest) Error() string {
	return fmt.Sprintf("[GET /account/{account_pubkey}/nonce][%d] getAccountNonceBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountNonceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountNonceNotFound creates a GetAccountNonceNotFound with default headers values
func NewGetAccountNonceNotFound() *GetAccountNonceNotFound {
	return &GetAccountNonceNotFound{}
}

/*GetAccountNonceNotFound handles this case with default header values.

Account not found
*/
type GetAccountNonceNotFound struct {
	Payload *models.Error
}

func (o *GetAccountNonceNotFound) Error() string {
	return fmt.Sprintf("[GET /account/{account_pubkey}/nonce][%d] getAccountNonceNotFound  %+v", 404, o.Payload)
}

func (o *GetAccountNonceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}