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

// GetBlockGenesisReader is a Reader for the GetBlockGenesis structure.
type GetBlockGenesisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlockGenesisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlockGenesisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlockGenesisOK creates a GetBlockGenesisOK with default headers values
func NewGetBlockGenesisOK() *GetBlockGenesisOK {
	return &GetBlockGenesisOK{}
}

/*GetBlockGenesisOK handles this case with default header values.

The genesis block
*/
type GetBlockGenesisOK struct {
	Payload *models.GenericBlock
}

func (o *GetBlockGenesisOK) Error() string {
	return fmt.Sprintf("[GET /block/genesis][%d] getBlockGenesisOK  %+v", 200, o.Payload)
}

func (o *GetBlockGenesisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericBlock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
