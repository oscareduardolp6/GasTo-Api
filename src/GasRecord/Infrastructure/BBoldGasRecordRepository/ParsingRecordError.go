package gasrecord_infrastructure_bbold

import "fmt"

type parsingError struct {
	rawData []byte
}

func (myError parsingError) Error() string {
	return fmt.Sprintf("Error parsing the record from the db to json: %v", myError.rawData)
}

func newParsingRecordError(rawData []byte) parsingError {
	return parsingError{rawData}
}
