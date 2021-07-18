package pricing

import (
	"errors"

	"github.com/mksoni7001/simple-go-service/pkg/models"
)

var (
	Items        = make(map[string]models.Item, 0)
	Rules        = make(map[string]models.PricingRule, 0)
	cartQuantity = make(map[string]uint, 0)
)

func (s *service) AddItem(item models.Item) (saved bool, err error) {
	// Validations
	if item == (models.Item{}) {
		err = errors.New("Empty item details")
		return
	}
	if item.SKU == "" {
		err = errors.New("SKU details is mandatory")
		return
	}
	Items[item.SKU] = item
	saved = true
	return
}

func (s *service) Scan(item string, quantity uint) {
	if _, ok := cartQuantity[item]; ok {
		cartQuantity[item] += quantity
	} else {
		cartQuantity[item] = quantity
	}
}

func RemoveItem(item string, quantity uint) {
	if _, ok := cartQuantity[item]; ok {
		if remainingQuantity := cartQuantity[item] - quantity; remainingQuantity > 0 {
			cartQuantity[item] = remainingQuantity
		} else {
			delete(cartQuantity, item)
		}
	}
}

func (s *service) AddPricingRule(rule models.PricingRule) {
	Rules[rule.SKU] = rule
}

func (s *service) Total() (total float64, err error) {
	if len(Items) == 0 {
		err = errors.New("No items in cart")
		return
	}
	// Calculate Total
	for item, quantity := range cartQuantity {
		if rule, ok := Rules[item]; ok {
			if quantity >= rule.Quantity {
				total += float64(rule.SpecialPrice) * float64(quantity/rule.Quantity)
				if individualUnits := quantity % rule.Quantity; individualUnits > 0 {
					total += float64(individualUnits) * float64(Items[item].UnitPrice)
				}
			} else {
				total += Items[item].UnitPrice * float64(quantity)
			}
		} else {
			total += Items[item].UnitPrice * float64(quantity)
		}
	}
	return
}

func (s *service) ClearCart() {
	cartQuantity = make(map[string]uint, 0)
}
