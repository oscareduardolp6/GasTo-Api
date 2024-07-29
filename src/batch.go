package main

// import (
// 	. "gasto-api/src/GasRecord"
// )

// type Mapper func(row string) GasRecord

// const (
// 	place_part     = 0
// 	liters_part    = 1
// 	start_km       = 2
// 	price_by_liter = 3
// 	total_price    = 4
// 	date_part      = 5
// 	roadtrip       = 7
// )

// func csvMapper(row string) *GasRecord {
// 	parts := strings.Split(row, ",")

// 	if len(parts) < 7 {
// 		log.Fatal("Invalid number of parts in the row", row)
// 	}
// 	place := parts[place_part]
// 	liters := parts[liters_part]
// 	traveled_km := parts[start_km]
// 	price_per_liter := parts[price_by_liter]
// 	total := parts[total_price]
// 	date := parts[date_part]
// 	has_roadtrip := parts[roadtrip]

// 	parsedLiters, parsingLitersError := strconv.ParseFloat(liters, 32)
// 	parsedTraveledKm, parsingTraveledKmError := strconv.ParseFloat(traveled_km, 32)
// 	parsedTotal, parsingTotalError := strconv.ParseFloat(total, 32)
// 	parsedPriceByLiter, parsingPriceByLiterError := strconv.ParseFloat(price_per_liter, 32)
// 	parsedDate, parsingDateError := parseDate(date)
// 	parsedRoadTrip, parsingRoadTripError := strconv.ParseBool(has_roadtrip)

// 	log.Printf("ParsedRoadTrip: %v \n", parsedRoadTrip)

// 	if parsingRoadTripError != nil {
// 		log.Fatal("RoadTrip parsing failed", row)
// 	}

// 	if parsingDateError != nil {
// 		log.Fatal("Date parsing Failed", row)
// 	}

// 	if parsingLitersError != nil {
// 		log.Fatal("Liters parsing failed", row)
// 	}

// 	if parsingTraveledKmError != nil {
// 		log.Fatal("Traveled Km parsing failed", row)
// 	}

// 	if parsingTotalError != nil {
// 		log.Fatal("Total price parsing failed", row)
// 	}

// 	if parsingPriceByLiterError != nil {
// 		log.Fatal("Price by Liter parsing failed", row)
// 	}

// 	primitives := GasRecord{
// 		Id:                 uuid.NewString(),
// 		Place:              place,
// 		Liters:             float32(parsedLiters),
// 		TraveledKilometers: float32(parsedTraveledKm),
// 		TotalPrice:         float32(parsedTotal),
// 		PriceByLiter:       float32(parsedPriceByLiter),
// 		Date:               parsedDate,
// 		RoadTrip:           parsedRoadTrip,
// 	}

// 	gasRecord, creationError := CreateGasRecord(primitives)

// 	if creationError != nil {
// 		log.Fatal("Error while creating the Gas Record with the next primitives", primitives)
// 	}

// 	return gasRecord
// }

// func parseDate(strDate string) (time.Time, error) {
// 	return time.Parse("1/2/2006", strDate)
// }

// func convertFileToGasRecords(filename string) {

// }

// func executeBatchFromFile(filename string, mapper Mapper, repo GasRecordRepository) {
// 	file, fileError := os.Open(filename)
// 	if fileError != nil {
// 		log.Fatal(fileError)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)

// 	createGasRecord := app.MakeCreateGasRecord(repo, nil)

// 	insertions := 0
// 	isHeader := true

// 	for scanner.Scan() {
// 		if isHeader {
// 			isHeader = false
// 			continue
// 		}

// 		line := scanner.Text()
// 		rawData := mapper(line)
// 		creationError := createGasRecord(rawData)
// 		if creationError != nil {
// 			log.Printf("Correct Insertions: %d", insertions)
// 			log.Fatal(creationError)
// 		}
// 		insertions++
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Printf("All record succesfull inserted")
// }
