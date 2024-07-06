package routes

import (
	"encoding/json"
	domain "gasto-api/src/GasRecord"
	"net/http"
)

func makeGetAllGetHandler(getAll func() []domain.GasRecord) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		jsonData, parsingError := json.Marshal(getAll())
		if parsingError != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write(nil)
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write(jsonData)
	}
}
