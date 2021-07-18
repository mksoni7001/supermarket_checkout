package pricing

import "github.com/mksoni7001/simple-go-service/pkg/models"

type Checkout interface {
	AddItem(item models.Item) (saved bool, err error)
	Scan(item string, count uint)
	AddPricingRule(rule models.PricingRule)
	Total() (total float64, err error)
	ClearCart()
}

type service struct {
}

func New() Checkout {
	return &service{}
}
