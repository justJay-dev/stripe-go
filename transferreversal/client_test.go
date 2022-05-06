package transferreversal

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v72"
	_ "github.com/stripe/stripe-go/v72/testing"
)

func TestReversalGet(t *testing.T) {
	reversal, err := Get("trr_123", &stripe.TransferReversalParams{
		Transfer: stripe.String("tr_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}

func TestReversalList(t *testing.T) {
	i := List(&stripe.TransferReversalListParams{
		Transfer: stripe.String("tr_123"),
	})

	// Verify that we can get at least one reversal
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.TransferReversal())
	assert.NotNil(t, i.TransferReversalList())
}

func TestReversalNew(t *testing.T) {
	reversal, err := New(&stripe.TransferReversalParams{
		Amount:   stripe.Int64(123),
		Transfer: stripe.String("tr_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}

func TestReversalUpdate(t *testing.T) {
	reversal, err := Update("trr_123", &stripe.TransferReversalParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
		Transfer: stripe.String("tr_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}
