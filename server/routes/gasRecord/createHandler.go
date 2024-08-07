package routes

import (
	"encoding/json"
	domain "gasto-api/src/GasRecord"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func makeGasRecordPutHandler(createGasRecord func(domain.GasRecordPrimitives) error) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, req *http.Request) {
		body, bodyError := io.ReadAll(req.Body)
		if bodyError != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
			responseWriter.Write(nil)
			return
		}
		defer req.Body.Close()

		var primitives domain.GasRecordPrimitives
		parsingError := json.Unmarshal(body, &primitives)
		if parsingError != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
			responseWriter.Write([]byte(parsingError.Error()))
			return
		}
		validationError := domain.ValidatePrimitives(primitives)
		if validationError != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
			responseWriter.Write([]byte(validationError.(validator.ValidationErrors).Error()))
			return
		}
		creationError := createGasRecord(primitives)
		if creationError != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			responseWriter.Write(nil)
			log.Print(creationError.Error())
			return
		}
		responseWriter.WriteHeader(http.StatusAccepted)
		responseWriter.Write(nil)
	}
}
