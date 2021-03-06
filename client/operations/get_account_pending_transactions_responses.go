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

// GetAccountPendingTransactionsReader is a Reader for the GetAccountPendingTransactions structure.
type GetAccountPendingTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountPendingTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountPendingTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAccountPendingTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAccountPendingTransactionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountPendingTransactionsOK creates a GetAccountPendingTransactionsOK with default headers values
func NewGetAccountPendingTransactionsOK() *GetAccountPendingTransactionsOK {
	return &GetAccountPendingTransactionsOK{}
}

/*GetAccountPendingTransactionsOK handles this case with default header values.

successful operation
*/
type GetAccountPendingTransactionsOK struct {
	Payload models.Transactions
}

func (o *GetAccountPendingTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /account/{account_pubkey}/pending_transactions][%d] getAccountPendingTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetAccountPendingTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountPendingTransactionsBadRequest creates a GetAccountPendingTransactionsBadRequest with default headers values
func NewGetAccountPendingTransactionsBadRequest() *GetAccountPendingTransactionsBadRequest {
	return &GetAccountPendingTransactionsBadRequest{}
}

/*GetAccountPendingTransactionsBadRequest handles this case with default header values.

Invalid account_pubkey provided
*/
type GetAccountPendingTransactionsBadRequest struct {
	Payload *models.Error
}

func (o *GetAccountPendingTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /account/{account_pubkey}/pending_transactions][%d] getAccountPendingTransactionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountPendingTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountPendingTransactionsNotFound creates a GetAccountPendingTransactionsNotFound with default headers values
func NewGetAccountPendingTransactionsNotFound() *GetAccountPendingTransactionsNotFound {
	return &GetAccountPendingTransactionsNotFound{}
}

/*GetAccountPendingTransactionsNotFound handles this case with default header values.

Account not found
*/
type GetAccountPendingTransactionsNotFound struct {
	Payload *models.Error
}

func (o *GetAccountPendingTransactionsNotFound) Error() string {
	return fmt.Sprintf("[GET /account/{account_pubkey}/pending_transactions][%d] getAccountPendingTransactionsNotFound  %+v", 404, o.Payload)
}

func (o *GetAccountPendingTransactionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
