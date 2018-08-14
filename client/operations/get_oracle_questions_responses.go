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

// GetOracleQuestionsReader is a Reader for the GetOracleQuestions structure.
type GetOracleQuestionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleQuestionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOracleQuestionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOracleQuestionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOracleQuestionsOK creates a GetOracleQuestionsOK with default headers values
func NewGetOracleQuestionsOK() *GetOracleQuestionsOK {
	return &GetOracleQuestionsOK{}
}

/*GetOracleQuestionsOK handles this case with default header values.

Active oracle questions
*/
type GetOracleQuestionsOK struct {
	Payload models.OracleQuestions
}

func (o *GetOracleQuestionsOK) Error() string {
	return fmt.Sprintf("[GET /oracle-questions][%d] getOracleQuestionsOK  %+v", 200, o.Payload)
}

func (o *GetOracleQuestionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleQuestionsBadRequest creates a GetOracleQuestionsBadRequest with default headers values
func NewGetOracleQuestionsBadRequest() *GetOracleQuestionsBadRequest {
	return &GetOracleQuestionsBadRequest{}
}

/*GetOracleQuestionsBadRequest handles this case with default header values.

Invalid parameters
*/
type GetOracleQuestionsBadRequest struct {
	Payload *models.Error
}

func (o *GetOracleQuestionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracle-questions][%d] getOracleQuestionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOracleQuestionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}