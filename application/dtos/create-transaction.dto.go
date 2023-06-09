package dto

import "fmt"

type CreateTransactionDTO struct {
	CardNumber      string   `json:"card_number" validate:"required"`
	CVV             string   `json:"cvv" validate:"required"`
	ExpirationMonth string   `json:"month" validate:"required"`
	ExpirationYear  string   `json:"year" validate:"required"`
	CardBrand       string   `json:"brand" validate:"required"`
	Holder          string   `json:"holder" validate:"required"`
	Amount          int      `json:"amount" validate:"required"`
	Installments    int      `json:"installments" validate:"required"`
	Type            string   `json:"type" validate:"required"`
	OrderID         string   `json:"order_id" validate:"required"`
	Customer        Customer `json:"customer"`
}

type Customer struct {
	Name string `json:"name" validate:"required"`
}

func (r *CreateTransactionDTO) Validate() error {
	if r.CardNumber == "" {
		return validatePropsIsRequire("card_number", "string")
	}
	if r.CVV == "" {
		return validatePropsIsRequire("cvv", "string")
	}
	if r.ExpirationMonth == "" {
		return validatePropsIsRequire("expireation month", "string")
	}
	if r.ExpirationYear == "" {
		return validatePropsIsRequire("expiration year", "string")
	}
	if r.CardBrand == "" {
		return validatePropsIsRequire("brand", "string")
	}
	if r.Holder == "" {
		return validatePropsIsRequire("holder", "string")
	}

	return nil
}

func validatePropsIsRequire(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
