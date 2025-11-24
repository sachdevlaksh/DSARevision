package LC2353

import (
	"cmp"
	"fmt"
	"slices"
)

type FoodRatings struct {
	foodrating map[string]Food
}

type Food struct {
	FoodName string
	Rating   int
	Cuisine  string
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	iniFood := make(map[string]Food, 0)
	for i := 0; i < len(cuisines); i++ {
		currFood := Food{
			FoodName: foods[i],
			Rating:   ratings[i],
			Cuisine:  cuisines[i],
		}
		iniFood[foods[i]] = currFood

	}

	for k, v := range iniFood {
		slice := v
		slices.SortFunc(slice.foodDetail, func(a, b FoodDetail) int {
			if n1 := cmp.Compare(b.Rating, a.Rating); n1 != 0 {
				return n1
			}
			return cmp.Compare(a.FoodName, b.FoodName)
		})
		iniFood[k] = slice
	}

	fmt.Println("Inside Constructor")
	fmt.Println(iniFood)
	return FoodRatings{iniFood}
}
func (this *FoodRatings) ChangeRating(food string, newRating int) {

	fd := this.foodrating[food]
	for i, v := range fd {

		slices.SortFunc(this.foodrating[cuisine].foodDetail, func(a, b FoodDetail) int {
			if n1 := cmp.Compare(b.Rating, a.Rating); n1 != 0 {
				return n1
			}
			return cmp.Compare(a.FoodName, b.FoodName)
		})

	}
	fmt.Println("Inside Change Rating")
	fmt.Println(cuisine)
	fmt.Println(this.foodrating[cuisine])

}

func (this *FoodRatings) HighestRated(cuisine string) string {
	fmt.Println("Inside Get Rating")
	fmt.Println(this.foodrating)

	return this.foodrating[cuisine].foodDetail[0].FoodName
}
