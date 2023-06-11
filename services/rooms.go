package services

import "log"

type RoomTypes string

const (
	Premium RoomTypes = "Premium"
	Gold              = "Gold"
	VIP               = "VIP"
	Commom            = "Common"
)

type RoomsService interface {
}

type Rooms struct {
	Number   int8
	Ocupied  bool
	RoomType RoomTypes
}

func NewRoom(number int8, roomtype RoomTypes) *Rooms {
	return &Rooms{
		Number:   number,
		Ocupied:  false,
		RoomType: roomtype,
	}
}

func (r *Rooms) SetRoomToOccupied() error {
	r.Ocupied = true
	log.Printf("Room %d occupied!", r.Number)
	return nil
}
