package hello

import (
	"fmt"

	"github.com/google/uuid"
)

// UUID to print UUID
func UUID() {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uu := u.String()
	fmt.Println("hello, " + uu)
}
