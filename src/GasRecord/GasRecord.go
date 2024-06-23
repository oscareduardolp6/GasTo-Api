package gasrecord

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

}

type GasRecord struct {
	Id                 string    `validate:"required,uuid4"`
	Place              string    `validate:"required"`
	Liters             float32   `validate:"gte=0"`
	TotalPrice         float32   `validate:"gte=0"`
	TraveledKilometers float32   `validate:"gte=0"`
	PriceByLiter       float32   `validate:"gte=0"`
	Date               time.Time `validate:"required"`
}

func CreateGasRecord(place string, liters, totalPrice, traveledKms, priceByLiter float32, date time.Time) (GasRecord, error) {
	gasRecord := GasRecord{
		Id:                 uuid.NewString(),
		Place:              place,
		Liters:             liters,
		TotalPrice:         totalPrice,
		TraveledKilometers: traveledKms,
		PriceByLiter:       priceByLiter,
		Date:               date,
	}
	err := validate.Struct(gasRecord)
	if err != nil {
		return GasRecord{}, err
	}
	return gasRecord, nil
}
