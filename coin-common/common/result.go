package common

type BizCode int

const SuccessCode BizCode = 0

type Result struct {
	Code    BizCode `json:"code"`
	Message string  `json:"message"`
	Data    any     `json:"data"`
}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Success(data any) *Result {
	r.Code = SuccessCode
	r.Message = "success"
	r.Data = data
	return r
}
func (r *Result) Fail(code BizCode, msg string) *Result {
	r.Code = code
	r.Message = msg
	return r
}

func (r *Result) Deal(data any, err error) *Result {
	if err != nil {
		r.Fail(-999, err.Error())
	} else {
		r.Success(data)
	}
	return r
}
