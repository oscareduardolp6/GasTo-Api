package gasrecord

import (
	share "gasto-api/src/Share"
	"math"
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
	RoadTrip           bool
	Performance        float32 //Calculated
	domainEvents       []share.Event
}

func calculatePerformance(previous GasRecord, next GasRecord) float32 {
	rawPerformance := next.TraveledKilometers / previous.Liters
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

func CreateGasRecord(gasRecordPrimitives GasRecord) (*GasRecord, error) {
	err := validate.Struct(gasRecordPrimitives)
	if err != nil {
		return nil, err
	}

	newGasRecord := gasRecordPrimitives
	newGasRecord.domainEvents = []share.Event{CreateGasRecordCreatedEvent(gasRecordPrimitives)}
	return &newGasRecord, nil
}
