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

// CallContractReader is a Reader for the CallContract structure.
type CallContractReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CallContractReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCallContractOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewCallContractForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCallContractOK creates a CallContractOK with default headers values
func NewCallContractOK() *CallContractOK {
	return &CallContractOK{}
}

/*CallContractOK handles this case with default header values.

Resulting state map
*/
type CallContractOK struct {
	Payload *models.CallResult
}

func (o *CallContractOK) Error() string {
	return fmt.Sprintf("[POST /contract/call][%d] callContractOK  %+v", 200, o.Payload)
}

func (o *CallContractOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCallContractForbidden creates a CallContractForbidden with default headers values
func NewCallContractForbidden() *CallContractForbidden {
	return &CallContractForbidden{}
}

/*CallContractForbidden handles this case with default header values.

Invalid contract
*/
type CallContractForbidden struct {
	Payload *models.Error
}

func (o *CallContractForbidden) Error() string {
	return fmt.Sprintf("[POST /contract/call][%d] callContractForbidden  %+v", 403, o.Payload)
}

func (o *CallContractForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
