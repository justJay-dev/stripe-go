package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestOrder_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Order
		err := json.Unmarshal([]byte(`"or_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "or_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Order{ID: "or_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "or_123", v.ID)
	}
}

func TestOrderItem_UnmarshalJSON(t *testing.T) {
	// Try unmarshaling a SKU order item
	{
		orderItemData := map[string]interface{}{
			"object": "order_item",
			"amount": 123,
			"parent": "TEST-SKU-123",
			"type":   "sku",
		}
		bytes, err := json.Marshal(&orderItemData)
		assert.NoError(t, err)

		var orderItem OrderItem
		err = json.Unmarshal(bytes, &orderItem)
		assert.NoError(t, err)
		assert.Equal(t, "TEST-SKU-123", orderItem.Parent.ID)
	}

	// Try unmarshaling a SKU order item with parent expanded
	{
		orderItemData := map[string]interface{}{
			"object": "order_item",
			"amount": 123,
			"parent": map[string]interface{}{
				"id":     "TEST-SKU-123",
				"object": "sku",
			},
			"type": "sku",
		}
		bytes, err := json.Marshal(&orderItemData)
		assert.NoError(t, err)

		var orderItem OrderItem
		err = json.Unmarshal(bytes, &orderItem)
		assert.NoError(t, err)
		assert.Equal(t, "TEST-SKU-123", orderItem.Parent.ID)
		assert.Equal(t, "sku", orderItem.Parent.Object)
		assert.Equal(t, "TEST-SKU-123", orderItem.Parent.ID)
	}

	// Try unmarshaling a coupon order item
	{
		orderItemData := map[string]interface{}{
			"object": "order_item",
			"amount": 0,
			"parent": "TEST-COUPON-123",
			"type":   "coupon",
		}
		bytes, err := json.Marshal(&orderItemData)
		assert.NoError(t, err)

		var orderItem OrderItem
		err = json.Unmarshal(bytes, &orderItem)
		assert.NoError(t, err)
		assert.Equal(t, "TEST-COUPON-123", orderItem.Parent.ID)
	}

	// Try unmarshaling a shipping order item
	{
		orderItemData := map[string]interface{}{
			"object": "order_item",
			"amount": 1000,
			"parent": "ship_MZmIpV7v14QZLlRR",
			"type":   "shipping",
		}
		bytes, err := json.Marshal(&orderItemData)
		assert.NoError(t, err)

		var orderItem OrderItem
		err = json.Unmarshal(bytes, &orderItem)
		assert.NoError(t, err)
		assert.Equal(t, "ship_MZmIpV7v14QZLlRR", orderItem.Parent.ID)
	}
}
