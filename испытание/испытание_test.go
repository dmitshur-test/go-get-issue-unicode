package испытание

import "testing"

func TestИспытаниеReturnsИспытание(t *testing.T) {
	got := Испытание()
	want := Эксперимент{Строка: "Испытание!"}
	if got != want {
		t.Errorf("Испытание(): got %#v, want %#v", got, want)
	}
}
