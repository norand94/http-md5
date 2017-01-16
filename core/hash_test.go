package core

import (
	"testing"
	"fmt"
)

func TestHash(t *testing.T)  {
	answer := GetHash(Message{10, "test_string"})
	fmt.Println(answer)
	if answer == "" {
		t.Error("Ответ не должен быть пустым!")
	}
}
