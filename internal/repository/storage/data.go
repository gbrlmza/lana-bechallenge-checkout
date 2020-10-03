package storage

import "github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"

func (s *storage) initializeData() {
	s.data.baskets = make(map[string]entities.Basket, 0)
	s.data.products = make(map[string]entities.Product, 0)
	s.data.discounts = make(map[string]entities.Discount, 0)

	//===========================================================================================
	// Discounts
	//===========================================================================================
	discount2X1 := "BUY2GET1FREE"
	s.data.discounts[discount2X1] = entities.Discount{
		Code:          discount2X1,
		Type:          "BuyXGetN",
		RequiredItems: 2,
		FreeItems:     1,
		Reduction:     0,
	}
	discount25Off := "BUY3GET25OFF"
	s.data.discounts[discount25Off] = entities.Discount{
		Code:          discount25Off,
		Type:          "Bulk",
		RequiredItems: 3,
		FreeItems:     0,
		Reduction:     25,
	}

	//===========================================================================================
	// Products
	//===========================================================================================
	s.data.products["PEN"] = entities.Product{
		ID:         "PEN",
		Name:       "Lana Pen",
		Price:      5.0,
		DiscountID: &discount2X1,
	}
	s.data.products["TSHIRT"] = entities.Product{
		ID:         "TSHIRT",
		Name:       "Lana T-Shirt",
		Price:      20.0,
		DiscountID: &discount25Off,
	}
	s.data.products["MUG"] = entities.Product{
		ID:         "MUG",
		Name:       "Lana Coffee Mug",
		Price:      7.50,
		DiscountID: nil,
	}
}
