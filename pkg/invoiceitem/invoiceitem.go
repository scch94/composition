package invoiceitem

type Item struct {
	id      uint
	product string
	value   float64
}

func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

func (i *Item) SetId(id uint) {
	i.id = id
}

func (i *Item) Id() uint {
	return i.id
}

func (i *Item) SetProduct(product string) {
	i.product = product
}

func (i *Item) Product() string {
	return i.product
}

func (i *Item) SetValue(value float64) {
	i.value = value
}

func (i *Item) Value() float64 {
	return i.value
}
