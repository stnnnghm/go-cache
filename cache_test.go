package cache

import (
	"testing"
)

func TestBasic(t *testing.T) {
	dns := New()
	dns.Set("apple.com", "17.253.144.10")
	ip, exists := dns.Get("apple.com")
	if !exists {
		t.Error("apple.com was not found")
	}

	if ip == nil {
		t.Error("dns[apple.com] is nil")
	}

	if ip != "17.253.144.10" {
		t.Error("dns[apple.com] != 17.253.144.10")
	}
}

func TestRemove(t *testing.T) {
	fruits := New()
	fruits.Set("Apple", 1.39)

	applePrice, exists := fruits.Get("Apple")
	if !exists {
		t.Error("Apple price was not set")
	}

	if applePrice == nil {
		t.Error("Apple price is nil")
	}

	if applePrice != 1.39 {
		t.Error("Apple price expected to be 1.39")
	}

	fruits.Remove("Apple")

	applePrice, exists = fruits.Get("Apple")
	if exists {
		t.Error("Apple price was not removed")
	}

	if applePrice != nil {
		t.Error("Apple price is not nil after removal")
	}
}