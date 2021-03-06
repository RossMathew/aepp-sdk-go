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

// NewPostNameTransferTxParams creates a new PostNameTransferTxParams object
// with the default values initialized.
func NewPostNameTransferTxParams() *PostNameTransferTxParams {
	var ()
	return &PostNameTransferTxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNameTransferTxParamsWithTimeout creates a new PostNameTransferTxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNameTransferTxParamsWithTimeout(timeout time.Duration) *PostNameTransferTxParams {
	var ()
	return &PostNameTransferTxParams{

		timeout: timeout,
	}
}

// NewPostNameTransferTxParamsWithContext creates a new PostNameTransferTxParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNameTransferTxParamsWithContext(ctx context.Context) *PostNameTransferTxParams {
	var ()
	return &PostNameTransferTxParams{

		Context: ctx,
	}
}

// NewPostNameTransferTxParamsWithHTTPClient creates a new PostNameTransferTxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNameTransferTxParamsWithHTTPClient(client *http.Client) *PostNameTransferTxParams {
	var ()
	return &PostNameTransferTxParams{
		HTTPClient: client,
	}
}

/*PostNameTransferTxParams contains all the parameters to send to the API endpoint
for the post name transfer tx operation typically these are written to a http.Request
*/
type PostNameTransferTxParams struct {

	/*Body
	  Creates new name transfer transaction

	*/
	Body *models.NameTransferTx

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post name transfer tx params
func (o *PostNameTransferTxParams) WithTimeout(timeout time.Duration) *PostNameTransferTxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post name transfer tx params
func (o *PostNameTransferTxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post name transfer tx params
func (o *PostNameTransferTxParams) WithContext(ctx context.Context) *PostNameTransferTxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post name transfer tx params
func (o *PostNameTransferTxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post name transfer tx params
func (o *PostNameTransferTxParams) WithHTTPClient(client *http.Client) *PostNameTransferTxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post name transfer tx params
func (o *PostNameTransferTxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post name transfer tx params
func (o *PostNameTransferTxParams) WithBody(body *models.NameTransferTx) *PostNameTransferTxParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post name transfer tx params
func (o *PostNameTransferTxParams) SetBody(body *models.NameTransferTx) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostNameTransferTxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
