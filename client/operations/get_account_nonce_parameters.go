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
)

// NewGetAccountNonceParams creates a new GetAccountNonceParams object
// with the default values initialized.
func NewGetAccountNonceParams() *GetAccountNonceParams {
	var ()
	return &GetAccountNonceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountNonceParamsWithTimeout creates a new GetAccountNonceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountNonceParamsWithTimeout(timeout time.Duration) *GetAccountNonceParams {
	var ()
	return &GetAccountNonceParams{

		timeout: timeout,
	}
}

// NewGetAccountNonceParamsWithContext creates a new GetAccountNonceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountNonceParamsWithContext(ctx context.Context) *GetAccountNonceParams {
	var ()
	return &GetAccountNonceParams{

		Context: ctx,
	}
}

// NewGetAccountNonceParamsWithHTTPClient creates a new GetAccountNonceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountNonceParamsWithHTTPClient(client *http.Client) *GetAccountNonceParams {
	var ()
	return &GetAccountNonceParams{
		HTTPClient: client,
	}
}

/*GetAccountNonceParams contains all the parameters to send to the API endpoint
for the get account nonce operation typically these are written to a http.Request
*/
type GetAccountNonceParams struct {

	/*AccountPubkey
	  Account pubkey to show nonce for

	*/
	AccountPubkey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account nonce params
func (o *GetAccountNonceParams) WithTimeout(timeout time.Duration) *GetAccountNonceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account nonce params
func (o *GetAccountNonceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account nonce params
func (o *GetAccountNonceParams) WithContext(ctx context.Context) *GetAccountNonceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account nonce params
func (o *GetAccountNonceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account nonce params
func (o *GetAccountNonceParams) WithHTTPClient(client *http.Client) *GetAccountNonceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account nonce params
func (o *GetAccountNonceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountPubkey adds the accountPubkey to the get account nonce params
func (o *GetAccountNonceParams) WithAccountPubkey(accountPubkey string) *GetAccountNonceParams {
	o.SetAccountPubkey(accountPubkey)
	return o
}

// SetAccountPubkey adds the accountPubkey to the get account nonce params
func (o *GetAccountNonceParams) SetAccountPubkey(accountPubkey string) {
	o.AccountPubkey = accountPubkey
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountNonceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_pubkey
	if err := r.SetPathParam("account_pubkey", o.AccountPubkey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
