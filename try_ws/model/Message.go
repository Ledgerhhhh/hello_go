package model

type Message struct {
	Name    string `json:"name"`
	To      string `json:"to"`
	Message string `json:"message"`
}
