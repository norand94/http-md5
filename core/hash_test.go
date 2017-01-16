package core

import "testing"

func TestHash(t *testing.T)  {
	answer := GetHash(10, "test_string")
	if answer == "" {
		t.Error("Ответ не должен быть пустым!")
	}
}
