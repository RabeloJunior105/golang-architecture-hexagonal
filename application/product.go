package application

import "errors"

type IProduct interface {
	IsValid() (bool, error)
	Enabled() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

//func (p *Product) IsValid() (bool, error) {

//}

func (p *Product) Enabled() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price most be greater than zero to enable the product")
}

//func (p *Product) Disable() error {

//}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
