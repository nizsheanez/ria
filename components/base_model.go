package components

import "github.com/astaxie/beego/validation"

type BaseModel struct {

}

func (this *BaseModel) Validate(scenario string) (b bool, err error) {
	validator := validation.Validation{}

	if this, ok := interface{}(this).(Validatable); ok {
		this.Validators(&validator, scenario)
	} else {
		panic("Model don't implement interface Validatable")
	}

	return validator.Valid(this)
}

type Validatable interface {
	Validators(validator *validation.Validation, scenario string)
}
