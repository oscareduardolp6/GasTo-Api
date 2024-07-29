package gasrecord

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidatePrimitives(gasRecord GasRecordPrimitives) error {
	return validate.Struct(gasRecord)
}

type GasTankPrimitives struct {
	InitialLiters float32 `json:"initial_liters" validate:"required,gte=0"`
	FinalLiters   float32 `json:"final_liters" validate:"required,gte=0"`
}

type GasRecordPrimitives struct {
	Id                 string            `json:"id" validate:"required,uuid4"`
	Place              string            `json:"place" validate:"required"`
	Liters             float32           `json:"liters" validate:"gte=0"`
	TotalPrice         float32           `json:"total_price" validate:"gte=0"`
	TraveledKilometers float32           `json:"traveled_kilometers" jsonvalidate:"gte=0"`
	PriceByLiter       float32           `json:"price_by_liter" validate:"gte=0"`
	Date               time.Time         `json:"date" validate:"required"`
	RoadTrip           bool              `json:"road_trip"`
	Tank               GasTankPrimitives `json:"tank"`
}
