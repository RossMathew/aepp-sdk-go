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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAccountBalanceParams creates a new GetAccountBalanceParams object
// with the default values initialized.
func NewGetAccountBalanceParams() *GetAccountBalanceParams {
	var ()
	return &GetAccountBalanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountBalanceParamsWithTimeout creates a new GetAccountBalanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountBalanceParamsWithTimeout(timeout time.Duration) *GetAccountBalanceParams {
	var ()
	return &GetAccountBalanceParams{

		timeout: timeout,
	}
}

// NewGetAccountBalanceParamsWithContext creates a new GetAccountBalanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountBalanceParamsWithContext(ctx context.Context) *GetAccountBalanceParams {
	var ()
	return &GetAccountBalanceParams{

		Context: ctx,
	}
}

// NewGetAccountBalanceParamsWithHTTPClient creates a new GetAccountBalanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountBalanceParamsWithHTTPClient(client *http.Client) *GetAccountBalanceParams {
	var ()
	return &GetAccountBalanceParams{
		HTTPClient: client,
	}
}

/*GetAccountBalanceParams contains all the parameters to send to the API endpoint
for the get account balance operation typically these are written to a http.Request
*/
type GetAccountBalanceParams struct {

	/*Address
	  Address to show balance for; it can be either account or contract

	*/
	Address string
	/*Hash
	  Hash of the block to show balance at

	*/
	Hash *string
	/*Height
	  Height of the block to show balance at

	*/
	Height *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account balance params
func (o *GetAccountBalanceParams) WithTimeout(timeout time.Duration) *GetAccountBalanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account balance params
func (o *GetAccountBalanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account balance params
func (o *GetAccountBalanceParams) WithContext(ctx context.Context) *GetAccountBalanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account balance params
func (o *GetAccountBalanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account balance params
func (o *GetAccountBalanceParams) WithHTTPClient(client *http.Client) *GetAccountBalanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account balance params
func (o *GetAccountBalanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the get account balance params
func (o *GetAccountBalanceParams) WithAddress(address string) *GetAccountBalanceParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the get account balance params
func (o *GetAccountBalanceParams) SetAddress(address string) {
	o.Address = address
}

// WithHash adds the hash to the get account balance params
func (o *GetAccountBalanceParams) WithHash(hash *string) *GetAccountBalanceParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the get account balance params
func (o *GetAccountBalanceParams) SetHash(hash *string) {
	o.Hash = hash
}

// WithHeight adds the height to the get account balance params
func (o *GetAccountBalanceParams) WithHeight(height *int64) *GetAccountBalanceParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the get account balance params
func (o *GetAccountBalanceParams) SetHeight(height *int64) {
	o.Height = height
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountBalanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if o.Hash != nil {

		// query param hash
		var qrHash string
		if o.Hash != nil {
			qrHash = *o.Hash
		}
		qHash := qrHash
		if qHash != "" {
			if err := r.SetQueryParam("hash", qHash); err != nil {
				return err
			}
		}

	}

	if o.Height != nil {

		// query param height
		var qrHeight int64
		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := swag.FormatInt64(qrHeight)
		if qHeight != "" {
			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
