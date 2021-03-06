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

// PostOracleExtendTxReader is a Reader for the PostOracleExtendTx structure.
type PostOracleExtendTxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOracleExtendTxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostOracleExtendTxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostOracleExtendTxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOracleExtendTxOK creates a PostOracleExtendTxOK with default headers values
func NewPostOracleExtendTxOK() *PostOracleExtendTxOK {
	return &PostOracleExtendTxOK{}
}

/*PostOracleExtendTxOK handles this case with default header values.

successful operation
*/
type PostOracleExtendTxOK struct {
	Payload *models.OracleRegisterResponse
}

func (o *PostOracleExtendTxOK) Error() string {
	return fmt.Sprintf("[POST /oracle-extend-tx][%d] postOracleExtendTxOK  %+v", 200, o.Payload)
}

func (o *PostOracleExtendTxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OracleRegisterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOracleExtendTxNotFound creates a PostOracleExtendTxNotFound with default headers values
func NewPostOracleExtendTxNotFound() *PostOracleExtendTxNotFound {
	return &PostOracleExtendTxNotFound{}
}

/*PostOracleExtendTxNotFound handles this case with default header values.

Oracle register transaction validation error
*/
type PostOracleExtendTxNotFound struct {
	Payload *models.Error
}

func (o *PostOracleExtendTxNotFound) Error() string {
	return fmt.Sprintf("[POST /oracle-extend-tx][%d] postOracleExtendTxNotFound  %+v", 404, o.Payload)
}

func (o *PostOracleExtendTxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
