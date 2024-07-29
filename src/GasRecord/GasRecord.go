package gasrecord

import (
	shared "gasto-api/src/Shared"
	"log"
	"time"
)

type Tank struct {
	Initial, Final Liters
}

type GasRecord struct {
	Id           GasRecordId
	Place        shared.NonEmptyString
	Liters       Liters
	Total        GasRecordPrice
	Traveled     Kilometers
	PriceByLiter GasRecordPrice
	Date         time.Time
	RoadTrip     bool
	Performance  shared.PositiveNumber
	Tank         Tank
	domainEvents []shared.Event
}

func calculatePerformance(previous GasRecord, next GasRecord) shared.PositiveNumber {
	utilizedLiters, litersError := previous.Tank.Final.Substract(next.Tank.Initial)
	if litersError != nil {
		log.Fatal(litersError)
	}

	rawPerformance := next.Traveled.Value() / utilizedLiters.Value()
	result, err := shared.CreatePositiveNumber(rawPerformance)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (gasRecord *GasRecord) UpadatePerformance(next GasRecord) {
	gasRecord.Performance = calculatePerformance(*gasRecord, next)
	gasRecord.domainEvents = append(gasRecord.domainEvents, CreateGasRecordPerformanceUpdated(*gasRecord))
}

func (gasRecord *GasRecord) PullAllDomainEvents() []shared.Event {
	numOfEvents := len(gasRecord.domainEvents)
	returnedEvents := make([]shared.Event, numOfEvents)
	copy(returnedEvents, gasRecord.domainEvents)
	gasRecord.domainEvents = []shared.Event{}
	return returnedEvents
}

func CreateGasRecord(gasRecordPrimitives GasRecordPrimitives) (*GasRecord, *GasRecordError) {

	id, idError := CreateGasRecordId(gasRecordPrimitives.Id)
	place, placeError := shared.CreateNonEmptyString(gasRecordPrimitives.Place)
	liters, litersError := CreateLiters(gasRecordPrimitives.Liters)
	total, totalError := CreateGasRecordPrice(gasRecordPrimitives.TotalPrice)
	traveled, traveledError := CreateKilometers(gasRecordPrimitives.TraveledKilometers)
	priceByLiter, priceByLiterError := CreateGasRecordPrice(gasRecordPrimitives.PriceByLiter)
	initialLiters, initialLitersError := CreateLiters(gasRecordPrimitives.Tank.InitialLiters)
	finalLiters, finalLitersError := CreateLiters(gasRecordPrimitives.Tank.FinalLiters)

	if shared.Any(shared.IsNil, idError, placeError, litersError, totalError, traveledError, priceByLiterError, initialLitersError, finalLitersError) {
		return nil, &GasRecordError{
			Id:           idError,
			Place:        placeError,
			Liters:       litersError,
			Total:        totalError,
			Traveled:     traveledError,
			PriceByLiter: priceByLiterError,
			Tank: TankError{
				Initial: initialLitersError,
				Final:   finalLitersError,
			},
		}
	}

	gasRecord := GasRecord{
		Id:           id,
		Place:        place,
		Liters:       liters,
		Total:        total,
		Traveled:     traveled,
		PriceByLiter: priceByLiter,
		Date:         gasRecordPrimitives.Date,
		RoadTrip:     gasRecordPrimitives.RoadTrip,
		Tank: Tank{
			Initial: initialLiters,
			Final:   finalLiters,
		},
		domainEvents: []shared.Event{CreateGasRecordCreatedEvent(gasRecordPrimitives)},
	}

	return &gasRecord, nil
}
