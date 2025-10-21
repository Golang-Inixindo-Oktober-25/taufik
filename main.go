package main

import (
	"fmt"

	"github.com/Golang-Inixindo-Oktober-25/taufik/helper" // ubah sesuai path project kamu
)

func main() {
	helper.Info("Testing fungsi email validator")

	emails := []string{
		"user@example.com",
		"invalid@",
		"noatsign.com",
		"taufik.typ@bankjatim.co.id",
	}

	for _, email := range emails {
		if helper.IsValidEmail(email) {
			fmt.Printf("%s -> %v\n", email, true)
		} else {
			helper.Error(fmt.Sprintf("Email tidak valid: %s", email))
		}
	}
}
