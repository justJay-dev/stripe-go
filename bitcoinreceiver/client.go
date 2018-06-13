// Package bitcoinreceiver provides the /bitcoin/receivers APIs.
//
// Note that this entire package is deprecated. Please use the new sources API
// instead.
package bitcoinreceiver

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /bitcoin/receivers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new bitcoin receivers.
// For more details see https://stripe.com/docs/api/#create_bitcoin_receiver
func New(params *stripe.BitcoinReceiverParams) (*stripe.BitcoinReceiver, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.BitcoinReceiverParams) (*stripe.BitcoinReceiver, error) {
	receiver := &stripe.BitcoinReceiver{}
	err := c.B.Call("POST", "/bitcoin/receivers", c.Key, params, receiver)
	return receiver, err
}

// Get returns the details of a bitcoin receiver.
// For more details see https://stripe.com/docs/api/#retrieve_bitcoin_receiver
func Get(id string, params *stripe.BitcoinReceiverParams) (*stripe.BitcoinReceiver, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.BitcoinReceiverParams) (*stripe.BitcoinReceiver, error) {
	path := stripe.FormatURLPath("/bitcoin/receivers/%s", id)
	bitcoinReceiver := &stripe.BitcoinReceiver{}
	err := c.B.Call("GET", path, c.Key, params, bitcoinReceiver)
	return bitcoinReceiver, err
}

// Update updates a bitcoin receiver's properties.
// For more details see https://stripe.com/docs/api#update_bitcoin_receiver.
func Update(id string, params *stripe.BitcoinReceiverUpdateParams) (*stripe.BitcoinReceiver, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.BitcoinReceiverUpdateParams) (*stripe.BitcoinReceiver, error) {
	path := stripe.FormatURLPath("/bitcoin/receivers/%s", id)
	receiver := &stripe.BitcoinReceiver{}
	err := c.B.Call("POST", path, c.Key, params, receiver)
	return receiver, err
}

// List returns a list of bitcoin receivers.
// For more details see https://stripe.com/docs/api/#list_bitcoin_receivers
func List(params *stripe.BitcoinReceiverListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.BitcoinReceiverListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.BitcoinReceiverList{}
		err := c.B.CallRaw("GET", "/bitcoin/receivers", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of BitcoinReceivers.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// BitcoinReceiver returns the most recent BitcoinReceiver
// visited by a call to Next.
func (i *Iter) BitcoinReceiver() *stripe.BitcoinReceiver {
	return i.Current().(*stripe.BitcoinReceiver)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
