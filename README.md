### Features

* **Easy to Use**. This library is designed to be easy to use. You can start using it in just a few minutes.
* **Well maintained**. This library is actively maintained by stakeholders.
* **Full Coverage**. All Kaspi.kz Merchant API methods are implemented.
* **No External Dependencies**. This library is written in pure Go, so it has no external dependencies.

### Installation

`go get -u github.com/abdymazhit/kaspi-merchant-api`

### Example

```go
package main

import (
	"context"
	"fmt"
	kma "github.com/abdymazhit/kaspi-merchant-api"
	"time"
)

func main() {
	api := kma.New("AUTH-TOKEN")

	orders, err := api.GetOrders(context.Background(), kma.GetOrdersRequest{
		PageNumber: 0,
		PageSize:   10,
		Filter: struct {
			Orders struct {
				CreationDateGe    time.Time
				CreationDateLe    time.Time
				State             kma.OrdersState
				Status            kma.OrdersStatus
				DeliveryType      kma.OrdersDeliveryType
				SignatureRequired bool
			}
		}{
			Orders: struct {
				CreationDateGe    time.Time
				CreationDateLe    time.Time
				State             kma.OrdersState
				Status            kma.OrdersStatus
				DeliveryType      kma.OrdersDeliveryType
				SignatureRequired bool
			}{
				CreationDateGe:    time.Now().AddDate(0, 0, -1),
				CreationDateLe:    time.Now(),
				State:             kma.OrdersStateNew,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", orders)
}
```

### Legal

This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by **Kaspi.kz** or any of its affiliates or subsidiaries.