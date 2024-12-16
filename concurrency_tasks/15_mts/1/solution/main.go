package main

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type OrderService struct {
	BookingService BookingService
	UserService    UserService
}

type UserService interface {
	LockUser(User) error
	UnlockUser(User) error
}

type User struct {
	ID string
}

type Receipt struct {
	ID          string
	BookingCode string
	BookedAt    string
}

type BookingService interface {
	BookFlight() (string, *BookingServiceError)
}

type BookingServiceError struct {
	error
	TryAgain bool
}

func (s *OrderService) HandleBookingOrder(user User) *Receipt {
	receipt := Receipt{ID: uuid.New().String()}

	if err := s.UserService.LockUser(user); err != nil {
		log.Logger.Err(err)
		return nil
	}
LOOP:
	for {
		bookingCode, err := s.BookingService.BookFlight()

		switch {
		case err == nil:
			receipt.BookedAt = time.Now().Format(time.RFC3339)
			receipt.BookingCode = bookingCode
			return &receipt
		case err.TryAgain:
			// continue
			// нет обработки TryAgain
		default:
			log.Logger.Err(err)
			break LOOP
		}
	}

	if err := s.UserService.UnlockUser(user); err != nil {
		log.Logger.Err(err)
		return nil
	}

	return &receipt
}

// 1. Метод BookFlight() не имеет вводных параметров.
// 2. Метод HandleBookingOrder не возвращает ошибку на уровень выше.
// 3. Добавить для OrderService возврат его специфической ошибки.
