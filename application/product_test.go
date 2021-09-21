package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Junior Rabelo"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enabled()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enabled()
	require.Equal(t, "the price most be greater than zero to enable the product", err.Error())
}
