package validator

import "github.com/go-playground/validator/v10"

var _validator *validator.Validate

func Init(){
	_validator = validator.New()
}