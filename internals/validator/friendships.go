package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Friends(f interface{}) error{
	err := _validator.Struct(f)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		for _,e := range err.(validator.ValidationErrors){
			fmt.Println(e.Tag())
			fmt.Println(e.Field())
		}

		return err
	}
	return nil
}