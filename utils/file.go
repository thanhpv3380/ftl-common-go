package utils

import (
	"log"
	"os"
)

func CheckExistOrMake(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal("Unable to create directory:", err)
		}
	}
}
