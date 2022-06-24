package csv

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"time"
)

func Writer() *os.File {

	var orders []*Order

	orders = append(orders, &Order{
		OrderId:        15252195855,
		PlaceOrderTime: time.Now(),
		PayableAmount:  100,
		Currency:       "人民币",
	})

	//dir, _ := os.MkdirTemp("", "*orders")
	//fmt.Println(dir)

	f, err := os.CreateTemp("/var/folders/mq/y1gb7qh53rl32lr089tbdrlm0000gn/T/2324137876orders", "*.csv")
	if err != nil {
		fmt.Println("create file err", err)
	}
	fmt.Println(f.Name())

	err = gocsv.MarshalFile(&orders, f)
	if err != nil {
		fmt.Println("marshal csv file err", err)
	}

	return f
}
