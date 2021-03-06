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

// PostSpendTxReader is a Reader for the PostSpendTx structure.
type PostSpendTxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSpendTxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostSpendTxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostSpendTxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSpendTxOK creates a PostSpendTxOK with default headers values
func NewPostSpendTxOK() *PostSpendTxOK {
	return &PostSpendTxOK{}
}

/*PostSpendTxOK handles this case with default header values.

successful operation
*/
type PostSpendTxOK struct {
}

func (o *PostSpendTxOK) Error() string {
	return fmt.Sprintf("[POST /spend-tx][%d] postSpendTxOK ", 200)
}

func (o *PostSpendTxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSpendTxNotFound creates a PostSpendTxNotFound with default headers values
func NewPostSpendTxNotFound() *PostSpendTxNotFound {
	return &PostSpendTxNotFound{}
}

/*PostSpendTxNotFound handles this case with default header values.

Spend transaction validation error
*/
type PostSpendTxNotFound struct {
	Payload *models.Error
}

func (o *PostSpendTxNotFound) Error() string {
	return fmt.Sprintf("[POST /spend-tx][%d] postSpendTxNotFound  %+v", 404, o.Payload)
}

func (o *PostSpendTxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
