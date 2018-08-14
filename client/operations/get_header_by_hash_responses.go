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

// GetHeaderByHashReader is a Reader for the GetHeaderByHash structure.
type GetHeaderByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHeaderByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHeaderByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetHeaderByHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetHeaderByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHeaderByHashOK creates a GetHeaderByHashOK with default headers values
func NewGetHeaderByHashOK() *GetHeaderByHashOK {
	return &GetHeaderByHashOK{}
}

/*GetHeaderByHashOK handles this case with default header values.

The header found
*/
type GetHeaderByHashOK struct {
	Payload *models.Header
}

func (o *GetHeaderByHashOK) Error() string {
	return fmt.Sprintf("[GET /header-by-hash][%d] getHeaderByHashOK  %+v", 200, o.Payload)
}

func (o *GetHeaderByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Header)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHeaderByHashBadRequest creates a GetHeaderByHashBadRequest with default headers values
func NewGetHeaderByHashBadRequest() *GetHeaderByHashBadRequest {
	return &GetHeaderByHashBadRequest{}
}

/*GetHeaderByHashBadRequest handles this case with default header values.

Invalid hash
*/
type GetHeaderByHashBadRequest struct {
	Payload *models.Error
}

func (o *GetHeaderByHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /header-by-hash][%d] getHeaderByHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetHeaderByHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHeaderByHashNotFound creates a GetHeaderByHashNotFound with default headers values
func NewGetHeaderByHashNotFound() *GetHeaderByHashNotFound {
	return &GetHeaderByHashNotFound{}
}

/*GetHeaderByHashNotFound handles this case with default header values.

Header not found
*/
type GetHeaderByHashNotFound struct {
	Payload *models.Error
}

func (o *GetHeaderByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /header-by-hash][%d] getHeaderByHashNotFound  %+v", 404, o.Payload)
}

func (o *GetHeaderByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}