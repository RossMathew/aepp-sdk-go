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

// NewGetTransactionFromBlockHeightParams creates a new GetTransactionFromBlockHeightParams object
// with the default values initialized.
func NewGetTransactionFromBlockHeightParams() *GetTransactionFromBlockHeightParams {
	var (
		txEncodingDefault = string("message_pack")
	)
	return &GetTransactionFromBlockHeightParams{
		TxEncoding: &txEncodingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransactionFromBlockHeightParamsWithTimeout creates a new GetTransactionFromBlockHeightParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTransactionFromBlockHeightParamsWithTimeout(timeout time.Duration) *GetTransactionFromBlockHeightParams {
	var (
		txEncodingDefault = string("message_pack")
	)
	return &GetTransactionFromBlockHeightParams{
		TxEncoding: &txEncodingDefault,

		timeout: timeout,
	}
}

// NewGetTransactionFromBlockHeightParamsWithContext creates a new GetTransactionFromBlockHeightParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTransactionFromBlockHeightParamsWithContext(ctx context.Context) *GetTransactionFromBlockHeightParams {
	var (
		txEncodingDefault = string("message_pack")
	)
	return &GetTransactionFromBlockHeightParams{
		TxEncoding: &txEncodingDefault,

		Context: ctx,
	}
}

// NewGetTransactionFromBlockHeightParamsWithHTTPClient creates a new GetTransactionFromBlockHeightParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTransactionFromBlockHeightParamsWithHTTPClient(client *http.Client) *GetTransactionFromBlockHeightParams {
	var (
		txEncodingDefault = string("message_pack")
	)
	return &GetTransactionFromBlockHeightParams{
		TxEncoding: &txEncodingDefault,
		HTTPClient: client,
	}
}

/*GetTransactionFromBlockHeightParams contains all the parameters to send to the API endpoint
for the get transaction from block height operation typically these are written to a http.Request
*/
type GetTransactionFromBlockHeightParams struct {

	/*Height
	  Height of the block to search for

	*/
	Height int64
	/*TxEncoding
	  Transactions encoding

	*/
	TxEncoding *string
	/*TxIndex
	  Index of the transaction in the block

	*/
	TxIndex int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) WithTimeout(timeout time.Duration) *GetTransactionFromBlockHeightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) WithContext(ctx context.Context) *GetTransactionFromBlockHeightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) WithHTTPClient(client *http.Client) *GetTransactionFromBlockHeightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) WithHeight(height int64) *GetTransactionFromBlockHeightParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) SetHeight(height int64) {
	o.Height = height
}

// WithTxEncoding adds the txEncoding to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) WithTxEncoding(txEncoding *string) *GetTransactionFromBlockHeightParams {
	o.SetTxEncoding(txEncoding)
	return o
}

// SetTxEncoding adds the txEncoding to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) SetTxEncoding(txEncoding *string) {
	o.TxEncoding = txEncoding
}

// WithTxIndex adds the txIndex to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) WithTxIndex(txIndex int64) *GetTransactionFromBlockHeightParams {
	o.SetTxIndex(txIndex)
	return o
}

// SetTxIndex adds the txIndex to the get transaction from block height params
func (o *GetTransactionFromBlockHeightParams) SetTxIndex(txIndex int64) {
	o.TxIndex = txIndex
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransactionFromBlockHeightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param height
	if err := r.SetPathParam("height", swag.FormatInt64(o.Height)); err != nil {
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

	// path param tx_index
	if err := r.SetPathParam("tx_index", swag.FormatInt64(o.TxIndex)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
