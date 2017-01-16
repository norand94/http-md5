package core

import (
	"encoding/json"
	"strings"
)

type Message struct {
	Id int
	Text string
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
