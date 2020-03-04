package main

import (
	"testing"
)

var plainText = "Kenneth William Miles | GBC4-3FDY-9DYQ-N2F8"
var key = "abcdefghijklmnopqrstuvwxyzabcdef"

func TestSha256(t *testing.T) {
	h := Sha256(plainText)
	//printf "Kenneth William Miles | GBC4-3FDY-9DYQ-N2F8" | openssl sha256"
	if h != "4a13494d676251683db639649843537eb4b85d4ede9baeeff19c99dfc92bdecd" {
		t.Fatal("Hash doesn't match")
	}
}

func TestHmacSha256(t *testing.T) {
	h, err := HmacSha256(plainText, key)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	//printf "Kenneth William Miles | GBC4-3FDY-9DYQ-N2F8" | openssl dgst -sha256 -hmac "abcdefghijklmnopqrstuvwxyzabcdef"
	if h != "0c4201b0ee812d71af1e2d11508277a68c84ca73bf4a33190bcc16fa1a5861b0" {
		t.Fatalf("Has doesn't match %s\n", h)
	}

}

func BenchmarkSha256(b *testing.B) {
	for n := 0; n < b.N; n++ {

		_ = Sha256(plainText)
	}
}

func BenchmarkHmacSha256(b *testing.B) {
	for n := 0; n < b.N; n++ {

		_, err := HmacSha256(plainText, key)
		if err != nil {
			b.Fatal("Error", err)
		}
	}
}
