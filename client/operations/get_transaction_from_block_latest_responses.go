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

// GetTransactionFromBlockLatestReader is a Reader for the GetTransactionFromBlockLatest structure.
type GetTransactionFromBlockLatestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionFromBlockLatestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionFromBlockLatestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetTransactionFromBlockLatestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionFromBlockLatestOK creates a GetTransactionFromBlockLatestOK with default headers values
func NewGetTransactionFromBlockLatestOK() *GetTransactionFromBlockLatestOK {
	return &GetTransactionFromBlockLatestOK{}
}

/*GetTransactionFromBlockLatestOK handles this case with default header values.

The transaction found
*/
type GetTransactionFromBlockLatestOK struct {
	Payload models.SingleTxObject
}

func (o *GetTransactionFromBlockLatestOK) Error() string {
	return fmt.Sprintf("[GET /block/tx/latest/{tx_index}][%d] getTransactionFromBlockLatestOK  %+v", 200, o.Payload)
}

func (o *GetTransactionFromBlockLatestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalSingleTxObject(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetTransactionFromBlockLatestNotFound creates a GetTransactionFromBlockLatestNotFound with default headers values
func NewGetTransactionFromBlockLatestNotFound() *GetTransactionFromBlockLatestNotFound {
	return &GetTransactionFromBlockLatestNotFound{}
}

/*GetTransactionFromBlockLatestNotFound handles this case with default header values.

Block or transaction not found
*/
type GetTransactionFromBlockLatestNotFound struct {
	Payload *models.Error
}

func (o *GetTransactionFromBlockLatestNotFound) Error() string {
	return fmt.Sprintf("[GET /block/tx/latest/{tx_index}][%d] getTransactionFromBlockLatestNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionFromBlockLatestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
