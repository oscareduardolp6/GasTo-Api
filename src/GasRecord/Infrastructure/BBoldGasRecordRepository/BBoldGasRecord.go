package gasrecord_infrastructure_bbold

import (
	. "gasto-api/src/GasRecord"
	"time"
)

type bboldGasRecord struct {
	ID                 string    `json:"id"`
	Place              string    `json:"place"`
	Liters             float32   `json:"liters"`
	TraveledKilometers float32   `json:"traveled_kilometers"`
	PriceByLiter       float32   `json:"price_by_liter"`
	TotalPrice         float32   `json:"total_price"`
	Date               time.Time `json:"date"`
	RoadTrip           bool      `json:"road_trip"`
	Performance        float32   `json:"performance"`
}

func toBBoldGasRecord(record GasRecord) bboldGasRecord {
	return bboldGasRecord{
		ID:                 record.Id,
		Place:              record.Place,
		Liters:             record.Liters,
		TraveledKilometers: record.TraveledKilometers,
		PriceByLiter:       record.PriceByLiter,
		TotalPrice:         record.TotalPrice,
		Date:               record.Date,
		RoadTrip:           record.RoadTrip,
		Performance:        record.Performance,
	}
}

func fromBBoldGasRecord(dto bboldGasRecord) GasRecord {
	return GasRecord{
		Id:                 dto.ID,
		Place:              dto.Place,
		Liters:             dto.Liters,
		TraveledKilometers: dto.TraveledKilometers,
		PriceByLiter:       dto.PriceByLiter,
		TotalPrice:         dto.TotalPrice,
		Date:               dto.Date,
		RoadTrip:           dto.RoadTrip,
		Performance:        dto.Performance,
	}
}
