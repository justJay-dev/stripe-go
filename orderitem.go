//
//
// File generated from our OpenAPI spec
//
//

package stripe

// A representation of the constituent items of any given order. Can be used to
// represent [SKUs](https://stripe.com/docs/api#skus), shipping costs, or taxes owed on the order.
//
// Related guide: [Orders](https://stripe.com/docs/orders/guide).
type OrderItem struct {
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount for the line item.
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Description of the line item, meant to be displayable to the user (e.g., `"Express shipping"`).
	Description string `json:"description"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ID of the associated object for this line item. Expandable if not null (e.g., expandable to a SKU).
	Parent *SKU `json:"parent"`
	// A positive integer representing the number of instances of `parent` that are included in this order item. Applicable/present only if `type` is `sku`.
	Quantity int64 `json:"quantity"`
	// The type of line item. One of `sku`, `tax`, `shipping`, or `discount`.
	Type string `json:"type"`
}
