package main

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

// You start at "5" and move up (to "2"), left (to "1"), and left (you can't, and stay on "1"), so the first button is 1.
// Starting from the previous button ("1"), you move right twice (to "3") and then down three times (stopping at "9" after two moves and ignoring the third), ending up with 9.
// Continuing from "9", you move left, up, right, down, and left, ending with 8.
// Finally, you move up four times (stopping at "2"), then down once, ending with 5.

func TestAverage(t *testing.T) {
	input := []string{
		"ULL",
		"RRDDD",
		"LURDL",
		"UUUUD",
	}
}
