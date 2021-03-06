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

// NewPostNameUpdateTxParams creates a new PostNameUpdateTxParams object
// with the default values initialized.
func NewPostNameUpdateTxParams() *PostNameUpdateTxParams {
	var ()
	return &PostNameUpdateTxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNameUpdateTxParamsWithTimeout creates a new PostNameUpdateTxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNameUpdateTxParamsWithTimeout(timeout time.Duration) *PostNameUpdateTxParams {
	var ()
	return &PostNameUpdateTxParams{

		timeout: timeout,
	}
}

// NewPostNameUpdateTxParamsWithContext creates a new PostNameUpdateTxParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNameUpdateTxParamsWithContext(ctx context.Context) *PostNameUpdateTxParams {
	var ()
	return &PostNameUpdateTxParams{

		Context: ctx,
	}
}

// NewPostNameUpdateTxParamsWithHTTPClient creates a new PostNameUpdateTxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNameUpdateTxParamsWithHTTPClient(client *http.Client) *PostNameUpdateTxParams {
	var ()
	return &PostNameUpdateTxParams{
		HTTPClient: client,
	}
}

/*PostNameUpdateTxParams contains all the parameters to send to the API endpoint
for the post name update tx operation typically these are written to a http.Request
*/
type PostNameUpdateTxParams struct {

	/*Body
	  Creates new name update transaction

	*/
	Body *models.NameUpdateTx

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post name update tx params
func (o *PostNameUpdateTxParams) WithTimeout(timeout time.Duration) *PostNameUpdateTxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post name update tx params
func (o *PostNameUpdateTxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post name update tx params
func (o *PostNameUpdateTxParams) WithContext(ctx context.Context) *PostNameUpdateTxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post name update tx params
func (o *PostNameUpdateTxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post name update tx params
func (o *PostNameUpdateTxParams) WithHTTPClient(client *http.Client) *PostNameUpdateTxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post name update tx params
func (o *PostNameUpdateTxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post name update tx params
func (o *PostNameUpdateTxParams) WithBody(body *models.NameUpdateTx) *PostNameUpdateTxParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post name update tx params
func (o *PostNameUpdateTxParams) SetBody(body *models.NameUpdateTx) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostNameUpdateTxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
