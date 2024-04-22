package invoice

import (
	"github.com/scch94/composition/pkg/customer"
	"github.com/scch94/composition/pkg/invoiceitem"
)

// una factura tiene muchos items
// una factura es para un solo cliente
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

// new retorna una nueva factura
func New(country string, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	invoice := Invoice{
		country: country,
		city:    city,
		total:   0.0,
		client:  client,
		items:   items,
	}
	invoice.setTotal()
	return invoice
}

func (i *Invoice) setTotal() {
	i.total = i.items.Total()
}
