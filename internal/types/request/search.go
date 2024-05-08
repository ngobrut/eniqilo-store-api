package request

type SearchQuery struct {
	Limit    *int
	Offset   *int
	Name     *string
	Category *string
	Sku      *string
	InStock  *bool
	Price    *string
}
