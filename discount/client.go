//
//
// File generated from our OpenAPI spec
//
//

// Package discount provides the TODO APIs
package discount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke TODO APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Del is the method for the `DELETE /v1/customers/{customer}/discount` API.
func Del(id string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	return getC().Del(id, params)
}

// Del is the method for the `DELETE /v1/customers/{customer}/discount` API.
func (c Client) Del(id string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	path := stripe.FormatURLPath("/v1/customers/%s/discount", id)
	discount := &stripe.Discount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, discount)
	return discount, err
}

// DelSub is the method for the `DELETE /v1/subscriptions/{subscription_exposed_id}/discount` API.
func DelSub(id string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	return getC().DelSub(id, params)
}

// DelSub is the method for the `DELETE /v1/subscriptions/{subscription_exposed_id}/discount` API.
func (c Client) DelSub(id string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s/discount", id)
	discount := &stripe.Discount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, discount)
	return discount, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
