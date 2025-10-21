package helper

import (
	"fmt"
	"regexp"
)

// Info mencetak pesan informasi ke konsol dengan format [INFO]: message
func Info(msg string) {
	fmt.Println("[INFO]:", msg)
}

// Error mencetak pesan error ke konsol dengan format [ERROR]: message
func Error(msg string) {
	fmt.Println("[ERROR]:", msg)
}

// IsValidEmail memeriksa apakah format email valid (true jika valid, false jika tidak)
func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
