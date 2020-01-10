package types

type ErrorSerializer struct {
	Code  CodeType `json:"code"`
	Error string   `json:"error"`
}

type Response struct {
	Code CodeType `json:"code"`
}
