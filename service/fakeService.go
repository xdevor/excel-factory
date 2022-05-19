package service

import (
	"fmt"
	"strconv"

	gofakeit "github.com/brianvoe/gofakeit/v6"
)

func GetFakeData(dataType string) string {
	switch dataType {
	case "FirstName":
		return gofakeit.FirstName()
	case "LastName":
		return gofakeit.LastName()
	case "Name":
		return gofakeit.Name()
	case "Email":
		return gofakeit.Email()
	case "Phone":
		return gofakeit.Phone()
	case "Gender":
		return gofakeit.Gender()
	case "SSN":
		return gofakeit.SSN()
	case "Hobby":
		return gofakeit.Hobby()
	case "Username":
		return gofakeit.Username()
	case "Password":
		return gofakeit.Password(true, true, true, false, false, 9)
	case "Country":
		return gofakeit.Country()
	case "City":
		return gofakeit.City()
	case "State":
		return gofakeit.State()
	case "Street":
		return gofakeit.Street()
	case "Zip":
		return gofakeit.Zip()
	case "Latitude":
		return fmt.Sprintf("%f", gofakeit.Latitude())
	case "Longitude":
		return fmt.Sprintf("%f", gofakeit.Longitude())
	case "BeerName":
		return gofakeit.BeerName()
	case "Noun":
		return gofakeit.Noun()
	case "Verb":
		return gofakeit.Verb()
	case "Adjective":
		return gofakeit.Adjective()
	case "Adverb":
		return gofakeit.Adverb()
	case "Word":
		return gofakeit.Word()
	case "Sentence":
		return gofakeit.Sentence(10)
	case "Fruit":
		return gofakeit.Fruit()
	case "Vegetable":
		return gofakeit.Vegetable()
	case "Breakfast":
		return gofakeit.Breakfast()
	case "Lunch":
		return gofakeit.Lunch()
	case "Dinner":
		return gofakeit.Dinner()
	case "Snack":
		return gofakeit.Snack()
	case "Dessert":
		return gofakeit.Dessert()
	case "Bool":
		return strconv.FormatBool(gofakeit.Bool())
	case "UUID":
		return gofakeit.UUID()
	case "URL":
		return gofakeit.URL()
	case "DomainName":
		return gofakeit.DomainName()
	case "IPv4Address":
		return gofakeit.IPv4Address()
	case "IPv6Address":
		return gofakeit.IPv6Address()
	case "Color":
		return gofakeit.Color()
	case "HexColor":
		return gofakeit.HexColor()
	case "Number":
		return strconv.Itoa(gofakeit.IntRange(0, 999999))
	default:
		return gofakeit.Word()
	}
}
