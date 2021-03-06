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

// GetAccountBalanceReader is a Reader for the GetAccountBalance structure.
type GetAccountBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAccountBalanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAccountBalanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountBalanceOK creates a GetAccountBalanceOK with default headers values
func NewGetAccountBalanceOK() *GetAccountBalanceOK {
	return &GetAccountBalanceOK{}
}

/*GetAccountBalanceOK handles this case with default header values.

successful operation
*/
type GetAccountBalanceOK struct {
	Payload *models.Balance
}

func (o *GetAccountBalanceOK) Error() string {
	return fmt.Sprintf("[GET /account/{address}/balance][%d] getAccountBalanceOK  %+v", 200, o.Payload)
}

func (o *GetAccountBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Balance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountBalanceBadRequest creates a GetAccountBalanceBadRequest with default headers values
func NewGetAccountBalanceBadRequest() *GetAccountBalanceBadRequest {
	return &GetAccountBalanceBadRequest{}
}

/*GetAccountBalanceBadRequest handles this case with default header values.

Invalid block hash or hash and height combination
*/
type GetAccountBalanceBadRequest struct {
	Payload *models.Error
}

func (o *GetAccountBalanceBadRequest) Error() string {
	return fmt.Sprintf("[GET /account/{address}/balance][%d] getAccountBalanceBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountBalanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountBalanceNotFound creates a GetAccountBalanceNotFound with default headers values
func NewGetAccountBalanceNotFound() *GetAccountBalanceNotFound {
	return &GetAccountBalanceNotFound{}
}

/*GetAccountBalanceNotFound handles this case with default header values.

Block or account not found
*/
type GetAccountBalanceNotFound struct {
	Payload *models.Error
}

func (o *GetAccountBalanceNotFound) Error() string {
	return fmt.Sprintf("[GET /account/{address}/balance][%d] getAccountBalanceNotFound  %+v", 404, o.Payload)
}

func (o *GetAccountBalanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
