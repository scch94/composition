package main

import (
	"fmt"

	"github.com/scch94/composition/pkg/customer"
	"github.com/scch94/composition/pkg/invoice"
	"github.com/scch94/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Pasto",
		customer.New("santiago canal", "calle 64#15-03", "3012415382"),
		[]invoiceitem.Item{
			invoiceitem.New(
				1,
				"mantequlla",
				12.32,
			),
			invoiceitem.New(
				2,
				"papa",
				11.11,
			),
		},
	)
	fmt.Printf("%+v\n", i)
}
