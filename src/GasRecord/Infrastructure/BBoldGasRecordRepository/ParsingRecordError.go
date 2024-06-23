package gasrecord_infrastructure_bbold

import "fmt"

type ParsingError struct {
	rawData []byte
}

func (myError ParsingError) Error() string {
	return fmt.Sprintf("Error parsing the record from the db to json: %v", myError.rawData)
}

func NewParsingRecordError(rawData []byte) ParsingError {
	return ParsingError{rawData}
}
