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

// GetMicroBlockTransactionsCountByHashReader is a Reader for the GetMicroBlockTransactionsCountByHash structure.
type GetMicroBlockTransactionsCountByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMicroBlockTransactionsCountByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMicroBlockTransactionsCountByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetMicroBlockTransactionsCountByHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetMicroBlockTransactionsCountByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMicroBlockTransactionsCountByHashOK creates a GetMicroBlockTransactionsCountByHashOK with default headers values
func NewGetMicroBlockTransactionsCountByHashOK() *GetMicroBlockTransactionsCountByHashOK {
	return &GetMicroBlockTransactionsCountByHashOK{}
}

/*GetMicroBlockTransactionsCountByHashOK handles this case with default header values.

Successful operation
*/
type GetMicroBlockTransactionsCountByHashOK struct {
	Payload *models.InlineResponse2002
}

func (o *GetMicroBlockTransactionsCountByHashOK) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/transactions/count][%d] getMicroBlockTransactionsCountByHashOK  %+v", 200, o.Payload)
}

func (o *GetMicroBlockTransactionsCountByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2002)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicroBlockTransactionsCountByHashBadRequest creates a GetMicroBlockTransactionsCountByHashBadRequest with default headers values
func NewGetMicroBlockTransactionsCountByHashBadRequest() *GetMicroBlockTransactionsCountByHashBadRequest {
	return &GetMicroBlockTransactionsCountByHashBadRequest{}
}

/*GetMicroBlockTransactionsCountByHashBadRequest handles this case with default header values.

Invalid hash
*/
type GetMicroBlockTransactionsCountByHashBadRequest struct {
	Payload *models.Error
}

func (o *GetMicroBlockTransactionsCountByHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/transactions/count][%d] getMicroBlockTransactionsCountByHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetMicroBlockTransactionsCountByHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicroBlockTransactionsCountByHashNotFound creates a GetMicroBlockTransactionsCountByHashNotFound with default headers values
func NewGetMicroBlockTransactionsCountByHashNotFound() *GetMicroBlockTransactionsCountByHashNotFound {
	return &GetMicroBlockTransactionsCountByHashNotFound{}
}

/*GetMicroBlockTransactionsCountByHashNotFound handles this case with default header values.

Block not found
*/
type GetMicroBlockTransactionsCountByHashNotFound struct {
	Payload *models.Error
}

func (o *GetMicroBlockTransactionsCountByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/transactions/count][%d] getMicroBlockTransactionsCountByHashNotFound  %+v", 404, o.Payload)
}

func (o *GetMicroBlockTransactionsCountByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
