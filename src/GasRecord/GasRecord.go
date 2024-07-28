package gasrecord

import (
	share "gasto-api/src/Share"
	"log"
	"math"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

}

type GasRecord struct {
	Id                 string    `json:"id" validate:"required,uuid4"`
	Place              string    `json:"place" validate:"required"`
	Liters             float32   `json:"liters" validate:"gte=0"`
	TotalPrice         float32   `json:"total_price" validate:"gte=0"`
	TraveledKilometers float32   `json:"traveled_kilometers" jsonvalidate:"gte=0"`
	PriceByLiter       float32   `json:"price_by_liter" validate:"gte=0"`
	Date               time.Time `json:"date" validate:"required"`
	RoadTrip           bool      `json:"road_trip"`
	Performance        float32   `json:"performance"`
	Tank               GasTank   `json:"tank"`
	domainEvents       []share.Event
}

func calculatePerformance(previous GasRecord, next GasRecord) float32 {
	utilizedLiters, litersError := previous.Tank.Final.Substract(next.Tank.Initial)
	if litersError != nil {
		log.Fatal(litersError)
	}
	rawPerformance := next.TraveledKilometers / utilizedLiters.Value
	return float32(math.Ceil(float64(rawPerformance)))
}

func (gasRecord *GasRecord) UpadatePerformance(next GasRecord) {
	gasRecord.Performance = calculatePerformance(*gasRecord, next)
	gasRecord.domainEvents = append(gasRecord.domainEvents, CreateGasRecordPerformanceUpdated(*gasRecord))
}

func (gasRecord *GasRecord) PullAllDomainEvents() []share.Event {
	numOfEvents := len(gasRecord.domainEvents)
	returnedEvents := make([]share.Event, numOfEvents)
	copy(returnedEvents, gasRecord.domainEvents)
	gasRecord.domainEvents = []share.Event{}
	return returnedEvents
}

func ValidatePrimitives(gasRecord GasRecord) error {
	return validate.Struct(gasRecord)
}

func CreateGasRecord(gasRecordPrimitives GasRecord) (*GasRecord, error) {
	validationError := ValidatePrimitives(gasRecordPrimitives)
	if validationError != nil {
		return nil, validationError
	}
	newGasRecord := gasRecordPrimitives
	newGasRecord.domainEvents = []share.Event{CreateGasRecordCreatedEvent(gasRecordPrimitives)}
	return &newGasRecord, nil
}
