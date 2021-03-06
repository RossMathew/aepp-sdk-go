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

// GetBlockTxsCountByHashReader is a Reader for the GetBlockTxsCountByHash structure.
type GetBlockTxsCountByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlockTxsCountByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlockTxsCountByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetBlockTxsCountByHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetBlockTxsCountByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlockTxsCountByHashOK creates a GetBlockTxsCountByHashOK with default headers values
func NewGetBlockTxsCountByHashOK() *GetBlockTxsCountByHashOK {
	return &GetBlockTxsCountByHashOK{}
}

/*GetBlockTxsCountByHashOK handles this case with default header values.

The count of transactions in the block
*/
type GetBlockTxsCountByHashOK struct {
	Payload *models.InlineResponse2003
}

func (o *GetBlockTxsCountByHashOK) Error() string {
	return fmt.Sprintf("[GET /block/txs/count/hash/{hash}][%d] getBlockTxsCountByHashOK  %+v", 200, o.Payload)
}

func (o *GetBlockTxsCountByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2003)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlockTxsCountByHashBadRequest creates a GetBlockTxsCountByHashBadRequest with default headers values
func NewGetBlockTxsCountByHashBadRequest() *GetBlockTxsCountByHashBadRequest {
	return &GetBlockTxsCountByHashBadRequest{}
}

/*GetBlockTxsCountByHashBadRequest handles this case with default header values.

Invalid hash
*/
type GetBlockTxsCountByHashBadRequest struct {
	Payload *models.Error
}

func (o *GetBlockTxsCountByHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /block/txs/count/hash/{hash}][%d] getBlockTxsCountByHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetBlockTxsCountByHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlockTxsCountByHashNotFound creates a GetBlockTxsCountByHashNotFound with default headers values
func NewGetBlockTxsCountByHashNotFound() *GetBlockTxsCountByHashNotFound {
	return &GetBlockTxsCountByHashNotFound{}
}

/*GetBlockTxsCountByHashNotFound handles this case with default header values.

Block not found
*/
type GetBlockTxsCountByHashNotFound struct {
	Payload *models.Error
}

func (o *GetBlockTxsCountByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /block/txs/count/hash/{hash}][%d] getBlockTxsCountByHashNotFound  %+v", 404, o.Payload)
}

func (o *GetBlockTxsCountByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
