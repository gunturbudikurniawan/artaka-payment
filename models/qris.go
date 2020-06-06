package models

import (
	"github.com/jinzhu/gorm"
)

type Qris struct {
	gorm.Model
	Event      string
	QrID       string
	Amount     float64
	ExternalID string
	QrString   string
	QrType     string
	Status     string
}
