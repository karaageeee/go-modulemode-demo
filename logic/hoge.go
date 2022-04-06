package logic

import (
	"fmt"

	"github.com/google/uuid"
)

// PrintUUID to print UUID
func PrintUUID() {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uu := u.String()
	fmt.Println(uu)
}
