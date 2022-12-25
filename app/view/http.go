package view

type HttpView interface {
	String(status int, str string)
	Json(status int, json any)
}
