package gasrecord

import (
	"time"

	"github.com/IBM/fp-go/either"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type gasRecord struct {
	id                 string    `validate:"required,uuid4"`
	place              string    `validate:"required"`
	liters             float32   `validate:"gte=0"`
	totalPrice         float32   `validate:"gte=0"`
	traveledKilometers float32   `validate:"gte=0"`
	priceByLiter       float32   `validate:"gte=0"`
	date               time.Time `validate:"required"`
}

type GasRecord interface {
	GetId() string
	GetPlace() string
	GetLiters() float32
	GetTotalPrice() float32
	GetTraveledKilometers() float32
	GetPriceByLiter() float32
	GetDate() time.Time
}

func (gr gasRecord) GetId() string {
	return gr.id
}

func (gr gasRecord) GetPlace() string {
	return gr.place
}

func (gr gasRecord) GetLiters() float32 {
	return gr.liters
}

func (gr gasRecord) GetTotalPrice() float32 {
	return gr.totalPrice
}

func (gr gasRecord) GetTraveledKilometers() float32 {
	return gr.traveledKilometers
}

func (gr gasRecord) GetPriceByLiter() float32 {
	return gr.priceByLiter
}

func (gr gasRecord) GetDate() time.Time {
	return gr.date
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func CreateGasRecord(place string, liters, totalPrice, traveledKms, priceByLiter float32, date time.Time) either.Either[error, GasRecord] {
	gasRecord := gasRecord{
		id:                 uuid.NewString(),
		place:              place,
		liters:             liters,
		totalPrice:         totalPrice,
		traveledKilometers: traveledKms,
		priceByLiter:       priceByLiter,
		date:               date,
	}

	err := validate.Struct(gasRecord)
	if err != nil {
		return either.Left[GasRecord](err)
	}
	return either.Right[error, GasRecord](gasRecord)
}
