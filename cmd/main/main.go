package main

import (
	"fmt"
	"lk_sut/internal/app"
)

// @title           Lk SUT Autocommitter

// @contact.name   	Maks Mikhaylov
// @contact.url		https://t.me/don101

func main() {
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
