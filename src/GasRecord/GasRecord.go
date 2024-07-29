package gasrecord

import (
	share "gasto-api/src/Share"
	"log"
	"time"
)

type Tank struct {
	Initial, Final Liters
}

type GasRecord struct {
	Id           GasRecordId
	Place        share.NonEmptyString
	Liters       Liters
	Total        GasRecordPrice
	Traveled     Kilometers
	PriceByLiter GasRecordPrice
	Date         time.Time
	RoadTrip     bool
	Performance  share.PositiveNumber
	Tank         Tank
	domainEvents []share.Event
}

func calculatePerformance(previous GasRecord, next GasRecord) share.PositiveNumber {
	utilizedLiters, litersError := previous.Tank.Final.Substract(next.Tank.Initial)
	if litersError != nil {
		log.Fatal(litersError)
	}

	rawPerformance := next.Traveled.Value() / utilizedLiters.Value()
	result, err := share.CreatePositiveNumber(rawPerformance)
	if err != nil {
		log.Fatal(err)
	}
	return result
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

func CreateGasRecord(gasRecordPrimitives GasRecordPrimitives) (*GasRecord, *GasRecordError) {

	id, idError := CreateGasRecordId(gasRecordPrimitives.Id)
	place, placeError := share.CreateNonEmptyString(gasRecordPrimitives.Place)
	liters, litersError := CreateLiters(gasRecordPrimitives.Liters)
	total, totalError := CreateGasRecordPrice(gasRecordPrimitives.TotalPrice)
	traveled, traveledError := CreateKilometers(gasRecordPrimitives.TraveledKilometers)
	priceByLiter, priceByLiterError := CreateGasRecordPrice(gasRecordPrimitives.PriceByLiter)
	initialLiters, initialLitersError := CreateLiters(gasRecordPrimitives.Tank.InitialLiters)
	finalLiters, finalLitersError := CreateLiters(gasRecordPrimitives.Tank.FinalLiters)

	if share.Any(share.IsNil, idError, placeError, litersError, totalError, traveledError, priceByLiterError, initialLitersError, finalLitersError) {
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
		domainEvents: []share.Event{CreateGasRecordCreatedEvent(gasRecordPrimitives)},
	}

	return &gasRecord, nil
}
