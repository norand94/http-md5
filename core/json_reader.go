package core

import (
	"encoding/json"
	"strings"
	"gopkg.in/go-playground/validator.v9"
)

type Message struct {
	Id int `validate:"min=0,required"`
	Text string `validate:"required,max=100"`
}

func ReadJSON(data string) Message  {
	dec := json.NewDecoder(strings.NewReader(data))

	var m Message
	err := dec.Decode(&m)
	if err != nil {
		panic(err)
	}

	return m

}

func ValidateMessage(m Message) error  {
	validate := validator.New()
	err := validate.Struct(m)
	return err
}

