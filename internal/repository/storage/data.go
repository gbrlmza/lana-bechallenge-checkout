package storage

import "github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"

func (s *storage) initializeData() {
	s.data.baskets = make(map[string]entities.Basket, 0)
	s.data.products = make(map[string]entities.Product, 0)
	s.data.promotions = make(map[string]entities.Promotion, 0)

	//===========================================================================================
	// Promotions
	//===========================================================================================
	promotion2X1 := "BUY2GET1FREE"
	s.data.promotions[promotion2X1] = entities.Promotion{
		ID:            promotion2X1,
		RequiredItems: 2,
		FreeItems:     1,
		Reduction:     0,
	}
	promotion25Off := "BUY3+GET25OFF"
	s.data.promotions[promotion25Off] = entities.Promotion{
		ID:            promotion25Off,
		RequiredItems: 3,
		FreeItems:     0,
		Reduction:     25,
	}

	//===========================================================================================
	// Products
	//===========================================================================================
	s.data.products["PEN"] = entities.Product{
		ID:          "PEN",
		Name:        "Lana Pen",
		Price:       5.0,
		PromotionID: &promotion2X1,
	}
	s.data.products["TSHIRT"] = entities.Product{
		ID:          "TSHIRT",
		Name:        "Lana T-Shirt",
		Price:       20.0,
		PromotionID: &promotion25Off,
	}
	s.data.products["MUG"] = entities.Product{
		ID:          "MUG",
		Name:        "Lana Coffee Mug",
		Price:       7.50,
		PromotionID: nil,
	}
}
