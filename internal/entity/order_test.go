package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "1234"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "1234", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}
func TestFinalPrice(t *testing.T) {
	order := Order{ID: "1234", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "1234", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
