// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/models"
)

// NewPostTxParams creates a new PostTxParams object
// with the default values initialized.
func NewPostTxParams() *PostTxParams {
	var ()
	return &PostTxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTxParamsWithTimeout creates a new PostTxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTxParamsWithTimeout(timeout time.Duration) *PostTxParams {
	var ()
	return &PostTxParams{

		timeout: timeout,
	}
}

// NewPostTxParamsWithContext creates a new PostTxParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTxParamsWithContext(ctx context.Context) *PostTxParams {
	var ()
	return &PostTxParams{

		Context: ctx,
	}
}

// NewPostTxParamsWithHTTPClient creates a new PostTxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTxParamsWithHTTPClient(client *http.Client) *PostTxParams {
	var ()
	return &PostTxParams{
		HTTPClient: client,
	}
}

/*PostTxParams contains all the parameters to send to the API endpoint
for the post tx operation typically these are written to a http.Request
*/
type PostTxParams struct {

	/*Body
	  A new transaction

	*/
	Body *models.Tx

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post tx params
func (o *PostTxParams) WithTimeout(timeout time.Duration) *PostTxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post tx params
func (o *PostTxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post tx params
func (o *PostTxParams) WithContext(ctx context.Context) *PostTxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post tx params
func (o *PostTxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post tx params
func (o *PostTxParams) WithHTTPClient(client *http.Client) *PostTxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post tx params
func (o *PostTxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post tx params
func (o *PostTxParams) WithBody(body *models.Tx) *PostTxParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post tx params
func (o *PostTxParams) SetBody(body *models.Tx) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
