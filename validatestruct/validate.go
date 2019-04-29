package main

import (
	"errors"
	"fmt"
	"reflect"

	validator "gopkg.in/validator.v2"
)

// NewU1Request - Test structure for default validation
type NewU1Request struct {
	Username string `validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Name     string `validate:"nonzero"`
	Age      int    `validate:"min=21"`
}

// NewU2Request - Test structure for custom validation
type NewU2Request struct {
	Devicetype string `validate:"devicetype"`
}

// NewU3Request - Test structure for custom tag
type NewU3Request struct {
	Players int `cricket:"min=11, max=15" footbal:"min=7, max=11"`
}

// ValidateDeviceType -
func ValidateDeviceType(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return validator.ErrUnsupported
	}
	if st.String() != "OLT" || st.String() != "ONU" {
		return errors.New("Provided DeviceType value is not OLT or ONU")
	}
	return nil
}

func main() {

	// Usecase - 1: Print validatation errors
	u1 := NewU1Request{Username: "sarali", Age: 2, Name: "Shivaraj"}
	if errs := validator.Validate(u1); errs != nil {
		fmt.Println("U1: ", errs)
	}

	// Usecase - 2: Adding Custom validation function
	validator.SetValidationFunc("devicetype", ValidateDeviceType)
	u2 := NewU2Request{"HP-Server"}
	if errs := validator.Validate(u2); errs != nil {
		fmt.Println("U2: ", errs)
	}

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
