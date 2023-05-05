package main

import (
	"fmt"
	"go-email/internal/domain/campaign"

	"github.com/go-playground/validator"
)

func main() {
	validate := validator.New()

	contacts := []campaign.Contact{{Email: "teste@gmail.com"}}
	campaign := campaign.Campaign{Contacts: contacts}

	errs := validate.Struct(campaign)
	if errs == nil {
		fmt.Println("No error")
	} else {
		errs := errs.(validator.ValidationErrors)
		for _, err := range errs {
			switch err.Tag() {
			case "required":
				fmt.Println(err.StructField() + " is required")
			case "min":
				fmt.Println(err.StructField() + " is required with min " + err.Param())
			case "max":
				fmt.Println(err.StructField() + " is required with max " + err.Param())
			case "email":
				fmt.Println(err.StructField() + " is invalid")
			}
		}
	}

}
