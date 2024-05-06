package constant

type Category string

const (
	Clothing    Category = "Clothing"
	Accessories Category = "Accessories"
	Footwear    Category = "Footwear"
	Beverages   Category = "Beverages"
)

var ValidCategory = map[string]bool{
	string(Clothing):    true,
	string(Accessories): true,
	string(Footwear):    true,
	string(Beverages):   true,
}

var Categories = []string{
	string(Clothing),
	string(Accessories),
	string(Footwear),
	string(Beverages),
}
