// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetTransactionFromBlockHashReader is a Reader for the GetTransactionFromBlockHash structure.
type GetTransactionFromBlockHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionFromBlockHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionFromBlockHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetTransactionFromBlockHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionFromBlockHashOK creates a GetTransactionFromBlockHashOK with default headers values
func NewGetTransactionFromBlockHashOK() *GetTransactionFromBlockHashOK {
	return &GetTransactionFromBlockHashOK{}
}

/*GetTransactionFromBlockHashOK handles this case with default header values.

The transaction found
*/
type GetTransactionFromBlockHashOK struct {
	Payload models.SingleTxObject
}

func (o *GetTransactionFromBlockHashOK) Error() string {
	return fmt.Sprintf("[GET /block/tx/hash/{hash}/{tx_index}][%d] getTransactionFromBlockHashOK  %+v", 200, o.Payload)
}

func (o *GetTransactionFromBlockHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalSingleTxObject(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetTransactionFromBlockHashNotFound creates a GetTransactionFromBlockHashNotFound with default headers values
func NewGetTransactionFromBlockHashNotFound() *GetTransactionFromBlockHashNotFound {
	return &GetTransactionFromBlockHashNotFound{}
}

/*GetTransactionFromBlockHashNotFound handles this case with default header values.

Block or transaction not found
*/
type GetTransactionFromBlockHashNotFound struct {
	Payload *models.Error
}

func (o *GetTransactionFromBlockHashNotFound) Error() string {
	return fmt.Sprintf("[GET /block/tx/hash/{hash}/{tx_index}][%d] getTransactionFromBlockHashNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionFromBlockHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}