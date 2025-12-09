package student

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}
	items, exist := menu[order]
	if !exist {
		return 404
	}
	return items.preptime
}
