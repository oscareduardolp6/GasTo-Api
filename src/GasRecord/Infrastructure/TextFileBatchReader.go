package infrastructure

import (
	"bufio"
	. "gasto-api/src/GasRecord"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func ReadRecordsFromTextFile(filePath string) (<-chan GasRecord, error) {
	file, openFileError := os.Open(filePath)
	if openFileError != nil {
		return nil, openFileError
	}
	recordsChan := make(chan GasRecord)

	go func() {
		defer close(recordsChan)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		isFirstLine := true
		for scanner.Scan() {
			if isFirstLine {
				isFirstLine = false
				continue
			}
			line := scanner.Text()
			record := csvMapper(line)
			recordsChan <- record
		}
	}()
	return recordsChan, nil
}

const (
	place_part     = 0
	liters_part    = 1
	start_km       = 2
	price_by_liter = 3
	total_price    = 4
	date_part      = 5
	roadtrip       = 7
)

func csvMapper(row string) GasRecord {
	parts := strings.Split(row, ",")

	if len(parts) < 7 {
		log.Fatal("Invalid number of parts in the row", row)
	}
	place := parts[place_part]
	liters := parts[liters_part]
	traveled_km := parts[start_km]
	price_per_liter := parts[price_by_liter]
	total := parts[total_price]
	date := parts[date_part]
	has_roadtrip := parts[roadtrip]

	parsedLiters, parsingLitersError := strconv.ParseFloat(liters, 32)
	parsedTraveledKm, parsingTraveledKmError := strconv.ParseFloat(traveled_km, 32)
	parsedTotal, parsingTotalError := strconv.ParseFloat(total, 32)
	parsedPriceByLiter, parsingPriceByLiterError := strconv.ParseFloat(price_per_liter, 32)
	parsedDate, parsingDateError := parseDate(date)
	parsedRoadTrip, parsingRoadTripError := strconv.ParseBool(has_roadtrip)

	log.Printf("ParsedRoadTrip: %v \n", parsedRoadTrip)

	if parsingRoadTripError != nil {
		log.Fatal("RoadTrip parsing failed", row)
	}

	if parsingDateError != nil {
		log.Fatal("Date parsing Failed", row)
	}

	if parsingLitersError != nil {
		log.Fatal("Liters parsing failed", row)
	}

	if parsingTraveledKmError != nil {
		log.Fatal("Traveled Km parsing failed", row)
	}

	if parsingTotalError != nil {
		log.Fatal("Total price parsing failed", row)
	}

	if parsingPriceByLiterError != nil {
		log.Fatal("Price by Liter parsing failed", row)
	}

	primitives := GasRecord{
		Id:                 uuid.NewString(),
		Place:              place,
		Liters:             float32(parsedLiters),
		TraveledKilometers: float32(parsedTraveledKm),
		TotalPrice:         float32(parsedTotal),
		PriceByLiter:       float32(parsedPriceByLiter),
		Date:               parsedDate,
		RoadTrip:           parsedRoadTrip,
	}

	return primitives
}

func parseDate(strDate string) (time.Time, error) {
	return time.Parse("1/2/2006", strDate)
}
