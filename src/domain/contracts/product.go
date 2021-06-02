package contracts

type Product struct {
	Id       string `json:"id"`
	Title	 string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Discount float64 `json:"discount"`
	ShippingCost float64 `json:"shipping_cost"`
	ImageUrl string `json:"image_url"`
	StoreId string `json:"store_id"`
}
