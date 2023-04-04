package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order := &Order{ID: id, Price: price, Tax: tax}
	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {

	if o.ID == "" {
		return errors.New("Id is required")
	}
	if o.Price == 0 {
		return errors.New("Price is required")
	}
	if o.Tax == 0 {
		return errors.New("Tax is required")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.Validate()
	if err != nil {
		return err
	}
	return nil

}
