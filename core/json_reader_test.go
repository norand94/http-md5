package core

import (
	"testing"
)

func TestReadJSON(t *testing.T) {
	data := `{"Id": 10, "Text" : "test_string"}`
	m := ReadJSON(data)

	if m.Id != 10 || m.Text != "test_string" {
		t.Errorf("Введенные данные не соответствуют полученным: id= %v , text = %v", m.Id, m.Text)
	}
}

func TestValidateMessage(t *testing.T) {
	var err error
	var m Message

	err = ValidateMessage(Message{10, "test_string"})
	if err != nil {
		t.Errorf("Ошибка валидации, с сообщением все в порядке: err = %v", err)
	}

	m = Message{-20, "test_string"}
	err = ValidateMessage(m)
	if err == nil {
		t.Errorf("Валидация не сработала, id отрицательный: id = %v", m.Id)
	}

	err = ValidateMessage(Message{0, "test_string"})
	if err == nil {
		t.Error("Валидация не сработала, id отсутствует")
	}


	err = ValidateMessage(Message{10, ""})
	if err == nil {
		t.Error("Валидация не сработала, поле text пустое")
	}

	err = ValidateMessage(Message{10, `
	dfsgksjsjksdfkjjsgijrsisrjfs;fe;fsd;flrgj;grkl
	wfjwfjwjw4rysfksgdfkgdrjwerwljg;lgkrjerijgetl,
	gwejwoijger,grfkerpojijwefkwgjoijwwgjerlgjwrlk
	fqkgjgerg'sehgdrggethrthtetertwfneflwgnerglbes
	sjsglkjgsrgoisrgkrirgioegslforgksekg;ldskgdkgpo
	gjgklreg;lksgokeohkd;krsghks'hkkhseklkwGOPHDS;'K
	sdlgd;gerg;lskdjgdohk;lsfkrg;lkdl;gseks;lgkg;ssg
	`})
	if err == nil {
		t.Error("Валидация не сработала, превышен размер text")
	}
}
