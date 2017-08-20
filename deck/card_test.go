package deck

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	c := NewCard("Th")
	if !c.valid {
		t.Fatal("NewCard(Th) isnt valid, should be")
	}
	if c.value != "T" {
		t.Fatal("Wrong card value ", c.value, " expected T")
	}
	if c.face != "h" {
		t.Fatal("Wrong card value ", c.face, " expected h")
	}

	c = NewCard("TT")
	if c.valid {
		t.Fatal("NewCard(TT) is vlid, shouldnt be")
	}
}

func TestValid(t *testing.T) {
	c := NewCard("Th")
	if !c.Valid() {
		t.Fatal("NewCard(Th) isnt valid, should be")
	}
	c = NewCard("TT")
	if c.Valid() {
		t.Fatal("NewCard(TT) is valid, shouldnt be")
	}
	c = NewCard("Ths")
	if c.Valid() {
		t.Fatal("NewCard(Ths) is valid, shouldnt be")
	}
	c = NewCard("Ths")
	if c.Valid() {
		t.Fatal("NewCard('') is valid, shouldnt be")
	}
}

func TestFace(t *testing.T) {
	c := NewCard("Th")
	if c.Face() != "h" {
		t.Fatal("NewCard(Th) has face ", c.Face(), " should have h")
	}
}

func TestValue(t *testing.T) {
	c := NewCard("Th")
	if c.Value() != "T" {
		t.Fatal("NewCard(Th) has value ", c.Value(), " should have T")
	}
}

func TestComparable(t *testing.T) {
	c := NewCard("Th")
	if c.Comparable() != 10 {
		t.Fatal("NewCard(Th) has comparable ", c.Comparable(), " should have 10")
	}
}
