package gasrecord

import (
	"log"
)

type TankError struct {
	Initial, Final error
}

type GasRecordError struct {
	Id, Place, Liters, Total, Traveled, PriceByLiter error
	Tank                                             TankError
}

func (err GasRecordError) Error() string {
	log.Printf("Id: <%v>", err.Id)
	log.Printf("Place: <%v>", err.Place)
	log.Printf("Liters: <%v>", err.Liters)
	log.Printf("Total: <%v>", err.Total)
	log.Printf("Traveled: <%v>", err.Traveled)
	log.Printf("PriceByLiter: <%v>", err.PriceByLiter)
	log.Printf("Initial Liters: <%v>", err.Tank.Initial)
	log.Printf("Final Liters: <%v>", err.Tank.Final)

	return "Error Creating Gas Record "
}
