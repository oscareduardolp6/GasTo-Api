package gasrecord_stub

import (
	domain "gasto-api/src/GasRecord"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func CreateBasicGasRecord() domain.GasRecordPrimitives {
	return domain.GasRecordPrimitives{
		Id:                 uuid.NewString(),
		Place:              createRandomPlace(),
		Liters:             createRandomLiters(),
		TotalPrice:         createRandomTotalPrice(),
		TraveledKilometers: createRandomTraveledKilometers(),
		PriceByLiter:       createRandomPriceByLiter(),
		Date:               createRandomDate(),
		RoadTrip:           false,
		Tank: domain.GasTankPrimitives{
			InitialLiters: createRandomLiters(),
			FinalLiters:   createRandomLiters(),
		},
	}
}

func createRandomDate() time.Time {
	start := time.Date(2024, 5, 10, 0, 0, 0, 0, time.UTC)
	end := time.Now()
	delta := end.Sub(start)
	seconds := rand.Int63n(int64(delta.Seconds()))
	return start.Add(time.Duration(seconds) * time.Second)
}

func createRandomTraveledKilometers() float32 {
	return float32(rand.Int31n(300))
}

func createRandomPriceByLiter() float32 {
	return float32(rand.Int31n(100))
}

func createRandomTotalPrice() float32 {
	return float32(rand.Int31n(1000))
}

func createRandomLiters() float32 {
	return float32(rand.Int31n(30))
}

func createRandomPlace() string {
	random := rand.Int31n(10)

	switch {
	case random <= 3:
		return "Gasolinería 1 Blvd. Aeropuerto"
	case random > 3 && random <= 6:
		return "Gasolinería 2 Torres Landa Oxxo Gas"
	case random > 6:
		return "Gasolinería 3 Gulf"
	}
	return "Gasolinería 4 Gulf camino a Silao"
}

// func createEmptyString() string {
// 	return ""
// }

// func createWhiteSpacesString(numWhiteSpaces int) string {
// 	return strings.Repeat(" ", numWhiteSpaces)
// }
