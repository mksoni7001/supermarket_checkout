package tests

import (
	"testing"

	"github.com/mksoni7001/simple-go-service/pkg/models"
	"github.com/mksoni7001/simple-go-service/pkg/pricing"
)

func TestPricing(t *testing.T) {
	pricingService := pricing.New()
	pricingService.AddItem(models.Item{
		SKU:       "A",
		UnitPrice: 50,
	})
	pricingService.AddItem(models.Item{
		SKU:       "B",
		UnitPrice: 30,
	})
	pricingService.AddItem(models.Item{
		SKU:       "C",
		UnitPrice: 20,
	})
	pricingService.AddItem(models.Item{
		SKU:       "D",
		UnitPrice: 15,
	})
	pricingService.AddPricingRule(models.PricingRule{
		SKU:          "A",
		Quantity:     3,
		SpecialPrice: 130,
	})
	pricingService.AddPricingRule(models.PricingRule{
		SKU:          "B",
		Quantity:     2,
		SpecialPrice: 45,
	})
	t.Run("Should calculate correctly when no special price applicable", func(t *testing.T) {
		pricingService.ClearCart()
		pricingService.Scan("A", 1)
		pricingService.Scan("B", 1)
		total, err := pricingService.Total()
		if err != nil || total != 80 {
			t.Errorf("Total should be 80")
		}
	})
	t.Run("Should calculate correctly when no special price appliable and quantity is > 1", func(t *testing.T) {
		pricingService.ClearCart()
		pricingService.Scan("A", 1)
		pricingService.Scan("A", 1)
		total, err := pricingService.Total()
		if err != nil || total != 100 {
			t.Errorf("Total should be 100")
		}
	})
	t.Run("Should calculate correctly when special price applicable and quantity > 1", func(t *testing.T) {
		pricingService.ClearCart()
		pricingService.Scan("A", 1)
		pricingService.Scan("A", 1)
		pricingService.Scan("A", 1)
		total, err := pricingService.Total()
		if err != nil || total != 130 {
			t.Errorf("Total should be 130")
		}
	})
	t.Run("Should calculate correctly when no special price appliable", func(t *testing.T) {
		pricingService.ClearCart()
		pricingService.Scan("C", 1)
		pricingService.Scan("D", 1)
		pricingService.Scan("B", 1)
		pricingService.Scan("A", 1)
		total, err := pricingService.Total()
		if err != nil || total != 115 {
			t.Errorf("Total should be 115")
		}
	})
	t.Run("Should calculate correctly when no special price appliable and multiple products added", func(t *testing.T) {
		pricingService.ClearCart()
		pricingService.Scan("C", 20)
		pricingService.Scan("D", 100)
		pricingService.Scan("B", 23)
		pricingService.Scan("A", 80)
		total, err := pricingService.Total()
		if err != nil || total != 5905 {
			t.Errorf("Total should be 5905")
		}
	})
}
