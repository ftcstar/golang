package main

import (
	"fmt"

	validator "gopkg.in/validator.v2"
)

// NewU3Request - Test structure for custom tag
type NewU3Request struct {
	Players int `cricket:"min=11, max=15" footbal:"min=7, max=11"`
}

func main() {
	// Usecase - 3: Custom tag
	validator.SetTag("cricket")
	u3 := NewU3Request{17}
	if errs := validator.Validate(u3); errs != nil {
		fmt.Println("Cricket:U3: ", errs)
	}

	validator.SetTag("footbal")
	u3.Players = 5
	if errs := validator.Validate(u3); errs != nil {
		fmt.Println("Footbal:U3: ", errs)
	}

}
