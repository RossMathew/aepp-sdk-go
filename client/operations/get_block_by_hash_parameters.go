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

// NewGetBlockByHashParams creates a new GetBlockByHashParams object
// with the default values initialized.
func NewGetBlockByHashParams() *GetBlockByHashParams {
	var (
		txEncodingDefault = string("json")
	)
	return &GetBlockByHashParams{
		TxEncoding: &txEncodingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlockByHashParamsWithTimeout creates a new GetBlockByHashParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlockByHashParamsWithTimeout(timeout time.Duration) *GetBlockByHashParams {
	var (
		txEncodingDefault = string("json")
	)
	return &GetBlockByHashParams{
		TxEncoding: &txEncodingDefault,

		timeout: timeout,
	}
}

// NewGetBlockByHashParamsWithContext creates a new GetBlockByHashParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlockByHashParamsWithContext(ctx context.Context) *GetBlockByHashParams {
	var (
		txEncodingDefault = string("json")
	)
	return &GetBlockByHashParams{
		TxEncoding: &txEncodingDefault,

		Context: ctx,
	}
}

// NewGetBlockByHashParamsWithHTTPClient creates a new GetBlockByHashParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlockByHashParamsWithHTTPClient(client *http.Client) *GetBlockByHashParams {
	var (
		txEncodingDefault = string("json")
	)
	return &GetBlockByHashParams{
		TxEncoding: &txEncodingDefault,
		HTTPClient: client,
	}
}

/*GetBlockByHashParams contains all the parameters to send to the API endpoint
for the get block by hash operation typically these are written to a http.Request
*/
type GetBlockByHashParams struct {

	/*Hash
	  Hash of the block to fetch

	*/
	Hash string
	/*TxEncoding
	  Transactions encoding

	*/
	TxEncoding *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get block by hash params
func (o *GetBlockByHashParams) WithTimeout(timeout time.Duration) *GetBlockByHashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get block by hash params
func (o *GetBlockByHashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get block by hash params
func (o *GetBlockByHashParams) WithContext(ctx context.Context) *GetBlockByHashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get block by hash params
func (o *GetBlockByHashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get block by hash params
func (o *GetBlockByHashParams) WithHTTPClient(client *http.Client) *GetBlockByHashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get block by hash params
func (o *GetBlockByHashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHash adds the hash to the get block by hash params
func (o *GetBlockByHashParams) WithHash(hash string) *GetBlockByHashParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the get block by hash params
func (o *GetBlockByHashParams) SetHash(hash string) {
	o.Hash = hash
}

// WithTxEncoding adds the txEncoding to the get block by hash params
func (o *GetBlockByHashParams) WithTxEncoding(txEncoding *string) *GetBlockByHashParams {
	o.SetTxEncoding(txEncoding)
	return o
}

// SetTxEncoding adds the txEncoding to the get block by hash params
func (o *GetBlockByHashParams) SetTxEncoding(txEncoding *string) {
	o.TxEncoding = txEncoding
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlockByHashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param hash
	if err := r.SetPathParam("hash", o.Hash); err != nil {
		return err
	}

	if o.TxEncoding != nil {

		// query param tx_encoding
		var qrTxEncoding string
		if o.TxEncoding != nil {
			qrTxEncoding = *o.TxEncoding
		}
		qTxEncoding := qrTxEncoding
		if qTxEncoding != "" {
			if err := r.SetQueryParam("tx_encoding", qTxEncoding); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
