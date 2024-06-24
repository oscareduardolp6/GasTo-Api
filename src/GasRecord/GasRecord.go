package gasrecord

import (
	"time"

	"github.com/go-playground/validator/v10"
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

func CreateGasRecord(gasRecordPrimitives GasRecord) (GasRecord, error) {
	err := validate.Struct(gasRecordPrimitives)
	if err != nil {
		return GasRecord{}, err
	}
	return gasRecordPrimitives, nil
}
