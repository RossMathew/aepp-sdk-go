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

// PostChannelDepositReader is a Reader for the PostChannelDeposit structure.
type PostChannelDepositReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostChannelDepositReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostChannelDepositOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostChannelDepositBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostChannelDepositOK creates a PostChannelDepositOK with default headers values
func NewPostChannelDepositOK() *PostChannelDepositOK {
	return &PostChannelDepositOK{}
}

/*PostChannelDepositOK handles this case with default header values.

Successful operation
*/
type PostChannelDepositOK struct {
	Payload *models.UnsignedTx
}

func (o *PostChannelDepositOK) Error() string {
	return fmt.Sprintf("[POST /tx/channel/deposit][%d] postChannelDepositOK  %+v", 200, o.Payload)
}

func (o *PostChannelDepositOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnsignedTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostChannelDepositBadRequest creates a PostChannelDepositBadRequest with default headers values
func NewPostChannelDepositBadRequest() *PostChannelDepositBadRequest {
	return &PostChannelDepositBadRequest{}
}

/*PostChannelDepositBadRequest handles this case with default header values.

Invalid transaction
*/
type PostChannelDepositBadRequest struct {
	Payload *models.Error
}

func (o *PostChannelDepositBadRequest) Error() string {
	return fmt.Sprintf("[POST /tx/channel/deposit][%d] postChannelDepositBadRequest  %+v", 400, o.Payload)
}

func (o *PostChannelDepositBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
