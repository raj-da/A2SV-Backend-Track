package concurrency

import (
	"fmt"
	"task4/services"
)

type ReservationRequest struct {
	BookID int
	MemberID int
}

func ProcessReservations(
	lib services.LibraryManager, 
	requests <-chan ReservationRequest,
) {
	for req := range requests {
		err := lib.ReserveBook(req.BookID, req.MemberID)
		if err != nil {
			fmt.Printf("\n[Worker] Failed to reserve book %d: %v\n", req.BookID, err)
		} else {
			fmt.Printf("\n[Worker] Book %d reserved successfully for member %d (5s timer started)\n", req.BookID, req.MemberID)
		}
	}
}