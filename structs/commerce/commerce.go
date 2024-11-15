package commerce

type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tag"`
}

func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
}

type Cart struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products,omitempty"`
}

func (c *Cart) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Cart) Total() float64 { // como no voy a modificar nada del carrito original no le agrego * a Cart
	var total float64
	for _, producto := range c.Products {
		total += producto.TotalPrice()
	}
	return total
}

func NewCart(userID uint) Cart {
	return Cart{UserID: userID}
}
