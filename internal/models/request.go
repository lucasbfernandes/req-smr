package models

type Request struct {
	Method 	string				`json:"method""`
	Url 	string				`json:"url"`
	Headers map[string][]string	`json:"headers"`
	Body 	[]byte				`json:"body"`
}