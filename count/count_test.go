package count

import "testing"

func TestCountLetter(t *testing.T) {
	availableSpace, noSpace := CountLetter("こんにちは　 My name is higashi.　")

	if availableSpace != 27 && noSpace != 21 {
		t.Fatal("failed CountLetter test")
	}
}
