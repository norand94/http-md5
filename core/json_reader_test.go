package core

import (
	"testing"
)

func TestReadJSON(t *testing.T) {
	//data := "{\"id\": 11,\"text\": \"test_string\"}"
	data := `{"Id": 10, "Text" : "test_string"}`
	m := ReadJSON(data)

	if m.Id != 10 || m.Text != "test_string" {
		t.Errorf("Введенные данные не соответствуют полученным: id= %v , text = %v", m.Id, m.Text)
	}
}
