package utils

import (
	"fmt"
)

func CheckNilErr(err error) {
	if err != nil {
		fmt.Println("Error :", err.Error())
	}
}
