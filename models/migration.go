package models

import "github.com/naufalihsan/artaka-payment/adapter"

func AutoMigrate() {
	adapter.GetDB().AutoMigrate(&Qris{}, &Va{})
}
