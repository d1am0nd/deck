package deck

import (
)

type Card struct {
    face string
    value string
    valid bool
}

var vals = map[string]int {
    "A": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "T": 10,
    "J": 11,
    "Q": 12,
    "K": 13 }
var faces = map[string]int {
    "h": 1,
    "s": 2,
    "c": 3,
    "d": 4 }

func NewCard(joined string) Card {
    if len(joined) != 2 {
        return Card{face: "", value: "", valid: false}
    }
    c := Card{face: string(joined[1]), value: string(joined[0]), valid: true}
    if _, ok := vals[c.value]; !ok {
        c.valid = false
    } else if _, ok := faces[c.face]; !ok {
        c.valid = false
    }
    return c
}

func (c *Card) Valid() bool {
    return c.valid
}

func (c *Card) Face() string {
    return c.face
}

func (c *Card) Value() string {
    return c.value
}

func (c *Card) Comparable() int {
    return vals[c.Value()]
}
