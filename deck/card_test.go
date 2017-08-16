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
