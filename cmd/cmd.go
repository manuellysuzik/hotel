package cmd

import (
	"fmt"
	"github.com/manuellysuzik/hotel.git/services"
)

func BootstrapGuestApi() {
	fmt.Println("iniciou guest api")
}
func BootstrapRoomsApi() {
	room := services.NewRoom(1, "Premium")
	isOccupied := room.SetRoomToOccupied()
	if isOccupied != nil {
		fmt.Println("Was not possible to occupy this room , try again later")
	}
	fmt.Println(fmt.Sprintf("Room %d occupied successfully", room.Number))

}
