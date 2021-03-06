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

// NewGetBlockTxsCountByHashParams creates a new GetBlockTxsCountByHashParams object
// with the default values initialized.
func NewGetBlockTxsCountByHashParams() *GetBlockTxsCountByHashParams {
	var ()
	return &GetBlockTxsCountByHashParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlockTxsCountByHashParamsWithTimeout creates a new GetBlockTxsCountByHashParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlockTxsCountByHashParamsWithTimeout(timeout time.Duration) *GetBlockTxsCountByHashParams {
	var ()
	return &GetBlockTxsCountByHashParams{

		timeout: timeout,
	}
}

// NewGetBlockTxsCountByHashParamsWithContext creates a new GetBlockTxsCountByHashParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlockTxsCountByHashParamsWithContext(ctx context.Context) *GetBlockTxsCountByHashParams {
	var ()
	return &GetBlockTxsCountByHashParams{

		Context: ctx,
	}
}

// NewGetBlockTxsCountByHashParamsWithHTTPClient creates a new GetBlockTxsCountByHashParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlockTxsCountByHashParamsWithHTTPClient(client *http.Client) *GetBlockTxsCountByHashParams {
	var ()
	return &GetBlockTxsCountByHashParams{
		HTTPClient: client,
	}
}

/*GetBlockTxsCountByHashParams contains all the parameters to send to the API endpoint
for the get block txs count by hash operation typically these are written to a http.Request
*/
type GetBlockTxsCountByHashParams struct {

	/*ExcludeTxTypes
	  Transactions types not to show. Comma separated. If a tx type appears in both tx_types and exclude_tx_types, it is excluded.

	*/
	ExcludeTxTypes *string
	/*Hash
	  Hash of the block to fetch

	*/
	Hash string
	/*TxTypes
	  Transactions types to show. Comma separated

	*/
	TxTypes *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) WithTimeout(timeout time.Duration) *GetBlockTxsCountByHashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) WithContext(ctx context.Context) *GetBlockTxsCountByHashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) WithHTTPClient(client *http.Client) *GetBlockTxsCountByHashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExcludeTxTypes adds the excludeTxTypes to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) WithExcludeTxTypes(excludeTxTypes *string) *GetBlockTxsCountByHashParams {
	o.SetExcludeTxTypes(excludeTxTypes)
	return o
}

// SetExcludeTxTypes adds the excludeTxTypes to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) SetExcludeTxTypes(excludeTxTypes *string) {
	o.ExcludeTxTypes = excludeTxTypes
}

// WithHash adds the hash to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) WithHash(hash string) *GetBlockTxsCountByHashParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) SetHash(hash string) {
	o.Hash = hash
}

// WithTxTypes adds the txTypes to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) WithTxTypes(txTypes *string) *GetBlockTxsCountByHashParams {
	o.SetTxTypes(txTypes)
	return o
}

// SetTxTypes adds the txTypes to the get block txs count by hash params
func (o *GetBlockTxsCountByHashParams) SetTxTypes(txTypes *string) {
	o.TxTypes = txTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlockTxsCountByHashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExcludeTxTypes != nil {

		// query param exclude_tx_types
		var qrExcludeTxTypes string
		if o.ExcludeTxTypes != nil {
			qrExcludeTxTypes = *o.ExcludeTxTypes
		}
		qExcludeTxTypes := qrExcludeTxTypes
		if qExcludeTxTypes != "" {
			if err := r.SetQueryParam("exclude_tx_types", qExcludeTxTypes); err != nil {
				return err
			}
		}

	}

	// path param hash
	if err := r.SetPathParam("hash", o.Hash); err != nil {
		return err
	}

	if o.TxTypes != nil {

		// query param tx_types
		var qrTxTypes string
		if o.TxTypes != nil {
			qrTxTypes = *o.TxTypes
		}
		qTxTypes := qrTxTypes
		if qTxTypes != "" {
			if err := r.SetQueryParam("tx_types", qTxTypes); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
