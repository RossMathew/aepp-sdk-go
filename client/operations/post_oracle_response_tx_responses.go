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

// PostOracleResponseTxReader is a Reader for the PostOracleResponseTx structure.
type PostOracleResponseTxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOracleResponseTxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostOracleResponseTxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostOracleResponseTxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOracleResponseTxOK creates a PostOracleResponseTxOK with default headers values
func NewPostOracleResponseTxOK() *PostOracleResponseTxOK {
	return &PostOracleResponseTxOK{}
}

/*PostOracleResponseTxOK handles this case with default header values.

successful operation
*/
type PostOracleResponseTxOK struct {
	Payload *models.OracleQueryResponse
}

func (o *PostOracleResponseTxOK) Error() string {
	return fmt.Sprintf("[POST /oracle-response-tx][%d] postOracleResponseTxOK  %+v", 200, o.Payload)
}

func (o *PostOracleResponseTxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OracleQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOracleResponseTxNotFound creates a PostOracleResponseTxNotFound with default headers values
func NewPostOracleResponseTxNotFound() *PostOracleResponseTxNotFound {
	return &PostOracleResponseTxNotFound{}
}

/*PostOracleResponseTxNotFound handles this case with default header values.

Oracle response transaction validation error
*/
type PostOracleResponseTxNotFound struct {
	Payload *models.Error
}

func (o *PostOracleResponseTxNotFound) Error() string {
	return fmt.Sprintf("[POST /oracle-response-tx][%d] postOracleResponseTxNotFound  %+v", 404, o.Payload)
}

func (o *PostOracleResponseTxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
