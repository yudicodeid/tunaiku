package main

type ResponseModel struct {

	Status bool
	Message string

}

func (model *ResponseModel) Success(msg string) {
	model.Status = true
	model.Message = msg
}

func (model *ResponseModel) Error(err error) {
	model.Status = false
	model.Message = err.Error()
}
