package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 10.0, tax, "Tax should be 10.0")

	tax, err = CalculateTax(0)
	assert.Error(t, err, "Amount must be greater than 0")
	//assert.Equal(t, 0.0, tax, "Tax should be 0.0")
	assert.Contains(t, err.Error(), "greater than 0")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	//quando receber vlr 10.0 no repo, retornar nil para o erro

	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
	//	a funcao saveTax devera ser chamada 3x
}
