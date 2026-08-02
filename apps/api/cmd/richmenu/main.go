package main

import (
	"github.com/car-rental-template/api/internal/services"
	"log"
	"os"
	"path/filepath"
)

func main() {
	token, liffID := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"), os.Getenv("LIFF_ID")
	if token == "" || liffID == "" {
		log.Fatal("LINE_CHANNEL_ACCESS_TOKEN and LIFF_ID are required")
	}
	service := services.NewRichMenuService(token)
	url, root := "https://liff.line.me/"+liffID, "assets/richmenu"
	rentalID, err := service.CreateMenu("Rental start menu", "เริ่มเช่ารถ", filepath.Join(root, "rental-menu.png"), url)
	if err != nil {
		log.Fatal(err)
	}
	bookingID, err := service.CreateMenu("Booking progress menu", "การเช่ารถของคุณ", filepath.Join(root, "booking-menu.png"), url+"?screen=booking")
	if err != nil {
		log.Fatal(err)
	}
	if err := service.SetDefault(rentalID); err != nil {
		log.Fatal(err)
	}
	log.Printf("Created rental menu: %s\nCreated booking menu: %s\nSet LINE_BOOKING_RICH_MENU_ID=%s", rentalID, bookingID, bookingID)
}
