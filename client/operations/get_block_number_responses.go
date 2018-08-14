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

// GetBlockNumberReader is a Reader for the GetBlockNumber structure.
type GetBlockNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlockNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlockNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlockNumberOK creates a GetBlockNumberOK with default headers values
func NewGetBlockNumberOK() *GetBlockNumberOK {
	return &GetBlockNumberOK{}
}

/*GetBlockNumberOK handles this case with default header values.

The current block's height
*/
type GetBlockNumberOK struct {
	Payload *models.BlockHeight
}

func (o *GetBlockNumberOK) Error() string {
	return fmt.Sprintf("[GET /block/number][%d] getBlockNumberOK  %+v", 200, o.Payload)
}

func (o *GetBlockNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlockHeight)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}