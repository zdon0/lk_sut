package main

import (
	"go.uber.org/fx"

	"lk_sut/internal/di"
)

// @title           Lk SUT Autocommitter

// @contact.name   	Maks Mikhaylov
// @contact.url		https://t.me/don101

func main() {
	fx.New(di.CreateApp()).Run()
}
