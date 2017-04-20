package moneticopaiement

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPaymentResponseCurrency(t *testing.T) {
	p := PaymentResponse{Amount: "100USD"}
	assert.Equal(t, float64(100), p.Value())
	assert.Equal(t, "usd", p.Currency())

	p.Amount = ""
	assert.Equal(t, float64(0), p.Value())
	assert.Equal(t, "", p.Currency())
}

func TestPaymentResultDate(t *testing.T) {
	d := PaymentResultDate("15/04/2017_AM_10:15:30")
	tm := d.Time()
	assert.Equal(t, 2017, tm.Year())
	assert.Equal(t, time.April, tm.Month())
	assert.Equal(t, 15, tm.Day())
	assert.Equal(t, 10, tm.Hour())
	assert.Equal(t, 15, tm.Minute())
	assert.Equal(t, 30, tm.Second())
}
