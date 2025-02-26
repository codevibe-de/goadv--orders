package model

type Order struct {
	customerPhoneNumber string
	quantities          map[string]int
}
