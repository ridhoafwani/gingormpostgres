package response

import (
	"time"

	"github.com/ridhoafwani/gingormpostgres/models"
)

type Order struct {
	OrderId      int
	CustomerName string
	OrderedAt    time.Time
	Items        []models.Item
}
