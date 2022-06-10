package model

type JsonResponse struct {
	Type string `json: "type"`
	Data []Student `json: "data"`
	Message string `json: "message"`
}